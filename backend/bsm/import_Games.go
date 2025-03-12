package bsm

import (
	"encoding/json"
	"errors"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/internal/pb"
	"log"
	"log/slog"
	"sync"
	"time"
)

type GameImportServiceApp interface {
	FindRecordsByFilter(
		collectionModelOrIdentifier any,
		filter string,
		sort string,
		limit int,
		offset int,
		params ...dbx.Params,
	) ([]*core.Record, error)
	FindFirstRecordByData(collectionModelOrIdentifier any, key string, value any) (*core.Record, error)
	Logger() *slog.Logger
	ExpandRecord(record *core.Record, expands []string, optFetchFunc core.ExpandFetchFunc) map[string]error
	FindCollectionByNameOrId(nameOrId string) (*core.Collection, error)
	Save(model core.Model) error
}

type GameImportService struct {
	App GameImportServiceApp
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
			leagueGroup, err := s.App.FindFirstRecordByData(LeagueGroupsCollection, "bsm_id", team.GetInt("bsm_league_group"))
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

			err = s.createOrUpdateEvents(matches, team)
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
	params[LeagueFilter] = league
	params[SearchFilter] = "skylarks"

	url := GetAPIURL("matches.json", params, apiKey)
	matches, _, err := FetchResource[[]Match](url.String())

	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (s GameImportService) createOrUpdateEvents(matches []Match, team *core.Record) (err error) {
	for _, match := range matches {
		record, foundErr := s.App.FindFirstRecordByData("events", "bsm_id", match.ID)

		location, err := s.createOrUpdateField(team, match.Field)
		if err != nil {
			s.App.Logger().Error("Error creating or updating event: ", "error", err)
		}

		// if not found, it throws an error, so create new record
		if foundErr != nil {
			collection, err := s.App.FindCollectionByNameOrId("events")
			if err != nil {
				s.App.Logger().Error("Error fetching event collection: ", "error", err)
				continue
			}

			record = core.NewRecord(collection)
		}
		if record == nil {
			s.App.Logger().Error("event record was unexpectedly nil: ", "error", err, "match", match)
			continue
		}
		// no error - update existing record
		err = s.setEventRecordValues(record, match, team.Id, location)
		if err != nil {
			s.App.Logger().Error("Error setting event record: ", "error", err)
			continue
		}

		s.App.Logger().Debug("Record data immediately before persist call", "record", record)
		if err := s.App.Save(record); err != nil {
			s.App.Logger().Error("Persisting event record failed", "error", err, "record", record, "teamID", team.Id, "match", match)
			continue
		}
	}
	return
}

func (s GameImportService) setEventRecordValues(record *core.Record, match Match, teamID string, location *Location) (err error) {
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
	record.Set("location", location.Id)

	s.App.Logger().Debug("Record populated with fields", "record", record)

	return
}

// createOrUpdateField Adds field/location to database that can then be set to new events as well as selected for
// manually created events in frontend.
func (s GameImportService) createOrUpdateField(team *core.Record, field Field) (*Location, error) {
	record, err := s.App.FindFirstRecordByData(pb.LocationCollection, "bsm_id", field.BSMID)

	if err != nil {
		collection, err := s.App.FindCollectionByNameOrId(pb.LocationCollection)
		if err != nil {
			s.App.Logger().Error("Error fetching location collection: ", "error", err)
			return nil, err
		}

		if collection != nil {
			record = core.NewRecord(collection)
		} else {
			return nil, errors.New("cannot create new record, collection is nil")
		}
	}
	location := &Location{}
	location.SetProxyRecord(record)

	location.SetBSMID(field.BSMID)
	location.SetName(field.Name)
	location.SetDescription(field.Description)
	location.SetAddressAddon(field.AddressAddon)
	location.SetCity(field.City)
	location.SetStreet(field.Street)
	location.SetPostalCode(field.PostalCode)
	location.SetCountry(field.Country)
	location.SetLatitude(field.Latitude)
	location.SetLongitude(field.Longitude)
	location.SetSpectatorTotal(field.SpectatorTotal)
	location.SetSpectatorSeats(field.SpectatorSeats)
	location.SetOtherInformation(field.OtherInformation)
	location.SetGroundRules(field.Groundrules)
	location.SetHumanCountry(field.HumanCountry)
	location.SetPhotoURL(field.PhotoURL)

	location.SetClub(team.GetString("club"))

	if err := s.App.Save(location); err != nil {
		s.App.Logger().Error("Persisting location failed", "error", err, "location", location)
		return nil, err
	}
	return location, nil
}
