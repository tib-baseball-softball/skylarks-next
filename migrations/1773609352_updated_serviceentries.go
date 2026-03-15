package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3454477543")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": "@request.auth.id != \"\"\n&& @request.body.club:changed = false\n&& (@request.body.club.admins.id ?= @request.body.board_member.id || @request.body.board_member:changed = false)\n&& (club.admins.id ?= @request.auth.id // club admins may edit\n  || member.id = @request.auth.id) // or you can edit your own"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3454477543")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": "@request.auth.id != \"\"\n&& (club.admins.id ?= board_member.id || @request.body.board_member:isset = false)\n&& (club.admins.id ?= @request.auth.id // club admins may edit\n  || member.id = @request.auth.id) // or you can edit your own"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
