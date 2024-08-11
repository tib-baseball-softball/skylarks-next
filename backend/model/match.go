package model

type Match struct {
    ID             int     `json:"id"`
    MatchID        string  `json:"match_id"`
    Time           string  `json:"time"`
    LeagueID       int     `json:"league_id"`
    HomeRuns       *int    `json:"home_runs,omitempty"`
    AwayRuns       *int    `json:"away_runs,omitempty"`
    CreatedAt      string  `json:"created_at"`
    UpdatedAt      string  `json:"updated_at"`
    State          string  `json:"state"`
    Season         int     `json:"season"`
    SecondLeagueID *int    `json:"second_league_id,omitempty"`
    PlannedInnings int     `json:"planned_innings"`
    Name           *string `json:"name,omitempty"`
    HumanState     string  `json:"human_state"`
    ScoresheetURL  *string `json:"scoresheet_url,omitempty"`
    HomeTeamName   string  `json:"home_team_name"`
    AwayTeamName   string  `json:"away_team_name"`
}
