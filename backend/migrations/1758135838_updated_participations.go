package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ngighvrbraklj58")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE UNIQUE INDEX `+"`"+`idx_YCrorbP`+"`"+` ON `+"`"+`participations`+"`"+` (\n  `+"`"+`user`+"`"+`,\n  `+"`"+`event`+"`"+`\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("ngighvrbraklj58")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX `+"`"+`idx_YCrorbP`+"`"+` ON `+"`"+`participations`+"`"+` (\n  `+"`"+`user`+"`"+`,\n  `+"`"+`event`+"`"+`\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
