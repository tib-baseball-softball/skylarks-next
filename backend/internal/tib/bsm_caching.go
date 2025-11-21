package tib

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	cacheLifetimeMinutes = 60
)

var cacheURLAllowlist = []*regexp.Regexp{
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/licenses\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/league_groups(/\d+)?(\.json|/.*)?(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/licenses/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/matches/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/club_functions\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/fields\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/matches\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/matches\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/(teams|team_clubs)\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/league_groups/\d+/table\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/(league_entries|people|clubs)/\d+/(statistics)(/.*)?\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/league_entries/\d+/(stats|statistics)(/.*)?\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/matches/\d+/match_boxscore\.json(?:\?.*)?$`),
}

type CachingApp interface {
	FindFirstRecordByData(collectionModelOrIdentifier any, key string, value any) (*core.Record, error)
	Logger() *slog.Logger
	Delete(model core.Model) error
	Save(model core.Model) error
	FindCollectionByNameOrId(nameOrId string) (*core.Collection, error)
}

// isValidBSMURL checks if the given url matches any regex in the cacheURLAllowlist
func isValidBSMURL(u *url.URL) bool {
	urlStr := u.String()
	for _, re := range cacheURLAllowlist {
		if re.MatchString(urlStr) {
			return true
		}
	}
	return false
}

func GetLeagueTop10Data(app CachingApp, client bsm.BaseClient, leagueID string, statsType string) (bsm.LeaderboardSummary, error) {
	leagueLeaderboard := bsm.LeaderboardSummary{
		LeagueID:  leagueID,
		StatsType: statsType,
		Data:      make([]bsm.LeaderboardData, 0),
	}

	var statsEndpoints []string

	switch statsType {
	case bsm.TypeBatting:
		statsEndpoints = bsm.GetBattingStatsTypes()
	case bsm.TypePitching:
		statsEndpoints = bsm.GetPitchingStatsTypes()
	case bsm.TypeFielding:
		statsEndpoints = bsm.GetFieldingStatsTypes()
	default:
		return leagueLeaderboard, errors.New("unknown stats type")
	}

	var wg sync.WaitGroup

	for _, endpoint := range statsEndpoints {
		wg.Add(1)
		go func() {
			defer wg.Done()
			query := make(map[string]string)
			resource := "league_groups/" + leagueID + "/top10/" + statsType + "/" + endpoint + ".json"
			apiURL := client.GetAPIURL(resource, query, "")

			rawResponseData, err := GetCachedBSMResponse(app, apiURL)
			if err != nil {
				return
			}
			var leaderboardData bsm.LeaderboardData
			err = json.Unmarshal([]byte(rawResponseData), &leaderboardData)
			if err != nil {
				return
			}

			leaderboardData.StatsCategory = endpoint
			leagueLeaderboard.Data = append(leagueLeaderboard.Data, leaderboardData)
		}()
	}

	wg.Wait()
	return leagueLeaderboard, nil
}

func GetCachedBSMResponse(app CachingApp, url *url.URL) (string, error) {
	var ret string

	if url.Host != os.Getenv("BSM_API_HOST") || !isValidBSMURL(url) {
		return ret, &bsm.URLAllowlistError{URL: url.Path, Message: "Only allowlisted BSM URLs are allowed"}
	}
	hash := dp.GetMD5Hash(url.String())

	var requestCache *dp.RequestCache
	record, err := app.FindFirstRecordByData(dp.RequestCacheCollection, "hash", hash)
	if err != nil {
		// no data - request has never been cached before
		record, err = saveBSMResponseToCache(app, url.String())
		if err != nil {
			return ret, err
		}
	}

	if record != nil {
		requestCache = &dp.RequestCache{}
		requestCache.SetProxyRecord(record)

		if isOutdated(requestCache.Updated()) {
			err = app.Delete(record)
			if err != nil {
				return ret, err
			}

			record, err = saveBSMResponseToCache(app, url.String())
			if err != nil {
				return ret, err
			}

			// Update requestCache with the new record
			requestCache = &dp.RequestCache{}
			requestCache.SetProxyRecord(record)
		}
		ret = requestCache.ResponseBody()
	}

	return ret, nil
}

func saveBSMResponseToCache(app CachingApp, url string) (*core.Record, error) {
	_, body, err := bsm.FetchResource[any](url)
	if err != nil {
		app.Logger().Error("Failed to get data from BSM", "err", err, "url", url)
		return nil, err
	}

	collection, err := app.FindCollectionByNameOrId(dp.RequestCacheCollection)
	if err != nil {
		return nil, err
	}

	record := core.NewRecord(collection)
	cache := &dp.RequestCache{}
	cache.SetProxyRecord(record)

	cache.SetResponseBody(body)
	cache.SetHash(dp.GetMD5Hash(url))
	cache.SetURL(url)

	err = app.Save(record)
	if err != nil {
		return record, err
	}

	return record, nil
}

func isOutdated(cacheTime types.DateTime) bool {
	currentTime := types.NowDateTime()
	cutoff := cacheTime.Add(cacheLifetimeMinutes * time.Minute)

	return currentTime.After(cutoff)
}

func rewriteURLForProxying(url url.URL) url.URL {
	bsmHost := os.Getenv("BSM_API_HOST")
	targetURL := url
	targetURL.Scheme = "https"
	targetURL.Host = bsmHost
	targetURL.Path = strings.TrimPrefix(url.Path, "/api/bsm/relay")

	newQuery := url.Query()
	newQuery.Set("api_key", os.Getenv("BSM_API_KEY"))
	targetURL.RawQuery = newQuery.Encode()

	return targetURL
}
