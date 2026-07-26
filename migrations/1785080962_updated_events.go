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

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"cascadeDelete": true,
			"collectionId": "fltkvrsbbvzf9cf",
			"help": "main team for this event",
			"hidden": false,
			"id": "prkespsn",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "team",
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

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"cascadeDelete": true,
			"collectionId": "fltkvrsbbvzf9cf",
			"help": "",
			"hidden": false,
			"id": "prkespsn",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "team",
			"presentable": false,
			"required": true,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
