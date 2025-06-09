package bsm

// SeriesStatus represents the status of a playoff series
type SeriesStatus int

const (
	SeriesStatusTied SeriesStatus = iota
	SeriesStatusLeading
	SeriesStatusWon
)

// PlayoffTeam represents a team in a playoff series
type PlayoffTeam struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

// PlayoffSeries represents a playoff series between two teams
type PlayoffSeries struct {
	SeriesGames  []Match                `json:"series_games"`
	Status       SeriesStatus           `json:"status"`
	Teams        map[string]PlayoffTeam `json:"teams"`
	LeadingTeam  *PlayoffTeam           `json:"leading_team,omitempty"`
	TrailingTeam *PlayoffTeam           `json:"trailing_team,omitempty"`
	SeriesLength int                    `json:"series_length"`
}

// NewPlayoffSeries creates a new PlayoffSeries instance
func NewPlayoffSeries(games []Match) PlayoffSeries {
	return PlayoffSeries{
		Status:       SeriesStatusTied,
		Teams:        make(map[string]PlayoffTeam),
		SeriesLength: 0,
		SeriesGames:  games,
		LeadingTeam:  nil,
		TrailingTeam: nil,
	}
}

// GetSeriesStatus calculates the current status of the playoff series.
// This function must be called with the series of games already set to receive a valid object.
func (ps *PlayoffSeries) GetSeriesStatus() {
	// Early return if no games
	if len(ps.SeriesGames) == 0 {
		return
	}

	ps.SeriesLength = len(ps.SeriesGames)

	// Initialize teams if not already present
	firstTeamName := ps.SeriesGames[0].AwayTeamName
	secondTeamName := ps.SeriesGames[0].HomeTeamName

	if _, exists := ps.Teams[firstTeamName]; !exists {
		ps.Teams[firstTeamName] = PlayoffTeam{Name: firstTeamName, Wins: 0}
	}

	if _, exists := ps.Teams[secondTeamName]; !exists {
		ps.Teams[secondTeamName] = PlayoffTeam{Name: secondTeamName, Wins: 0}
	}

	// Reset win counts
	firstTeam := ps.Teams[firstTeamName]
	firstTeam.Wins = 0
	ps.Teams[firstTeamName] = firstTeam

	secondTeam := ps.Teams[secondTeamName]
	secondTeam.Wins = 0
	ps.Teams[secondTeamName] = secondTeam

	// Count wins for each team using a single loop
	for _, match := range ps.SeriesGames {
		winner := match.GetWinnerForMatch()
		var winnerName string

		switch winner {
		case GameWinnerHome:
			winnerName = match.HomeTeamName
		case GameWinnerAway:
			winnerName = match.AwayTeamName
		default:
			continue
		}

		team := ps.Teams[winnerName]
		team.Wins++
		ps.Teams[winnerName] = team
	}

	// Get updated team references
	firstTeam = ps.Teams[firstTeamName]
	secondTeam = ps.Teams[secondTeamName]

	// Series hasn't started yet (no wins)
	if firstTeam.Wins == 0 && secondTeam.Wins == 0 {
		return
	}

	// Calculate series status more efficiently
	ps.LeadingTeam = nil
	ps.TrailingTeam = nil

	if firstTeam.Wins == secondTeam.Wins {
		ps.Status = SeriesStatusTied
		return
	}

	// Determine the leading and trailing teams once
	var leading, trailing PlayoffTeam
	if firstTeam.Wins > secondTeam.Wins {
		leading, trailing = firstTeam, secondTeam
	} else {
		leading, trailing = secondTeam, firstTeam
	}

	// Set status based on win count
	if leading.Wins > ps.SeriesLength/2 {
		ps.Status = SeriesStatusWon
	} else {
		ps.Status = SeriesStatusLeading
	}

	// Store pointers to the teams
	ps.LeadingTeam = &leading
	ps.TrailingTeam = &trailing
}
