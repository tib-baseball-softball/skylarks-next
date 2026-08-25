package dp

import (
	"errors"
	"slices"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ValidateEventTimes ensures that event times are valid and within the expected range.
func ValidateEventTimes(e *core.RecordRequestEvent) error {
	startTime := e.Record.GetDateTime("starttime")
	endTime := e.Record.GetDateTime("endtime")
	meetingTime := e.Record.GetDateTime("meetingtime")

	if !endTime.IsZero() && startTime.After(endTime) {
		return e.BadRequestError("Event start time cannot be after end time", nil)
	}

	if !endTime.IsZero() && meetingTime.After(endTime) {
		return e.BadRequestError("Event meeting time cannot be after end time", nil)
	}

	return e.Next()
}

func ValidateEvent(e *core.RecordEvent) error {
	event := &Event{}
	event.SetProxyRecord(e.Record)

	eventClubID := event.Club()
	eventTeamID := event.Team()
	additionalTeams := event.AdditionalTeams()

	if eventClubID == "" && eventTeamID == "" {
		err := LogicError{"An event needs to have either a club or a primary team set"}
		e.App.Logger().Error("event failed validation", "err", err, "event", event)
		return err
	}

	if eventClubID != "" && (eventTeamID != "" || event.HasAdditionalTeams()) {
		err := LogicError{"An event cannot both be club-wide and team-scoped"}
		e.App.Logger().Error("event failed validation", "err", err, "event", event)
		return err
	}

	if eventTeamID == "" && event.HasAdditionalTeams() {
		err := LogicError{"An event cannot have additional teams without a primary team"}
		e.App.Logger().Error("event failed validation", "err", err, "event", event)
		return err
	}

	if slices.Contains(additionalTeams, eventTeamID) {
		err := LogicError{"Cannot assign a team to an event both as primary and additional team"}
		e.App.Logger().Error("event failed validation", "err", err, "event", event)
		return err
	}

	// from here: additional teams validations and sanity checks
	if !event.HasAdditionalTeams() || eventClubID != "" {
		return e.Next()
	}

	if errs := e.App.ExpandRecord(event.Record, []string{"team", "additional_teams"}, nil); len(errs) > 0 {
		e.App.Logger().Error("failed to expand:", "errors", errs, "event", event, "file", "hooks_event.go")
		return errors.New("Couldn't validate event")
	}

	primaryTeam := &Team{}
	primaryTeamRecord := event.ExpandedOne("team")
	if primaryTeamRecord != nil {
		primaryTeam.SetProxyRecord(primaryTeamRecord)
	}

	additionalTeamRecords := event.ExpandedAll("additional_teams")
	for _, record := range additionalTeamRecords {
		team := &Team{}
		team.SetProxyRecord(record)

		if primaryTeam.Club() != team.Club() {
			err := LogicError{"Additional teams must belong to the same club as the primary team"}
			e.App.Logger().Error("event failed validation", "err", err, "event", event)
			return err
		}
	}

	return e.Next()
}

// NotifyNewEvent sends a push notification to all subscribed team members when a new event is created.
func NotifyNewEvent(e *core.RecordEvent, ps PushService) error {
	event := &Event{}
	event.SetProxyRecord(e.Record)

	// if an event is part of a series, don't flood users with notifications as the series will handle this
	if event.Series() != "" {
		return e.Next()
	}

	if errs := e.App.ExpandRecord(event.Record, []string{"location"}, nil); len(errs) > 0 {
		e.App.Logger().Warn("failed to expand:", "errors", errs, "event", event, "file", "hooks_event.go")
		return e.Next()
	}

	team := &Team{}
	teamRecord, err := e.App.FindRecordById(TeamsCollection, event.Team())
	if err != nil {
		e.App.Logger().Warn("Failed to load team record for event, data corrupted", "teamID", event.Team(), "error", err, "file", "hooks_event.go")
		return e.Next()
	}
	team.SetProxyRecord(teamRecord)

	subs, err := GetSubscriptionsForTeamOrClub(team.Id, TeamsCollection, e.App)
	if err != nil {
		e.App.Logger().Warn("Error fetching subscriptions", "error", err, "teamID", teamRecord.Id, "file", "hooks_event.go")
		return e.Next()
	}

	timeLocation, err := LoadAppTimeZone()
	if err != nil {
		e.App.Logger().Error("Failed to load app timezone", "error", err, "file", "hooks_event.go")
		return e.Next()
	}

	eventStartInAppTimeZone, err := types.ParseDateTime(event.StartTime().Time().In(timeLocation))
	if err != nil {
		e.App.Logger().Error("Failed to parse event start time", "error", err, "file", "hooks_event.go")
		return e.Next()
	}

	messageBody := event.Title() + " (" + cases.Title(language.English, cases.Compact).String(event.Type()) + ") - starts at " + eventStartInAppTimeZone.Time().Format(time.RFC1123)

	locationRecord := event.ExpandedOne("location")
	if locationRecord != nil {
		location := &Location{}
		location.SetProxyRecord(locationRecord)
		messageBody += " at " + location.Name() + " (" + location.AddressAddon() + ")"
	}

	for _, sub := range subs {
		msg := &PushMessage{
			Title: "New Event in " + team.Name(),
			Body:  messageBody,
			Tag:   "team_new_event",
			Actions: []PushActionData{
				{
					Action: PushActionLink,
					Title:  "Go to Event",
					Type:   "button",
				},
			},
			Data: map[string]any{
				"navigate": "/account/event/" + event.Id, // TODO: centralized place for frontend routes in backend context
			},
		}
		err := ps.SendPushMessage(msg, new(sub.ToWebPushSubscription()))
		if err != nil {
			e.App.Logger().Warn("Error sending push notification", "error", err, "sub", sub, "teamID", team.Id)
			return e.Next()
		}
	}

	return e.Next()
}
