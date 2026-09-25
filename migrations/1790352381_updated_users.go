package migrations

import (
	"encoding/json/v2"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": "@request.auth.id != \"\"\n&&\n// user editing themselves, no setting team and club info \n(\n(id = @request.auth.id \n&& @request.body.club:isset = false \n&& @request.body.teams:isset = false\n) \n\n// editing user is Club admin\n|| \n(club.admins.id ?= @request.auth.id \n&& \n// club admins can only add users to their own club\n(\n@request.body.club.admins.id ?= @request.auth.id \n|| @request.body.club:isset = false\n)\n) \n|| // editing user is Team admin for a team the user is a member of\n    (@request.body.club:changed = false\n      && \n\t\t(\n\t\t (club.id ?= @request.body.teams.club.id && club.id ?= teams.club.id)\n\t\t || \n\t\t (@request.body.teams:length = 0 || teams:length = 0)\n\t\t)\n      &&\n      (@request.body.teams.admins.id ?= @request.auth.id \n        || teams.admins.id ?= @request.auth.id\n      ) // team admins can only add users to their own teams\n    )\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"updateRule": "@request.auth.id != \"\"\n&&\n// user editing themselves, no setting team and club info \n(\n(id = @request.auth.id \n&& @request.body.club:isset = false \n&& @request.body.teams:isset = false\n) \n\n// editing user is Club admin\n|| \n(club.admins.id ?= @request.auth.id \n&& \n// club admins can only add users to their own club\n(\n@request.body.club.admins.id ?= @request.auth.id \n|| @request.body.club:isset = false\n)\n) \n|| // editing user is Team admin for a team the user is a member of\n    (@request.body.club:changed = false\n      && \n\t\t(\n\t\t (club.id ?= @request.body.teams.club.id && club.id ?= teams.club.id)\n\t\t || @request.body.teams:length = 0\n\t\t)\n      &&\n      (@request.body.teams.admins.id ?= @request.auth.id \n        || teams.admins.id ?= @request.auth.id\n      ) // team admins can only add users to their own teams\n    )\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
