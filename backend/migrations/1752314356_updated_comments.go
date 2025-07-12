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
			"updateRule": "@request.auth.id != \"\" && ((@request.body.event:isset = false &&  @request.body.announcement:isset = false && user.id = @request.auth.id) || @request.auth.club.id ?= @request.body.event.team.club.id  || @request.auth.club.id ?= @request.body.announcement.club.id || @request.body.event.team.club.admins.id ?= @request.auth.id  || @request.body.event.team.admins.id ?= @request.auth.id || @request.body.announcement.club.admins.id ?= @request.auth.id || @request.body.announcement.team.admins.id ?= @request.auth.id)"
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
			"updateRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= @request.body.event.team.club.id  || @request.auth.club.id ?= @request.body.announcement.club.id || @request.body.event.team.club.admins.id ?= @request.auth.id  || @request.body.event.team.admins.id ?= @request.auth.id || @request.body.announcement.club.admins.id ?= @request.auth.id || @request.body.announcement.team.admins.id ?= @request.auth.id)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
