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
			"updateRule": "@request.auth.id != \"\" && ((@request.body.team.id = \"\" && @request.body.club.admins.id ?= @request.auth.id) || (@request.body.club.id = \"\" && @request.body.team.admins.id ?= @request.auth.id ))"
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
			"updateRule": "@request.auth.id != \"\" && ((team.id = \"\" && club.admins.id ?= @request.auth.id) || (club.id = \"\" && team.admins.id ?= @request.auth.id ))"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
