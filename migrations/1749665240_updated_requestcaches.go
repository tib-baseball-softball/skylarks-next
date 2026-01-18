package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ap0vz28x8jh6gv6")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX `+"`"+`idx_8P8482G`+"`"+` ON `+"`"+`requestcaches`+"`"+` (`+"`"+`hash`+"`"+`)",
				"CREATE INDEX `+"`"+`idx_0ty87dm7lr`+"`"+` ON `+"`"+`requestcaches`+"`"+` (`+"`"+`identifier`+"`"+`)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text1999537002",
			"max": 0,
			"min": 0,
			"name": "identifier",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"exceptDomains": null,
			"hidden": false,
			"id": "d66krk0w",
			"name": "url",
			"onlyDomains": null,
			"presentable": true,
			"required": false,
			"system": false,
			"type": "url"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ap0vz28x8jh6gv6")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX `+"`"+`idx_8P8482G`+"`"+` ON `+"`"+`requestcaches`+"`"+` (`+"`"+`hash`+"`"+`)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text1999537002")

		// update field
		if err := collection.Fields.AddMarshaledJSONAt(3, []byte(`{
			"exceptDomains": null,
			"hidden": false,
			"id": "d66krk0w",
			"name": "url",
			"onlyDomains": null,
			"presentable": true,
			"required": true,
			"system": false,
			"type": "url"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
