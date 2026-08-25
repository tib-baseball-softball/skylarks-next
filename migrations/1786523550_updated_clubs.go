package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("z806hmy5um3qy5x")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"autogeneratePattern": "",
			"help": "",
			"hidden": false,
			"id": "text522183215",
			"max": 0,
			"min": 0,
			"name": "bsm_search_string",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"autogeneratePattern": "",
			"help": "",
			"hidden": false,
			"id": "text2411891325",
			"max": 0,
			"min": 0,
			"name": "team_name",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("z806hmy5um3qy5x")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text522183215")

		// remove field
		collection.Fields.RemoveById("text2411891325")

		return app.Save(collection)
	})
}
