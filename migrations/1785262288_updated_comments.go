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
			"createRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id\n|| @request.auth.club.id ?= event.club.id\n|| @request.auth.club.id ?= event.additional_teams.club.id\n|| @request.auth.club.id ?= announcement.club.id \n|| event.team.club.admins.id ?= @request.auth.id  \n|| event.team.admins.id ?= @request.auth.id \n|| event.club.admins.id ?= @request.auth.id\n|| event.additional_teams.admins.id ?= @request.auth.id\n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id\n)",
			"deleteRule": "@request.auth.id != \"\" \n&& \n(\nuser.id = @request.auth.id \n|| event.club.admins.id ?= @request.auth.id\n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins ?= @request.auth.id \n|| event.additional_teams.admins.id ?= @request.auth.id\n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id\n)",
			"listRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id  \n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)",
			"updateRule": "@request.auth.id != \"\" \n&& \n(\n(@request.body.event:isset = false && @request.body.announcement:isset = false && user.id = @request.auth.id) \n|| @request.auth.club.id ?= @request.body.event.team.club.id\n|| @request.auth.club.id ?= @request.body.event.club.id\n|| @request.auth.teams.id ?= @request.body.event.team.id\n|| @request.auth.teams.id ?= @request.body.event.additional_teams.id  \n|| @request.auth.club.id ?= @request.body.announcement.club.id \n|| @request.body.event.team.club.admins.id ?= @request.auth.id  \n|| @request.body.event.team.admins.id ?= @request.auth.id \n|| @request.body.announcement.club.admins.id ?= @request.auth.id \n|| @request.body.announcement.team.admins.id ?= @request.auth.id\n)",
			"viewRule": "@request.auth.id != \"\" \n&& \n(\n@request.auth.club.id ?= event.team.club.id  \n|| @request.auth.club.id ?= announcement.club.id \n|| @request.auth.club.id ?= announcement.team.club.id \n|| @request.auth.teams.id ?= announcement.team.id \n|| announcement.club.admins.id ?= @request.auth.id \n|| announcement.team.admins.id ?= @request.auth.id \n|| event.team.admins.id ?= @request.auth.id \n|| event.team.club.admins.id ?= @request.auth.id\n)"
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
			"createRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || event.team.club.admins.id ?= @request.auth.id  || event.team.admins.id ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"deleteRule": "@request.auth.id != \"\" && (user.id = @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"listRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)",
			"updateRule": "@request.auth.id != \"\" && ((@request.body.event:isset = false &&  @request.body.announcement:isset = false && user.id = @request.auth.id) || @request.auth.club.id ?= @request.body.event.team.club.id  || @request.auth.club.id ?= @request.body.announcement.club.id || @request.body.event.team.club.admins.id ?= @request.auth.id  || @request.body.event.team.admins.id ?= @request.auth.id || @request.body.announcement.club.admins.id ?= @request.auth.id || @request.body.announcement.team.admins.id ?= @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
