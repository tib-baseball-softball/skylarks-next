package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("zyst8ardf4onbwz")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_zuKuzAF` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `bsm_id` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_8ixkwcxr4q` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `type` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_miocwua0bp` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `team` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_snxc3kufmb` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `series` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_0c501exuc1` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `location` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_kji3y3rlcp` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `prev` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_ytoiuzjffa` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `next` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_38rk28bmkd` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `team` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(15, []byte(`{
			"cascadeDelete": false,
			"collectionId": "zyst8ardf4onbwz",
			"help": "previous event in a series",
			"hidden": false,
			"id": "relation3168962645",
			"maxSelect": 0,
			"minSelect": 0,
			"name": "prev",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"cascadeDelete": false,
			"collectionId": "zyst8ardf4onbwz",
			"help": "",
			"hidden": false,
			"id": "relation70193212",
			"maxSelect": 0,
			"minSelect": 0,
			"name": "next",
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

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_zuKuzAF` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `starttime` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `title` + "`" + `,\n  ` + "`" + `bsm_id` + "`" + `,\n  ` + "`" + `meetingtime` + "`" + `,\n  ` + "`" + `team` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation3168962645")

		// remove field
		collection.Fields.RemoveById("relation70193212")

		return app.Save(collection)
	})
}
