package cronjobs

import (
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/tools/types"
    "github.com/tib-baseball-softball/skylarks-next/bsm"
    "github.com/tib-baseball-softball/skylarks-next/model"
    "log"
    "strconv"
    "sync"
)


func ImportLeagueGroups(app *pocketbase.PocketBase) (err error)  {
	clubs, err := app.Dao().FindRecordsByFilter("clubs", "bsm_id != 0", "", 0, 0)
	if err != nil {
        return err
    }
	
	
	var wg sync.WaitGroup
	
	for _, club := range clubs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			leagueGroups, err := fetchLeagueGroupsForCurrentSeason(club.GetString("bsm_api_key"))
			if err != nil {
				log.Print(err)
				return
			}
			err = createOrUpdateLeagueGroups(app, leagueGroups)
			if err != nil {
                log.Print(err)
                return
            }
		}()
	}
	
	wg.Wait()
	log.Println("League Group Import successfully imported all league groups")
	return nil
}

// the API key used determines which club LeagueGroups are loaded for
func fetchLeagueGroupsForCurrentSeason(apiKey string) ([]model.LeagueGroup, error) {
	currentYear := types.NowDateTime().Time().Year()
	
	params := make(map[string]string)
	params["filters[seasons][]"] = strconv.Itoa(currentYear)
	
	url := bsm.GetAPIURL("league_groups.json", params, apiKey)
	leagueGroups, _, err := bsm.FetchResource[[]model.LeagueGroup](url.String())
	if err != nil {
		return nil, err
	}
    return leagueGroups, nil
}

// TODO: try to make generic
func createOrUpdateLeagueGroups(app *pocketbase.PocketBase, leagueGroups []model.LeagueGroup) (err error) {
    for _, leagueGroup := range leagueGroups {
        record, err := app.Dao().FindFirstRecordByData("leaguegroups", "bsm_id", leagueGroup.ID)

        // if not found, it throws an error, so create new record
        if err != nil {
            collection, err := app.Dao().FindCollectionByNameOrId("leaguegroups")
            if err != nil {
                return err
            }

            record = models.NewRecord(collection)
            err = setLeagueGroupRecordValues(record, leagueGroup)
            if err != nil {
                return err
            }
        }
        // no error - update existing record
        err = setLeagueGroupRecordValues(record, leagueGroup)
        if err != nil {
            return err
        }

        if err := app.Dao().SaveRecord(record); err != nil {
            log.Print("Persisting LeagueGroup record failed: ", err)
            return err
        }
    }
    return
}

func setLeagueGroupRecordValues(record *models.Record, leagueGroup model.LeagueGroup) (err error) {
	record.Set("bsm_id", leagueGroup.ID)
	record.Set("season", leagueGroup.Season)
	record.Set("name", leagueGroup.Name)
	record.Set("acronym", leagueGroup.Acronym)

    return
}