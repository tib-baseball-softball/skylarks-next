package cronjobs

import (
	"encoding/json"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"log"
	"sync"
	"time"
)

func ImportGames(app core.App) {
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
				app.Logger().Error("Error fetching League Group record: ", "error", err)
				return
			}
			if leagueGroup.GetInt("season") != currentYear {
				return
			}

			if errs := app.ExpandRecord(team, []string{"club"}, nil); len(errs) > 0 {
				app.Logger().Error("failed to expand:", "errors", errs)
				return
			}
			club := team.ExpandedOne("club")
			if club == nil {
				app.Logger().Error("Could not load club data for API key")
				return
			}

			matches, err := fetchMatchesForLeagueGroup(team.GetString("bsm_league_group"), club.GetString("bsm_api_key"))
			if err != nil {
				app.Logger().Error("Could not fetch Matches for leagueGroup", "error", err, "team", team, "apiKey", club.GetString("bsm_api_key"))
				return
			}

			err = createOrUpdateEvents(app, matches, team.Id)
			if err != nil {
				app.Logger().Error("Error creating or updating events", "error", err, "team", team.Id)
				return
			}
		}()
	}

	wg.Wait()
	app.Logger().Info("Game Import successfully imported new game events")
}

func fetchMatchesForLeagueGroup(league string, apiKey string) ([]bsm.Match, error) {
	params := make(map[string]string)
	params["filters[leagues][]"] = league
	params["search"] = "skylarks"

	url := bsm.GetAPIURL("matches.json", params, apiKey)
	matches, _, err := bsm.FetchResource[[]bsm.Match](url.String())

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func createOrUpdateEvents(app core.App, matches []bsm.Match, teamID string) (err error) {
	for _, match := range matches {
		record, err := app.FindFirstRecordByData("events", "bsm_id", match.ID)

		// if not found, it throws an error, so create new record
		if err != nil {
			collection, err := app.FindCollectionByNameOrId("events")
			if err != nil {
				app.Logger().Error("Error fetching event collection: ", "error", err)
				continue
			}

			record = core.NewRecord(collection)
			err = setEventRecordValues(app, record, match, teamID)
			if err != nil {
				app.Logger().Error("Error setting event record: ", "error", err)
				continue
			}
		} else {
			// no error - update existing record
			err = setEventRecordValues(app, record, match, teamID)
			if err != nil {
				app.Logger().Error("Error setting event record: ", "error", err)
				continue
			}
		}

		app.Logger().Debug("Record data immediately before persist call", "record", record)
		if err := app.Save(record); err != nil {
			app.Logger().Error("Persisting event record failed", "error", err, "record", record, "teamID", teamID, "match", match)
			continue
		}
	}
	return
}

func setEventRecordValues(app core.App, record *core.Record, match bsm.Match, teamID string) (err error) {
	starttime, err := types.ParseDateTime(match.Time)
	if err != nil {
		return err
	}
	endtime := starttime.Time().Add(time.Hour * 3)
	meetingtime := starttime.Time().Add(-2*time.Hour - 30*time.Minute)

	app.Logger().Debug("Current Date Values", "starttime", starttime, "endtime", endtime, "meetingtime", meetingtime)

	if meetingtime.String() == "" {
		meetingtime = starttime.Time()
	}

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

	app.Logger().Debug("Record populated with fields", "record", record)

	return
}
