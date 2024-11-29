package model

type LeaderboardDataset struct {
	Person      Person      `json:"person"`
	LeagueEntry LeagueEntry `json:"league_entry"`
	Rank        int         `json:"rank"`
	Value       interface{} `json:"value"` // BSM returns both strings and integers for different endpoints
}

type LeaderboardData struct {
	League        LeagueGroup          `json:"league"`
	StatsCategory string               `json:"stats_category"`
	Data          []LeaderboardDataset `json:"data"`
}

type LeaderboardSummary struct {
	LeagueID  string            `json:"league_id"`
	StatsType string            `json:"stats_type"`
	Data      []LeaderboardData `json:"data"`
}
