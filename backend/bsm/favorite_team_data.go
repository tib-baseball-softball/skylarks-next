package bsm

import (
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/tib"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
)

// LoadHomeData loads structured data to display on a single team's home dashboard.
func LoadHomeData(app core.App, teamID int) ([]HomeDataset, error) {
	apiKey := os.Getenv("BSM_API_KEY")
	if apiKey == "" {
		return nil, errors.New("default API key not set, aborting")
	}

	teamChannel := make(chan Team)
	leagueChannel := make(chan []LeagueGroup)

	go func() {
		team, err := loadSingleTeamByID(teamID, apiKey)
		if err != nil {
			app.Logger().Error("Error fetching team", "error", err, "team", teamID)
			return
		}
		teamChannel <- team
	}()

	go func() {
		leagueGroup, err := fetchLeagueGroupsForCurrentSeason(apiKey, time.Now().Year())
		if err != nil {
			app.Logger().Error("Error fetching league groups", "error", err, "team", teamID)
			return
		}
		leagueChannel <- leagueGroup
	}()

	clubTeam := <-teamChannel
	leagueGroups := <-leagueChannel

	if len(clubTeam.LeagueEntries) == 0 {
		return nil, errors.New("team does not have any league entries associated with it")
	}

	filteredLeagueGroups := slices.Collect(func(yield func(LeagueGroup) bool) {
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
		go func(clubTeam Team) {
			defer wg.Done()

			var dataset HomeDataset
			dataset.LeagueGroup = leagueGroup

			var wg sync.WaitGroup

			wg.Add(1)
			go func() {
				defer wg.Done()

				var matches []Match

				params := make(map[string]string)
				params[SeasonFilter] = strconv.Itoa(leagueGroup.Season)
				params[SearchFilter] = tib.TeamSearchParam
				params[LeagueFilter] = strconv.Itoa(leagueGroup.ID)
				params[GamedayFilter] = GameDayAny

				matches, err := loadMatchesWithFilterParams(params, apiKey)
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
			go func(clubTeam Team) {
				defer wg.Done()

				err := LoadHomeTeamTable(clubTeam, leagueGroup.ID, &dataset, apiKey)
				if err != nil {
					app.Logger().Error("Error fetching home team table", "error", err, "team", clubTeam.ID, "leagueGroup", leagueGroup.ID)
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

func createPlayoffSeriesData(dataset *HomeDataset, games []Match) {
	playoffGames := slices.Collect(func(yield func(Match) bool) {
		for _, game := range games {
			if game.IsPlayoffGame() {
				if !yield(game) {
					return
				}
			}
		}
	})
	playoffSeries := NewPlayoffSeries(playoffGames)

	playoffSeries.GetSeriesStatus()
	dataset.PlayoffSeries = &playoffSeries
}

// LoadHomeTeamTable loads the home team table
func LoadHomeTeamTable(team Team, leagueGroupID int, dataset *HomeDataset, apiKey string) error {
	table, err := loadSingleTable(leagueGroupID, apiKey)
	if err != nil {
		return err
	}

	row := determineTableRow(team, &table)

	dataset.Table = table
	dataset.TableRow = row

	return nil
}
