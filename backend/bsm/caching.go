package bsm

import (
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"net/url"
	"sync"
)

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

			leaderboardData, _, err := FetchResource[model.LeaderboardData](apiURL.String())
			if err != nil {
				app.Logger().Error("Failed to get top10 league data from BSM", "err", err)
				return
			}

			leaderboardData.StatsCategory = endpoint
			leagueLeaderboard.Data = append(leagueLeaderboard.Data, leaderboardData)
		}()
	}

	wg.Wait()
	return leagueLeaderboard, nil
}

func GetCachedBSMResponse[T any](app core.App, url url.URL) (T, error) {

}
