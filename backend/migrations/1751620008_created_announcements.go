package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `{
			"createRule": null,
			"deleteRule": null,
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
					"id": "text724990059",
					"max": 0,
					"min": 0,
					"name": "title",
					"pattern": "",
					"presentable": true,
					"primaryKey": false,
					"required": true,
					"system": false,
					"type": "text"
				},
				{
					"convertURLs": false,
					"hidden": false,
					"id": "editor4152030090",
					"maxSize": 0,
					"name": "bodytext",
					"presentable": true,
					"required": true,
					"system": false,
					"type": "editor"
				},
				{
					"exceptDomains": [],
					"hidden": false,
					"id": "url917281265",
					"name": "link",
					"onlyDomains": [],
					"presentable": false,
					"required": false,
					"system": false,
					"type": "url"
				},
				{
					"hidden": false,
					"id": "select1655102503",
					"maxSelect": 1,
					"name": "priority",
					"presentable": false,
					"required": true,
					"system": false,
					"type": "select",
					"values": [
						"info",
						"warning",
						"danger"
					]
				},
				{
					"cascadeDelete": true,
					"collectionId": "z806hmy5um3qy5x",
					"hidden": false,
					"id": "relation3102619762",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "club",
					"presentable": false,
					"required": false,
					"system": false,
					"type": "relation"
				},
				{
					"cascadeDelete": true,
					"collectionId": "fltkvrsbbvzf9cf",
					"hidden": false,
					"id": "relation3303056927",
					"maxSelect": 1,
					"minSelect": 0,
					"name": "team",
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
			"id": "pbc_3866499052",
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_yJqhZETWyh` + "`" + ` ON ` + "`" + `announcements` + "`" + ` (\n  ` + "`" + `club` + "`" + `,\n  ` + "`" + `team` + "`" + `\n)"
			],
			"listRule": null,
			"name": "announcements",
			"system": false,
			"type": "base",
			"updateRule": null,
			"viewRule": null
		}`

		collection := &core.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3866499052")
		if err != nil {
			return err
		}

		return app.Delete(collection)
	})
}
