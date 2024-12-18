package cronjobs

import (
	"encoding/json"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"log"
	"sync"
	"time"
)

func ImportGames(app *pocketbase.PocketBase) {
	teams, err := app.FindRecordsByFilter("teams", "bsm_league_group != 0", "", 0, 0)
	if err != nil {
		log.Print("Error fetching existing team records: ", err)
		return
	}

	currentYear := types.NowDateTime().Time().Year()
	var wg sync.WaitGroup

	for _, team := range teams {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// only run this job for events this season (refreshing games past years should rarely ever be relevant)
			leagueGroup, err := app.FindFirstRecordByData("leaguegroups", "bsm_id", team.GetInt("bsm_league_group"))
			if err != nil {
				log.Print("Error fetching League Group record: ", err)
				return
			}
			if leagueGroup.GetInt("season") != currentYear {
				return
			}

			if errs := app.ExpandRecord(team, []string{"club"}, nil); len(errs) > 0 {
				log.Printf("failed to expand: %v", errs)
				return
			}
			club := team.ExpandedOne("club")
			if club == nil {
				log.Print("ERROR: Could not load club data for API key, aborting...")
				return
			}

			matches, err := fetchMatchesForLeagueGroup(team.GetString("bsm_league_group"), club.GetString("bsm_api_key"))
			if err != nil {
				log.Print(err)
				return
			}

			err = createOrUpdateEvents(app, matches, team.Id)
			if err != nil {
				log.Print(err)
				return
			}
		}()
	}

	wg.Wait()
	log.Println("Game Import successfully imported new game events")
}

func fetchMatchesForLeagueGroup(league string, apiKey string) ([]model.Match, error) {
	params := make(map[string]string)
	params["filters[leagues][]"] = league
	params["search"] = "skylarks"

	url := bsm.GetAPIURL("matches.json", params, apiKey)
	matches, _, err := bsm.FetchResource[[]model.Match](url.String())

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func createOrUpdateEvents(app *pocketbase.PocketBase, matches []model.Match, teamID string) (err error) {
	for _, match := range matches {
		record, err := app.FindFirstRecordByData("events", "bsm_id", match.ID)

		// if not found, it throws an error, so create new record
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("events")
			if err != nil {
				return err
			}

			record = core.NewRecord(collection)
			err = setEventRecordValues(record, match, teamID)
			if err != nil {
				return err
			}
		}
		// no error - update existing record
		err = setEventRecordValues(record, match, teamID)
		if err != nil {
			return err
		}

		if err := app.Save(record); err != nil {
			log.Print("Persisting event record failed: ", err)
			return err
		}
	}
	return
}

func setEventRecordValues(record *core.Record, match model.Match, teamID string) (err error) {
	starttime, err := types.ParseDateTime(match.Time)
	if err != nil {
		return err
	}
	endtime := starttime.Time().Add(time.Hour * 3)
	meetingtime := starttime.Time().Add(-2*time.Hour - 30*time.Minute)

	matchJSON, err := json.Marshal(match)
	if err != nil {
		return err
	}

	record.Set("title", match.AwayTeamName+" @ "+match.HomeTeamName)
	record.Set("bsm_id", match.ID)
	record.Set("starttime", starttime.String())
	record.Set("endtime", endtime.String())
	record.Set("meetingtime", meetingtime.String())
	record.Set("type", "game")
	record.Set("team", teamID)
	record.Set("match_json", string(matchJSON))

	return
}
