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
			"createRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || event.team.club.admins.id ?= @request.auth.id  || event.team.admins.id ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"deleteRule": "@request.auth.id != \"\" && (user.id = @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"listRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id)",
			"updateRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= @request.body.event.team.club.id  || @request.auth.club.id ?= @request.body.announcement.club.id || @request.body.event.team.club.admins.id ?= @request.auth.id  || @request.body.event.team.admins.id ?= @request.auth.id || @request.body.announcement.club.admins.id ?= @request.auth.id || @request.body.announcement.team.admins.id ?= @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id)"
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
			"createRule": null,
			"deleteRule": null,
			"listRule": "",
			"updateRule": null,
			"viewRule": null
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
