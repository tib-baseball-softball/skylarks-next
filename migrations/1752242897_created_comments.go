package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		// language=json
		jsonData := `{
			"createRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || event.team.club.admins.id ?= @request.auth.id  || event.team.admins.id ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"deleteRule": "@request.auth.id != \"\" && (user.id = @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins ?= @request.auth.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id)",
			"fields": [
				{
					"autogeneratePattern": "[a-z0-9]{15}",
					"hidden": false,
					"id": "text3208210256",
					"max": 15,
					"min": 15,
					"name": "id",
					"pattern": "^[a-z0-9]+$",
					"presentable": false,
					"primaryKey": true,
					"required": true,
					"system": true,
					"type": "text"
				},
				{
					"autogeneratePattern": "",
					"hidden": false,
					"id": "text999008199",
					"max": 500,
					"min": 0,
					"name": "text",
					"pattern": "",
					"presentable": false,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"cascadeDelete": true,
					"collectionId": "_pb_users_auth_",
					"hidden": false,
					"id": "relation2375276105",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "user",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "pbc_3866499052",
					"hidden": false,
					"id": "relation1304025372",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "announcement",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "zyst8ardf4onbwz",
					"hidden": false,
					"id": "relation1001261735",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "event",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"hidden": false,
					"id": "autodate2990389176",
					"name": "created",
					"onCreate": true,
					"onUpdate": false,
					"presentable": false,
					"system": false,
					"type": "autodate"
				},
				{
					"hidden": false,
					"id": "autodate3332085495",
					"name": "updated",
					"onCreate": true,
					"onUpdate": true,
					"presentable": false,
					"system": false,
					"type": "autodate"
				}
			],
			"id": "pbc_533777971",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_yMUhqMcso9` + "`" + ` ON ` + "`" + `comments` + "`" + ` (\n  ` + "`" + `user` + "`" + `,\n  ` + "`" + `announcement` + "`" + `,\n  ` + "`" + `event` + "`" + `\n)"
			],
			"listRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)",
			"name": "comments",
			"system": false,
			"type": "base",
			"updateRule": "@request.auth.id != \"\" && ((@request.body.event:isset = false &&  @request.body.announcement:isset = false && user.id = @request.auth.id) || @request.auth.club.id ?= @request.body.event.team.club.id  || @request.auth.club.id ?= @request.body.announcement.club.id || @request.body.event.team.club.admins.id ?= @request.auth.id  || @request.body.event.team.admins.id ?= @request.auth.id || @request.body.announcement.club.admins.id ?= @request.auth.id || @request.body.announcement.team.admins.id ?= @request.auth.id)",
			"viewRule": "@request.auth.id != \"\" && (@request.auth.club.id ?= event.team.club.id  || @request.auth.club.id ?= announcement.club.id || @request.auth.club.id ?= announcement.team.club.id || @request.auth.teams.id ?= announcement.team.id || announcement.club.admins.id ?= @request.auth.id || announcement.team.admins.id ?= @request.auth.id || event.team.admins.id ?= @request.auth.id || event.team.club.admins.id ?= @request.auth.id)"
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_533777971")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
