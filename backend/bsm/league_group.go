package bsm

import "strconv"

// LeagueGroup represents the BSM structure "Liga/Gruppe".
// It is always unique to single season.
// Caution: several LeagueGroup might have the same League.
type LeagueGroup struct {
	ID      int    `json:"id"`
	Season  int    `json:"season"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
	League  League `json:"league"`
}

// the API key used determines which club LeagueGroups are loaded for
func FetchLeagueGroupsForSeason(apiKey string, season int) ([]LeagueGroup, error) {
	params := make(map[string]string)
	params[SeasonFilter] = strconv.Itoa(season)

	url := GetAPIURL("league_groups.json", params, apiKey)
	leagueGroups, _, err := FetchResource[[]LeagueGroup](url.String())
	if err != nil {
		return nil, err
	}
	return leagueGroups, nil
}
