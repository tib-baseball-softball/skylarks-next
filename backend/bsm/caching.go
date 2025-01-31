package bsm

import (
	"encoding/json"
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"github.com/tib-baseball-softball/skylarks-next/utility"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"
)

const (
	cacheLifetimeMinutes = 60
)

var cacheURLWhitelist = []*regexp.Regexp{
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+/licenses\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/clubs/\d+\.json(?:\?.*)?$`),
	regexp.MustCompile(`^https://bsm\.baseball-softball\.de/league_groups/\d+(\.json|/.*)?(?:\?.*)?$`),
}

// isValidBSMURL checks if the given url matches any regex in cacheURLWhitelist
func isValidBSMURL(u *url.URL) bool {
	urlStr := u.String()
	for _, re := range cacheURLWhitelist {
		if re.MatchString(urlStr) {
			return true
		}
	}
	return false
}

func GetLeagueTop10Data(app core.App, leagueID string, statsType string) (model.LeaderboardSummary, error) {
	leagueLeaderboard := model.LeaderboardSummary{
		LeagueID:  leagueID,
		StatsType: statsType,
		Data:      make([]model.LeaderboardData, 0),
	}

	var statsEndpoints []string

	switch statsType {
	case TypeBatting:
		statsEndpoints = GetBattingStatsTypes()
	case TypePitching:
		statsEndpoints = GetPitchingStatsTypes()
	case TypeFielding:
		statsEndpoints = GetFieldingStatsTypes()
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
			apiURL := GetAPIURL(resource, query, "")

			rawResponseData, err := GetCachedBSMResponse(app, apiURL)
			if err != nil {
				return
			}
			var leaderboardData model.LeaderboardData
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
		return ret, &URLWhitelistError{url.String(), "Only whitelisted BSM URLs are allowed"}
	}
	hash := utility.GetMD5Hash(url.String())

	var record *core.Record
	record, err := app.FindFirstRecordByData("requestcaches", "hash", hash)
	if err != nil {
		// no data - request has never been cached before
		record, err = saveToCache(app, url.String())
		if err != nil {
			return ret, err
		}
	}

	if record != nil {
		currentTime := types.NowDateTime()
		updated := record.GetDateTime("updated")
		cutoff := updated.Add(cacheLifetimeMinutes * time.Minute)

		if cutoff.Before(currentTime) {
			// cache is outdated, delete and load new

			err = app.Delete(record)
			if err != nil {
				return ret, err
			}

			record, err = saveToCache(app, url.String())
			if err != nil {
				return ret, err
			}
		}
		ret = record.GetString("responseBody")
	}

	return ret, nil
}

func saveToCache(app core.App, url string) (*core.Record, error) {
	_, body, err := FetchResource[any](url)
	if err != nil {
		app.Logger().Error("Failed to get data from BSM", "err", err)
		return nil, err
	}

	collection, err := app.FindCollectionByNameOrId("requestcaches")
	if err != nil {
		return nil, err
	}

	record := core.NewRecord(collection)
	record.Set("responseBody", body)
	record.Set("hash", utility.GetMD5Hash(url))
	record.Set("url", url)

	err = app.Save(record)
	if err != nil {
		return record, err
	}

	return record, nil
}
