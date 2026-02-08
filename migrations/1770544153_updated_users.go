package migrations

import (
	"encoding/json"

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
			"updateRule": "@request.auth.id != \"\" \n\n&& \n\n(\n  (id = @request.auth.id && @request.body.club:isset = false && @request.body.teams:isset = false) // user editing themselves, no setting team and club info \n  \n|| // editing user is Club admin\n  \n    (club.admins.id ?= @request.auth.id \n      && \n    (@request.body.club.admins.id ?= @request.auth.id || @request.body.club:isset = false)\n    ) // club admins can only add users to their own club\n  \n|| // editing user is Team admin for a team the user is a member of\n    (@request.body.club:isset = false\n      && (club.id ?= @request.body.teams.club)\n      &&\n      (@request.body.teams.admins.id ?= @request.auth.id \n        || teams.admins.id ?= @request.auth.id\n      ) // team admins can only add users to their own teams\n    )\n)"
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
			"updateRule": "@request.auth.id != \"\" \n\n&& \n\n(\n  (id = @request.auth.id && @request.body.club:isset = false && @request.body.teams:isset = false) // user editing themselves, no setting team and club info \n  \n|| // editing user is Club admin\n  \n    (club.admins.id ?= @request.auth.id \n      && \n    (@request.body.club.admins.id ?= @request.auth.id || @request.body.club:isset = false)\n    ) // club admins can only add users to their own club\n  \n|| // editing user is Team admin for a team the user is a member of\n    (@request.body.club:isset = false\n      &&\n    (@request.body.teams.admins.id ?= @request.auth.id || teams.admins.id ?= @request.auth.id || @request.body.teams:isset = false) // team admins can only add users to their own teams\n    )\n)"
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
