package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3866499052")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": "@request.auth.id != \"\" && ((team.id = \"\" && club.admins.id ?= @request.auth.id) || (club.id = \"\" && team.admins.id ?= @request.auth.id ))",
			"deleteRule": "@request.auth.id != \"\" && ((team.id = \"\" && club.admins.id ?= @request.auth.id) || (club.id = \"\" && team.admins.id ?= @request.auth.id ))",
			"listRule": "@request.auth.id != \"\" && (club.id ?~ @request.auth.club.id || team.id ?~ @request.auth.teams.id)",
			"updateRule": "@request.auth.id != \"\" && ((team.id = \"\" && club.admins.id ?= @request.auth.id) || (club.id = \"\" && team.admins.id ?= @request.auth.id ))",
			"viewRule": "@request.auth.id != \"\" && (club.id ?~ @request.auth.club.id || team.id ?~ @request.auth.teams.id)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3866499052")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"createRule": null,
			"deleteRule": null,
			"listRule": null,
			"updateRule": null,
			"viewRule": null
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
