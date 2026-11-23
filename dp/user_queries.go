package dp

import (
	"ganigeorgiev.com/pocketbase/dbx"
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func GetUsersOnTeam(teamID string, app core.App) ([]base.User, error) {
	records, err := app.FindRecordsByFilter(
		base.UserCollection,
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

func GetUsersOnClub(clubID string, app core.App) ([]base.User, error) {
	records, err := app.FindRecordsByFilter(
		base.UserCollection,
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

func GetAdminsForTeam(team base.Team, app core.App) ([]base.User, error) {
	records, err := app.FindRecordsByIds(base.UserCollection, team.Admins())
	if err != nil {
		return nil, err
	}
	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetAdminsForClub(club base.Club, app core.App) ([]base.User, error) {
	records, err := app.FindRecordsByIds(base.UserCollection, club.Admins())
	if err != nil {
		return nil, err
	}
	users, err := transformUserRecords(records)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func transformUserRecords(records []*core.Record) ([]base.User, error) {
	var users []base.User
	for _, record := range records {
		user := base.User{}
		user.SetProxyRecord(record)
		users = append(users, user)
	}
	return users, nil
}
