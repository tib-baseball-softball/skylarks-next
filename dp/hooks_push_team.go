package dp

import "github.com/pocketbase/pocketbase/core"

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
		_ = PushTeamAnnouncement(e.App, announcement, targetID, ps)
	} else {
		//_ = PushClubAnnouncement(e.App, targetID, ps)
	}
	return e.Next()
}

func PushTeamAnnouncement(app core.App, announcement *Announcement, teamID string, ps PushService) error {
	teamMembers, err := GetUsersOnTeam(teamID, app)
	if err != nil {
		app.Logger().Error("Error fetching team members", "error", err, "team", teamID)
		return err
	}

	teamRecord, err := app.FindRecordById(TeamsCollection, teamID)
	if teamRecord == nil {
		app.Logger().Error("Error fetching team record", "error", err, "teamID", teamID)
		return err
	}
	team := &Team{}
	team.SetProxyRecord(teamRecord)

	userIDs := make([]string, len(teamMembers))
	for i, member := range teamMembers {
		userIDs[i] = member.Id
	}

	subs, err := GetSubscriptionsForUserIDs(userIDs, app)
	if err != nil {
		app.Logger().Error("Error fetching subscriptions", "error", err, "team", teamID)
		return err
	}
	for _, sub := range subs {
		app.Logger().Debug("Sending push notification to user", "user", sub.User)

		msg := &PushMessage{
			Title: "New Announcement in " + team.Name(),
			Body:  announcement.Title(),
			Tag:   "team_club_new_announcement",
		}
		ws := sub.ToWebPushSubscription()

		err := ps.SendPushMessage(msg, &ws)
		if err != nil {
			return err
		}
	}

	return nil
}
