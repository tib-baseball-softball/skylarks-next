package bsm

import (
	"sync"
)

// HomeDataset represents the data for a team's home page
type HomeDataset struct {
	LeagueGroup   LeagueGroup    `json:"league_group"`
	NextGame      *Match         `json:"next_game,omitempty"`
	LastGame      *Match         `json:"last_game,omitempty"`
	PlayoffSeries *PlayoffSeries `json:"playoff_series,omitempty"`
	Table         Table          `json:"table"`
	TableRow      Row            `json:"table_row"`
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

// TODO: implement
// createStreakDataEntries creates streak data entries for a team
//func createStreakDataEntries(games []Match, row Row) []StreakDataEntry {
//	var entries []StreakDataEntry
//	wins := 0
//
//	for i, game := range games {
//		if game.State != "planned" && game.State != "cancelled" &&
//			(game.AwayTeamName == row.TeamName || game.HomeTeamName == row.TeamName) {
//			number := i + 1
//			won := false
//
//			if skylarksWon := game.SkylarksWin(); skylarksWon != nil {
//				if *skylarksWon {
//					wins++
//					won = true
//				}
//			}
//			entries = append(entries, StreakDataEntry{
//				Game:      strconv.Itoa(number),
//				WonGame:   won,
//				WinsCount: wins,
//			})
//		}
//	}
//	return entries
//}
