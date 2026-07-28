package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ngighvrbraklj58")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" \n&& \n(\nuser.teams.id ?= event.team.id\n|| user.teams.id ?= event.additional_teams.id\n|| user.club.id ?= event.club.id\n)\n&& \n(\nuser.id = @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n) ",
			"deleteRule": "@request.auth.id != \"\" \n&& \n(\nevent.team.admins.id ?= @request.auth.id \n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n|| event.team.club.admins.id ?= @request.auth.id\n)",
			"listRule": "@request.auth.id != \"\" \n&&\n(\nevent.team.id ?= @request.auth.teams.id\n|| event.club.id ?= @request.auth.club.id\n|| event.additional_teams.id ?= @request.auth.teams.id\n)",
			"updateRule": "@request.auth.id != \"\" \n&& \n(\nuser.teams.id ?= event.team.id\n|| user.teams.id ?= event.additional_teams.id\n|| user.club.id ?= event.club.id\n)\n&& \n(\nuser.id = @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n) ",
			"viewRule": "@request.auth.id != \"\" \n&&\n(\nevent.team.id ?= @request.auth.teams.id\n|| event.club.id ?= @request.auth.club.id\n|| event.additional_teams.id ?= @request.auth.teams.id\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ngighvrbraklj58")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" && user.teams.id ?= event.team.id && (user.id = @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id) ",
			"deleteRule": "@request.auth.id != \"\" && (event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)",
			"listRule": "@request.auth.id != \"\" && event.team.id ?= @request.auth.teams.id",
			"updateRule": "@request.auth.id != \"\" && (user.id = @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && event.team.id ?= @request.auth.teams.id"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
