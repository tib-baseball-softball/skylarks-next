package dp

import "github.com/pocketbase/pocketbase/core"

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
		err := ps.SendPushMessage(msg, new(sub.ToWebPushSubscription()))
		if err != nil {
			return
		}
	}

	return
}
