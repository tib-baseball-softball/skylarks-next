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
		if err := collection.Fields.AddMarshaledJSONAt(7, []byte(`{
			"exceptDomains": null,
			"help": "",
			"hidden": false,
			"id": "email1203275662",
			"name": "admin_email",
			"onlyDomains": null,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "email"
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
		collection.Fields.RemoveById("email1203275662")

		return app.Save(collection)
	})
}
