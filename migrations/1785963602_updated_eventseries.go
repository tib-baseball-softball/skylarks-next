package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_29523945")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" && \n(\nteam.admins.id ?= @request.auth.id\n|| team.club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"deleteRule": "@request.auth.id != \"\" && \n(\nteam.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_QVU4dUhywn` + "`" + ` ON ` + "`" + `eventseries` + "`" + ` (` + "`" + `team` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_JwbMaKE47A` + "`" + ` ON ` + "`" + `eventseries` + "`" + ` (` + "`" + `series_end` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_w2tzr4um8d` + "`" + ` ON ` + "`" + `eventseries` + "`" + ` (` + "`" + `additional_teams` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\" &&\n(\n@request.auth.teams.id ?= team.id\n|| @request.auth.teams.id ?= additional_teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"updateRule": "@request.auth.id != \"\" && \n(\nteam.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\" &&\n(\n@request.auth.teams.id ?= team.id\n|| @request.auth.teams.id ?= additional_teams.id\n|| team.admins.id ?= @request.auth.id \n|| team.club.admins.id ?= @request.auth.id\n|| additional_teams.admins.id ?= @request.auth.id\n|| additional_teams.club.admins.id ?= @request.auth.id\n)"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"cascadeDelete": false,
			"collectionId": "fltkvrsbbvzf9cf",
			"help": "",
			"hidden": false,
			"id": "relation3929037161",
			"maxSelect": 10,
			"minSelect": 0,
			"name": "additional_teams",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_29523945")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"deleteRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_QVU4dUhywn` + "`" + ` ON ` + "`" + `eventseries` + "`" + ` (` + "`" + `team` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_JwbMaKE47A` + "`" + ` ON ` + "`" + `eventseries` + "`" + ` (` + "`" + `series_end` + "`" + `)"
			],
			"listRule": "@request.auth.id != \"\" && \n(team.id ?~ @request.auth.teams.id || team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"updateRule": "@request.auth.id != \"\" && \n(team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && \n(team.id ?~ @request.auth.teams.id || team.admins ?~ @request.auth.id || team.club.admins ?~ @request.auth.id)"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation3929037161")

		return app.Save(collection)
	})
}
