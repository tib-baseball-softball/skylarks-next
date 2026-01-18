package tib

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/diamond-planner/diamond-planner/dp"
)

// LoadHomeData loads structured data to display on a single team's home dashboard.
func LoadHomeData(app dp.LogOnlyApp, client bsm.APIClient, teamID int, season int) ([]HomeDataset, error) {
	apiKey := os.Getenv("BSM_API_KEY")
	if apiKey == "" {
		return nil, errors.New("default API key not set, aborting")
	}

	teamChannel := make(chan bsm.Team)
	leagueChannel := make(chan []bsm.LeagueGroup)

	go func() {
		team, err := client.LoadSingleTeamByID(teamID, apiKey)
		if err != nil {
			app.Logger().Error("Error fetching team", "error", err, "team", teamID)
			return
		}
		teamChannel <- team
	}()

	go func() {
		leagueGroupsResponse, err := client.FetchLeagueGroupsForSeason(apiKey, season)
		if err != nil {
			app.Logger().Error("Error fetching league groups", "error", err, "team", teamID)
			return
		}
		leagueChannel <- leagueGroupsResponse
	}()

	clubTeam := <-teamChannel
	leagueGroups := <-leagueChannel

	if len(clubTeam.LeagueEntries) == 0 {
		return nil, errors.New("team does not have any league entries associated with it")
	}

	filteredLeagueGroups := slices.Collect(func(yield func(bsm.LeagueGroup) bool) {
		for _, leagueGroup := range leagueGroups {
			if leagueGroup.League.ID == clubTeam.LeagueEntries[0].League.ID {
				if !yield(leagueGroup) {
					return
				}
			}
		}
	})

	datasetContainer := DatasetContainer{
		datasets: make([]HomeDataset, 0),
	}

	var wg sync.WaitGroup

	for _, leagueGroup := range filteredLeagueGroups {
		wg.Add(1)
		go func(clubTeam bsm.Team) {
			defer wg.Done()

			var dataset HomeDataset
			dataset.LeagueGroup = leagueGroup
			dataset.TeamID = teamID
			dataset.Season = season
			dataset.LeagueGroupID = leagueGroup.ID

			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				defer wg.Done()

				var matches []bsm.Match

				params := make(map[string]string)
				params[bsm.SeasonFilter] = strconv.Itoa(leagueGroup.Season)
				params[bsm.SearchFilter] = TeamSearchParam
				params[bsm.LeagueFilter] = strconv.Itoa(leagueGroup.ID)
				params[bsm.GamedayFilter] = bsm.GameDayAny
				params[bsm.CompactFilter] = "true"

				matches, err := client.LoadMatchesWithFilterParams(params, apiKey)
				if err != nil {
					app.Logger().Error("Error fetching matches", "error", err, "team", teamID)
					return
				}
				displayGames := findNextAndPreviousGame(matches, time.Now())

				if displayGames.Next != nil {
					dataset.NextGame = displayGames.Next
				}

				if displayGames.Last != nil {
					dataset.LastGame = displayGames.Last
				}
				createPlayoffSeriesData(&dataset, matches)
				dataset.StreakData = createStreakDataEntries(matches, clubTeam.Name)
			}()

			wg.Add(1)
			go func(clubTeam bsm.Team) {
				defer wg.Done()

				err := LoadHomeTeamTable(client, clubTeam, leagueGroup.ID, &dataset, apiKey)
				if err != nil {
					app.Logger().Warn("Error fetching home team table", "error", err, "team", clubTeam.ID, "leagueGroup", leagueGroup.ID)
					return
				}
			}(clubTeam)
			wg.Wait()

			datasetContainer.Add(dataset)
		}(clubTeam)
	}
	wg.Wait()

	return datasetContainer.datasets, nil
}

func createPlayoffSeriesData(dataset *HomeDataset, games []bsm.Match) {
	playoffGames := slices.Collect(func(yield func(bsm.Match) bool) {
		for _, game := range games {
			if game.IsPlayoffGame() {
				if !yield(game) {
					return
				}
			}
		}
	})
	playoffSeries := bsm.NewPlayoffSeries(playoffGames)

	playoffSeries.GetSeriesStatus()
	dataset.PlayoffSeries = &playoffSeries
}

// LoadHomeTeamTable loads the home team table
func LoadHomeTeamTable(client bsm.APIClient, team bsm.Team, leagueGroupID int, dataset *HomeDataset, apiKey string) error {
	table, err := client.LoadSingleTable(leagueGroupID, apiKey)
	if err != nil {
		return err
	}

	row := bsm.DetermineTableRow(team, &table)

	dataset.Table = table
	dataset.TableRow = row

	return nil
}
