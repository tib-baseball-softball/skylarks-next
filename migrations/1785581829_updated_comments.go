package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_533777971")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" \n&& \nuser.id = @request.auth.id // admins can edit, but not create comments for others\n&&\n(\n@request.auth.club.id ?= event.team.club.id\n|| @request.auth.club.id ?= event.club.id\n|| @request.auth.club.id ?= event.additional_teams.club.id\n|| @request.auth.club.id ?= announcement.club.id \n|| event.team.club.admins.id ?= @request.auth.id  \n|| event.team.admins.id ?= @request.auth.id \n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id\n)",
			"listRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id\n|| @request.auth.club.id ?= event.club.id\n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id \n|| @request.auth.club.id ?= event.club.id\n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_533777971")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id\n|| @request.auth.club.id ?= event.club.id\n|| @request.auth.club.id ?= event.additional_teams.club.id\n|| @request.auth.club.id ?= announcement.club.id \n|| event.team.club.admins.id ?= @request.auth.id  \n|| event.team.admins.id ?= @request.auth.id \n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id\n)",
			"listRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id  \n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id  \n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
