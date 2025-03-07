package bsm

import (
	"encoding/json"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"log"
	"sync"
	"time"
)

type GameImportService struct {
	App core.App
}

func (s GameImportService) ImportGames() {
	teams, err := s.App.FindRecordsByFilter("teams", "bsm_league_group != 0", "", 0, 0)
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
			leagueGroup, err := s.App.FindFirstRecordByData("leaguegroups", "bsm_id", team.GetInt("bsm_league_group"))
			if err != nil {
				s.App.Logger().Error("Error fetching League Group record: ", "error", err)
				return
			}
			if leagueGroup.GetInt("season") != currentYear {
				return
			}

			if errs := s.App.ExpandRecord(team, []string{"club"}, nil); len(errs) > 0 {
				s.App.Logger().Error("failed to expand:", "errors", errs)
				return
			}
			club := team.ExpandedOne("club")
			if club == nil {
				s.App.Logger().Error("Could not load club data for API key")
				return
			}

			matches, err := s.fetchMatchesForLeagueGroup(team.GetString("bsm_league_group"), club.GetString("bsm_api_key"))
			if err != nil {
				s.App.Logger().Error("Could not fetch Matches for leagueGroup", "error", err, "team", team, "apiKey", club.GetString("bsm_api_key"))
				return
			}

			err = s.createOrUpdateEvents(matches, team.Id)
			if err != nil {
				s.App.Logger().Error("Error creating or updating events", "error", err, "team", team.Id)
				return
			}
		}()
	}

	wg.Wait()
	s.App.Logger().Info("Game Import successfully imported new game events")
}

func (s GameImportService) fetchMatchesForLeagueGroup(league string, apiKey string) ([]Match, error) {
	params := make(map[string]string)
	params["filters[leagues][]"] = league
	params["search"] = "skylarks"

	url := GetAPIURL("matches.json", params, apiKey)
	matches, _, err := FetchResource[[]Match](url.String())

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (s GameImportService) createOrUpdateEvents(matches []Match, teamID string) (err error) {
	for _, match := range matches {
		record, err := s.App.FindFirstRecordByData("events", "bsm_id", match.ID)

		// if not found, it throws an error, so create new record
		if err != nil {
			collection, err := s.App.FindCollectionByNameOrId("events")
			if err != nil {
				s.App.Logger().Error("Error fetching event collection: ", "error", err)
				continue
			}

			record = core.NewRecord(collection)
			err = s.setEventRecordValues(record, match, teamID)
			if err != nil {
				s.App.Logger().Error("Error setting event record: ", "error", err)
				continue
			}
		} else {
			// no error - update existing record
			err = s.setEventRecordValues(record, match, teamID)
			if err != nil {
				s.App.Logger().Error("Error setting event record: ", "error", err)
				continue
			}
		}

		s.App.Logger().Debug("Record data immediately before persist call", "record", record)
		if err := s.App.Save(record); err != nil {
			s.App.Logger().Error("Persisting event record failed", "error", err, "record", record, "teamID", teamID, "match", match)
			continue
		}
	}
	return
}

func (s GameImportService) setEventRecordValues(record *core.Record, match Match, teamID string) (err error) {
	starttime, err := types.ParseDateTime(match.Time)
	if err != nil {
		return err
	}
	endtime := starttime.Time().Add(time.Hour * 3)
	meetingtime := starttime.Time().Add(-2*time.Hour - 30*time.Minute)

	s.App.Logger().Debug("Current Date Values", "starttime", starttime, "endtime", endtime, "meetingtime", meetingtime)

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

	s.App.Logger().Debug("Record populated with fields", "record", record)

	return
}
