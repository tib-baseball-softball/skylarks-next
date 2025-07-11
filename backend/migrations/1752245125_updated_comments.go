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
			"listRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)"
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
			"listRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id)",
			"viewRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
