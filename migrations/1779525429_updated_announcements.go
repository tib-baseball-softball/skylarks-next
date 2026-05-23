package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3866499052")
		if err != nil {
			return err
		}

		_, err = app.DB().
			NewQuery("UPDATE announcements SET TEMP_BODYTEXT_DATA = bodytext;").
			Execute()
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("editor4152030090")

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"autogeneratePattern": "",
			"help": "",
			"hidden": false,
			"id": "text4152030090",
			"max": 8000,
			"min": 0,
			"name": "bodytext",
			"pattern": "",
			"presentable": true,
			"primaryKey": false,
			"required": true,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3866499052")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"convertURLs": false,
			"help": "",
			"hidden": false,
			"id": "editor4152030090",
			"maxSize": 0,
			"name": "bodytext",
			"presentable": true,
			"required": true,
			"system": false,
			"type": "editor"
		}`)); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text4152030090")

		return app.Save(collection)
	})
}
