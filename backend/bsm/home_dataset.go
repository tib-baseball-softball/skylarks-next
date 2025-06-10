package bsm

import (
	"strconv"
	"sync"
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
