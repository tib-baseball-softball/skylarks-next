package bsm

import (
	"encoding/json"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/pb"
	"strconv"
	"sync"
	"time"
)

// HomeDataset represents the data for a team's home page
type HomeDataset struct {
	LeagueGroup   LeagueGroup       `json:"league_group"`
	NextGame      *Match            `json:"next_game,omitempty"`
	LastGame      *Match            `json:"last_game,omitempty"`
	PlayoffSeries *PlayoffSeries    `json:"playoff_series,omitempty"`
	Table         Table             `json:"table"`
	TableRow      Row               `json:"table_row"`
	StreakData    []StreakDataEntry `json:"streak_data"`
}

type DatasetContainer struct {
	mu       sync.Mutex
	datasets []HomeDataset
}

func (c *DatasetContainer) Add(dataset HomeDataset) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.datasets = append(c.datasets, dataset)
}

// StreakDataEntry represents a single entry in a team's streak data
type StreakDataEntry struct {
	Game      string `json:"game"`
	WonGame   bool   `json:"won_game"`
	WinsCount int    `json:"wins_count"`
}

type DisplayGames struct {
	Next *Match
	Last *Match
}

func GetCachedDatasetResponse(app core.App, teamID int, season int) ([]byte, error) {
	var ret []byte
	cacheIdentifier := strconv.Itoa(season) + "_" + strconv.Itoa(teamID) + "_home_dataset"

	requestCache := &pb.RequestCache{}
	cacheRecord, err := pb.GetCacheEntryByIdentifier(app, cacheIdentifier)
	requestCache.SetProxyRecord(cacheRecord)

	if err != nil || (requestCache != nil && isOutdated(requestCache.Updated()) && time.Now().Year() <= season) {
		// This identifier has not been cached before or has expired.
		// For this specific set of data, we assume it's not going to change after the season ends.

		if cacheRecord != nil {
			err := app.Delete(cacheRecord)
			if err != nil {
				return nil, err
			}
		}

		homeDatasets, err := LoadHomeData(app, teamID)
		if err != nil {
			app.Logger().Error("Failed to load home data", "error", err)
			return ret, err
		}

		ret, err = json.Marshal(homeDatasets)
		if err != nil {
			app.Logger().Error("Failed to marshal home data", "error", err)
			return ret, err
		}
		err = pb.SaveDataToCache(app, cacheIdentifier, string(ret))
		if err != nil {
			app.Logger().Error("Failed to save data to cache", "error", err, "cacheIdentifier", cacheIdentifier)
			return nil, err
		}
		return ret, nil
	} else {
		// cache hit
		ret = []byte(requestCache.ResponseBody())
		return ret, nil
	}
}

// createStreakDataEntries creates streak data entries for a team. Usable for data visualization in a chart.
func createStreakDataEntries(games []Match, teamName string) []StreakDataEntry {
	var entries []StreakDataEntry
	wins := 0

	for i, game := range games {
		state := game.GetMatchState(teamName)
		if state == MatchStateCancelled || state == MatchStateNotYetPlayed {
			continue
		}

		won := state == MatchStateWon

		if won {
			wins++
		}

		entries = append(entries, StreakDataEntry{
			Game:      strconv.Itoa(i + 1),
			WonGame:   won,
			WinsCount: wins,
		})
	}
	return entries
}
