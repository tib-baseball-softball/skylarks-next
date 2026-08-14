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
			"createRule": "@request.auth.id != \"\"\n&&\n(\nteam.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"deleteRule": "@request.auth.id != \"\"\n&&\n(\nteam.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"listRule": "@request.auth.id != \"\"\n&&\n(\nteam.id ?= @request.auth.teams.id \n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"updateRule": "@request.auth.id != \"\"\n&&\n(\nteam.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\"\n&&\n(\nteam.id ?= @request.auth.teams.id \n|| additional_teams.id ?= @request.auth.teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)"
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
			"createRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"deleteRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"listRule": "@request.auth.id != \"\" && \n(team.id ?~ @request.auth.teams.id || team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"updateRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && \n(team.id ?~ @request.auth.teams.id || team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
