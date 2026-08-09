package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("zyst8ardf4onbwz")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.id != \"\"\n&&\n(\n@request.auth.teams.id ?= team.id\n|| @request.auth.club.id ?= club.id \n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\"\n&&\n(\n@request.auth.teams.id ?= team.id\n|| @request.auth.club.id ?= club.id \n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("zyst8ardf4onbwz")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"listRule": "@request.auth.id != \"\"\n&&\n(\nteam.id ?= @request.auth.teams.id\n|| club.id ?= @request.auth.club.id\n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\"\n&&\n(\nteam.id ?= @request.auth.teams.id\n|| club.id ?= @request.auth.club.id\n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
