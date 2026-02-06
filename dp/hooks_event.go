package dp

import (
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

// NotifyNewEvent sends a push notification to all subscribed team members when a new event is created.
func NotifyNewEvent(e *core.RecordEvent, ps PushService) error {
	event := &Event{}
	event.SetProxyRecord(e.Record)

	// if an event is part of a series, don't flood users with notifications as the series will handle this
	if event.Series() != "" {
		return e.Next()
	}

	if errs := e.App.ExpandRecord(event.Record, []string{"location"}, nil); len(errs) > 0 {
		e.App.Logger().Warn("failed to expand:", "errors", errs, "event", event, "file", "hooks_push.go")
		return e.Next()
	}

	team := &Team{}
	teamRecord, err := e.App.FindRecordById(TeamsCollection, event.Team())
	if err != nil {
		e.App.Logger().Warn("Failed to load team record for event, data corrupted", "teamID", event.Team(), "error", err, "file", "hooks_push.go")
		return e.Next()
	}
	team.SetProxyRecord(teamRecord)

	subs, err := GetSubscriptionsForTeamOrClub(team.Id, TeamsCollection, e.App)
	if err != nil {
		e.App.Logger().Warn("Error fetching subscriptions", "error", err, "teamID", teamRecord.Id, "file", "hooks_push.go")
		return e.Next()
	}

	timeLocation, err := LoadAppTimeZone()
	if err != nil {
		e.App.Logger().Error("Failed to load app timezone", "error", err, "file", "hooks_push.go")
		return e.Next()
	}

	eventStartInAppTimeZone, err := types.ParseDateTime(event.StartTime().Time().In(timeLocation))
	if err != nil {
		e.App.Logger().Error("Failed to parse event start time", "error", err, "file", "hooks_push.go")
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
		}
		ws := sub.ToWebPushSubscription()

		err := ps.SendPushMessage(msg, &ws)
		if err != nil {
			e.App.Logger().Warn("Error sending push notification", "error", err, "sub", sub, "teamID", team.Id)
			return e.Next()
		}
	}

	return e.Next()
}
