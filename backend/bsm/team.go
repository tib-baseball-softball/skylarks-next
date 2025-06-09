package bsm

import "strconv"

type Team struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Season        int           `json:"season"`
	ShortName     string        `json:"short_name"`
	Clubs         []Club        `json:"clubs"`
	LeagueEntries []LeagueEntry `json:"league_entries"`
}

func loadSingleTeamByID(ID int, apiKey string) (Team, error) {
	var team Team
	url := GetAPIURL("teams/"+strconv.Itoa(ID)+".json", make(map[string]string), apiKey)
	team, _, err := FetchResource[Team](url.String())
	if err != nil {
		return team, err
	}
	return team, nil
}
