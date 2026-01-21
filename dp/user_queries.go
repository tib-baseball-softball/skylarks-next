package dp

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func GetUsersOnTeam(teamID string, app core.App) ([]User, error) {
	records, err := app.FindRecordsByFilter(
		UserCollection,
		"teams ?= {:team}",
		"",
		0,
		0,
		dbx.Params{"team": teamID},
	)
	if err != nil {
		return nil, err
	}

	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersOnClub(clubID string, app core.App) ([]User, error) {
	records, err := app.FindRecordsByFilter(
		UserCollection,
		"club ?= {:club}",
		"",
		0,
		0,
		dbx.Params{"club": clubID},
	)
	if err != nil {
		return nil, err
	}

	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAdminsForTeam(team Team, app core.App) ([]User, error) {
	records, err := app.FindRecordsByIds(UserCollection, team.Admins())
	if err != nil {
		return nil, err
	}
	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAdminsForClub(club Club, app core.App) ([]User, error) {
	records, err := app.FindRecordsByIds(UserCollection, club.Admins())
	if err != nil {
		return nil, err
	}
	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetSubscriptionsForUserIDs(userIDs []string, app core.App) ([]PushSubscription, error) {
	ids := make([]interface{}, len(userIDs))
	records := make([]*core.Record, len(ids))

	err := app.RecordQuery(PushSubscriptionsCollection).
		AndWhere(dbx.In("user", ids...)).
		All(&records)
	if err != nil {
		return nil, err
	}

	subs := make([]PushSubscription, len(records))
	for _, record := range records {
		sub := PushSubscription{}
		sub.SetProxyRecord(record)
		subs = append(subs, sub)
	}
	return subs, nil
}

func transformUserRecords(records []*core.Record) ([]User, error) {
	var users []User
	for _, record := range records {
		user := User{}
		user.SetProxyRecord(record)
		users = append(users, user)
	}
	return users, nil
}
