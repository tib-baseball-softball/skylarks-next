package cronjobs

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"strconv"
	"sync"
)

// ImportLeagueGroups imports league groups concurrently, either for one given club or all clubs in the database.
func ImportLeagueGroups(app *pocketbase.PocketBase, clubID *string, season *int) (err error) {
	filter := "bsm_id != 0"

	if clubID != nil {
		filter = filter + " && id = '" + *clubID + "'"
	}

	clubs, err := app.FindRecordsByFilter("clubs", filter, "", 0, 0)
	if err != nil {
		return err
	}

	selectedSeason := types.NowDateTime().Time().Year()
	if season != nil {
		selectedSeason = *season
	}

	var wg sync.WaitGroup

	for _, club := range clubs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			leagueGroups, err := fetchLeagueGroupsForCurrentSeason(club.GetString("bsm_api_key"), selectedSeason)
			if err != nil {
				app.Logger().Error(err.Error())
				return
			}
			err = createOrUpdateLeagueGroups(app, leagueGroups, club)
			if err != nil {
				app.Logger().Error(err.Error())
				return
			}
		}()
	}

	wg.Wait()
	app.Logger().Info("League Group Import successfully imported all league groups")
	return nil
}

// the API key used determines which club LeagueGroups are loaded for
func fetchLeagueGroupsForCurrentSeason(apiKey string, season int) ([]model.LeagueGroup, error) {
	params := make(map[string]string)
	params["filters[seasons][]"] = strconv.Itoa(season)

	url := bsm.GetAPIURL("league_groups.json", params, apiKey)
	leagueGroups, _, err := bsm.FetchResource[[]model.LeagueGroup](url.String())
	if err != nil {
		return nil, err
	}
	return leagueGroups, nil
}

func createOrUpdateLeagueGroups(app *pocketbase.PocketBase, leagueGroups []model.LeagueGroup, club *core.Record) (err error) {
	for _, leagueGroup := range leagueGroups {
		record, err := app.FindFirstRecordByData("leaguegroups", "bsm_id", leagueGroup.ID)

		// if not found, it throws an error, so create new record
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("leaguegroups")
			if err != nil {
				return err
			}

			record = core.NewRecord(collection)
			err = setLeagueGroupRecordValues(record, leagueGroup, club)
			if err != nil {
				return err
			}
		}
		// no error - update existing record
		err = setLeagueGroupRecordValues(record, leagueGroup, club)
		if err != nil {
			return err
		}

		if err := app.Save(record); err != nil {
			app.Logger().Error("Persisting LeagueGroup record failed: ", err)
			return err
		}
	}
	return
}

func setLeagueGroupRecordValues(record *core.Record, leagueGroup model.LeagueGroup, club *core.Record) (err error) {
	record.Set("bsm_id", leagueGroup.ID)
	record.Set("season", leagueGroup.Season)
	record.Set("name", leagueGroup.Name)
	record.Set("acronym", leagueGroup.Acronym)
	record.Set("clubs+", club.Id)

	return
}
