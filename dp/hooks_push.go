package dp

import (
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TODO: localization for all push messages

func NotifyNewAnnouncement(e *core.RecordEvent, ps PushService) error {
	announcement := &Announcement{}
	announcement.SetProxyRecord(e.Record)

	if announcement.Club() == "" && announcement.Team() == "" {
		e.App.Logger().Error("Announcement has no club or team, cannot handle", "announcement", announcement)
		return e.Next()
	}

	var targetID string
	var target string

	if announcement.Club() != "" && announcement.Team() == "" {
		targetID = announcement.Club()
		target = ClubsCollection
	} else if announcement.Club() == "" && announcement.Team() != "" {
		targetID = announcement.Team()
		target = TeamsCollection
	}

	if targetID == "" {
		return e.Next()
	}

	if target == TeamsCollection {
		PushTeamOrClubAnnouncement(e.App, TeamsCollection, announcement, targetID, ps)
	} else {
		PushTeamOrClubAnnouncement(e.App, ClubsCollection, announcement, targetID, ps)
	}
	return e.Next()
}

// NotifyNewEvent sends a push notification to all subscribed team members when a new event is created.
func NotifyNewEvent(e *core.RecordEvent, ps PushService) error {
	event := &Event{}
	event.SetProxyRecord(e.Record)

	if errs := e.App.ExpandRecord(event.Record, []string{"location"}, nil); len(errs) > 0 {
		e.App.Logger().Warn("failed to expand:", "errors", errs, "event", event, "file", "hooks_push.go")
		return e.Next()
	}
	locationRecord := event.ExpandedOne("location")
	if locationRecord == nil {
		e.App.Logger().Error("Could not load location data for API key")
		return e.Next()
	}
	location := &Location{}
	location.SetProxyRecord(locationRecord)

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

	for _, sub := range subs {
		msg := &PushMessage{
			Title: "New Event in " + team.Name(),
			Body:  event.Title() + " (" + cases.Title(language.English, cases.Compact).String(event.Type()) + ") - starts at " + eventStartInAppTimeZone.Time().Format(time.RFC1123) + " at " + location.Name() + " (" + location.AddressAddon() + ")",
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

type NamedRecordProxy interface {
	core.RecordProxy
	Name() string
}

// PushTeamOrClubAnnouncement sends a push notification to all subscribed users when a new announcement is created for a team or club.
func PushTeamOrClubAnnouncement(app core.App, coll string, announcement *Announcement, recordID string, ps PushService) {
	record, err := app.FindRecordById(coll, recordID)
	if record == nil {
		app.Logger().Error("Error fetching club record", "error", err, "teamID", recordID)
		return
	}

	var proxy NamedRecordProxy

	if coll == TeamsCollection {
		proxy = &Team{}
	} else {
		proxy = &Club{}
	}
	proxy.SetProxyRecord(record)

	subs, err := GetSubscriptionsForTeamOrClub(recordID, coll, app)
	if err != nil {
		app.Logger().Error("Error fetching subscriptions", "error", err, coll, recordID)
		return
	}

	for _, sub := range subs {
		app.Logger().Debug("Sending push notification to user", "user", sub.User)

		msg := &PushMessage{
			Title: "New Announcement in " + proxy.Name(),
			Body:  announcement.Title(),
			Tag:   "team_club_new_announcement",
		}
		ws := sub.ToWebPushSubscription()

		err := ps.SendPushMessage(msg, &ws)
		if err != nil {
			return
		}
	}

	return
}
