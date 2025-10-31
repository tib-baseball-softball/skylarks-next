package tib

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
	"github.com/tib-baseball-softball/skylarks-next/internal/utility"
)

const (
	cacheLifetimeMinutes = 60
)

var cacheURLAllowlist = []*regexp.Regexp{
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/licenses\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/league_groups/\d+(\.json|/.*)?(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/licenses/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/club_functions\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/fields\.json(?:\?.*)?$`),
}

// isValidBSMURL checks if the given url matches any regex in cacheURLAllowlist
func isValidBSMURL(u *url.URL) bool {
	urlStr := u.String()
	for _, re := range cacheURLAllowlist {
		if re.MatchString(urlStr) {
			return true
		}
	}
	return false
}

func GetLeagueTop10Data(app core.App, leagueID string, statsType string) (bsm.LeaderboardSummary, error) {
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
			apiURL := bsm.GetAPIURL(resource, query, "")

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

func GetCachedBSMResponse(app core.App, url *url.URL) (string, error) {
	var ret string

	if url.Host != os.Getenv("BSM_API_HOST") || !isValidBSMURL(url) {
		return ret, &bsm.URLAllowlistError{URL: url.String(), Message: "Only allowlisted BSM URLs are allowed"}
	}
	hash := utility.GetMD5Hash(url.String())

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

func saveBSMResponseToCache(app core.App, url string) (*core.Record, error) {
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
	cache.SetHash(utility.GetMD5Hash(url))
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
