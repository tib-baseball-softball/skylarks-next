package hooks

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"github.com/tib-baseball-softball/skylarks-next/model"
)

func AddEventParticipationData(app core.App, event *core.RecordEnrichEvent) error {
	event.Record.WithCustomData(true)
	err := enrichParticipationData(app, event.RequestInfo.Auth, event.Record)
	if err != nil {
		app.Logger().Error("EnrichParticipationData error", err)
		return err
	}

	return event.Next()
}

func enrichParticipationData(app core.App, user *core.Record, record *core.Record) error {
	err := addUserParticipation(app, user, record)
	if err != nil {
		return err
	}

	err = addParticipationsByType(app, record)
	if err != nil {
		return err
	}

	return nil
}

func addParticipationsByType(app core.App, record *core.Record) error {
	participations, err := app.FindRecordsByFilter(
		"participations",
		"event = {:eventID}",
		"",
		0,
		0,
		dbx.Params{"eventID": record.Id},
	)
	if err != nil {
		app.Logger().Error("EnrichParticipationData error", err)
		return err
	}
	if errs := app.ExpandRecords(participations, []string{"user"}, nil); len(errs) > 0 {
		return fmt.Errorf("failed to expand: %v", errs)
	}

	participationsByType := model.ParticipationsByType{
		In:    []*core.Record{},
		Out:   []*core.Record{},
		Maybe: []*core.Record{},
	}

	for _, participation := range participations {
		state, err := enum.ParseParticipationType(participation.GetString("state"))
		if err != nil {
			return err
		}
		switch state {
		case enum.In:
			participationsByType.In = append(participationsByType.In, participation)
		case enum.Out:
			participationsByType.Out = append(participationsByType.Out, participation)
		case enum.Maybe:
			participationsByType.Maybe = append(participationsByType.Maybe, participation)
		}
	}
	record.Set("participations", participationsByType)

	return nil
}

func addUserParticipation(app core.App, user *core.Record, record *core.Record) error {
	userParticipation, err := app.FindFirstRecordByFilter(
		"participations",
		"user = {:userID} && event = {:eventID}",
		dbx.Params{"userID": user.Id, "eventID": record.Id},
	)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// not found - that is perfectly valid, user has not participated in this event yet, so set to null in JSON
		record.Set("userParticipation", nil)
		return nil
	}
	if errs := app.ExpandRecord(userParticipation, []string{"user"}, nil); len(errs) > 0 {
		return fmt.Errorf("failed to expand: %v", errs)
	}

	record.Set("userParticipation", userParticipation)
	return nil
}
