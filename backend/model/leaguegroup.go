package model

type LeagueGroup struct {
	ID      int    `json:"id"`
	Season  int    `json:"season"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}
