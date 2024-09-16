package model

type League struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Acronym  string `json:"acronym"`
	Sport    string `json:"sport"`
	Sort     int    `json:"sort"`
	Season   int    `json:"season"`
	AgeGroup string `json:"age_group"`
}