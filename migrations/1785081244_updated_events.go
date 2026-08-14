package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("zyst8ardf4onbwz")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(10, []byte(`{
			"cascadeDelete": false,
			"collectionId": "z806hmy5um3qy5x",
			"help": "event is valid for a whole club, not just a single team",
			"hidden": false,
			"id": "relation3102619762",
			"maxSelect": 0,
			"minSelect": 0,
			"name": "club",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"cascadeDelete": false,
			"collectionId": "fltkvrsbbvzf9cf",
			"help": "additional teams that this event is valid for beyond the main team",
			"hidden": false,
			"id": "relation3929037161",
			"maxSelect": 999,
			"minSelect": 0,
			"name": "additional_teams",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("zyst8ardf4onbwz")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation3102619762")

		// remove field
		collection.Fields.RemoveById("relation3929037161")

		return app.Save(collection)
	})
}
