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

func transformUserRecords(records []*core.Record) ([]User, error) {
	var users []User
	for _, record := range records {
		user := User{}
		user.SetProxyRecord(record)
		users = append(users, user)
	}
	return users, nil
}
