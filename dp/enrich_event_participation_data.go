package dp

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func AddEventParticipationData(app core.App, event *core.RecordEnrichEvent) error {
	event.Record.WithCustomData(true)

	eventItem := &Event{}
	eventItem.SetProxyRecord(event.Record)

	err := enrichParticipationData(app, event.RequestInfo.Auth, eventItem)
	if err != nil {
		app.Logger().Error("EnrichParticipationData error", "error", err)
		return err
	}

	return event.Next()
}

func enrichParticipationData(app core.App, user *core.Record, event *Event) error {
	err := addUserParticipation(app, user, event)
	if err != nil {
		return err
	}

	err = addParticipationsByType(app, event)
	if err != nil {
		return err
	}

	return nil
}

func addParticipationsByType(app core.App, event *Event) error {
	participations, err := app.FindRecordsByFilter(
		ParticipationsCollection,
		"event = {:eventID}",
		"",
		0,
		0,
		dbx.Params{"eventID": event.Id},
	)
	if err != nil {
		app.Logger().Error("EnrichParticipationData error", "error", err)
		return err
	}
	if errs := app.ExpandRecords(participations, []string{"user"}, nil); len(errs) > 0 {
		return fmt.Errorf("failed to expand: %v", errs)
	}

	unspecifiedUsers, err := getUserDTOsWithoutParticipations(app, event.Id, event.Team())
	if err != nil {
		app.Logger().Error("EnrichParticipationData error", "error", err, "event", event, "team", event.Team())
	}

	participationsByType := ParticipationsByType{
		In:          []*core.Record{},
		Out:         []*core.Record{},
		Maybe:       []*core.Record{},
		Unspecified: unspecifiedUsers,
	}

	for _, participation := range participations {
		state, err := ParseParticipationType(participation.GetString("state"))
		if err != nil {
			return err
		}
		switch state {
		case In:
			participationsByType.In = append(participationsByType.In, participation)
		case Out:
			participationsByType.Out = append(participationsByType.Out, participation)
		case Maybe:
			participationsByType.Maybe = append(participationsByType.Maybe, participation)
		}
	}
	event.SetParticipations(participationsByType)

	return nil
}

func addUserParticipation(app core.App, user *core.Record, event *Event) error {
	userParticipation, err := app.FindFirstRecordByFilter(
		"participations",
		"user = {:userID} && event = {:eventID}",
		dbx.Params{"userID": user.Id, "eventID": event.Id},
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// not found - that is perfectly valid, user has not 
		// participated in this event yet, so set to null in JSON
		event.SetUserParticipation(nil)
		return nil
	}
	if errs := app.ExpandRecord(userParticipation, []string{"user"}, nil); len(errs) > 0 {
		return fmt.Errorf("failed to expand: %v", errs)
	}

	event.SetUserParticipation(userParticipation)
	return nil
}

func getUserDTOsWithoutParticipations(app core.App, eventID string, teamID string) ([]UserParticipationDTO, error) {
	participationDTOS := []UserParticipationDTO{}

	err := app.DB().
		NewQuery(`
			SELECT id, first_name, last_name, display_name
			FROM users
			WHERE EXISTS (SELECT 1 FROM json_each(users.teams) WHERE value = {:teamID})
			AND (SELECT id FROM participations WHERE event = {:eventID} AND user = users.id) IS NULL`).
		Bind(dbx.Params{"eventID": eventID, "teamID": teamID}).
		All(&participationDTOS)

	if err != nil {
		app.Logger().Error("getUserDTOsWithoutParticipations error", "error", err, "eventID", eventID)
		return nil, err
	}
	return participationDTOS, nil
}
