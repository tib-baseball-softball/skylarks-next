package dp

import (
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/push"
)

// TODO: localization for all push messages

func NotifyNewAnnouncement(e *core.RecordEvent, ps push.PushService) error {
	announcement := &base.Announcement{}
	announcement.SetProxyRecord(e.Record)

	if announcement.Club() == "" && announcement.Team() == "" {
		e.App.Logger().Error("Announcement has no club or team, cannot handle", "announcement", announcement)
		return e.Next()
	}

	var targetID string
	var target string

	if announcement.Club() != "" && announcement.Team() == "" {
		targetID = announcement.Club()
		target = base.ClubsCollection
	} else if announcement.Club() == "" && announcement.Team() != "" {
		targetID = announcement.Team()
		target = base.TeamsCollection
	}

	if targetID == "" {
		return e.Next()
	}

	if target == base.TeamsCollection {
		PushTeamOrClubAnnouncement(e.App, base.TeamsCollection, announcement, targetID, ps)
	} else {
		PushTeamOrClubAnnouncement(e.App, base.ClubsCollection, announcement, targetID, ps)
	}
	return e.Next()
}

type NamedRecordProxy interface {
	core.RecordProxy
	Name() string
}

// PushTeamOrClubAnnouncement sends a push notification to all subscribed users when a new announcement is created for a team or club.
func PushTeamOrClubAnnouncement(app core.App, coll string, announcement *base.Announcement, recordID string, ps push.PushService) {
	record, err := app.FindRecordById(coll, recordID)
	if record == nil {
		app.Logger().Error("Error fetching club record", "error", err, "teamID", recordID)
		return
	}

	var proxy NamedRecordProxy

	if coll == base.TeamsCollection {
		proxy = &base.Team{}
	} else {
		proxy = &base.Club{}
	}
	proxy.SetProxyRecord(record)

	subs, err := push.GetSubscriptionsForTeamOrClub(recordID, coll, app)
	if err != nil {
		app.Logger().Error("Error fetching subscriptions", "error", err, coll, recordID)
		return
	}

	for _, sub := range subs {
		app.Logger().Debug("Sending push notification to user", "user", sub.User)

		msg := &push.Message{
			Title: "New Announcement in " + proxy.Name(),
			Body:  announcement.Title(),
			Tag:   "team_club_new_announcement",
			Actions: []push.ActionData{
				{
					Action: push.PushActionLink,
					Title:  "Read Announcement",
					Type:   "button",
				},
			},
			Data: map[string]any{
				"navigate": RouteAnnouncementDetail(announcement.Id),
			},
		}
		err := ps.SendPushMessage(msg, new(sub.ToWebPushSubscription()))
		if err != nil {
			return
		}
	}
}
