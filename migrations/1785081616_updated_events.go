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
				"CREATE INDEX ` + "`" + `idx_38rk28bmkd` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `team` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_vizqfueccq` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `club` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_nfoliwl5a7` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `additional_teams` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_nrf9dfdobe` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `club` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_sffndaod6n` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `additional_teams` + "`" + `\n)"
			]
		}`), &collection); err != nil {
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
				"CREATE INDEX ` + "`" + `idx_zuKuzAF` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `bsm_id` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_8ixkwcxr4q` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `type` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_miocwua0bp` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `team` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_snxc3kufmb` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `series` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_0c501exuc1` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `location` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_38rk28bmkd` + "`" + ` ON ` + "`" + `events` + "`" + ` (\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `team` + "`" + `\n)"
			]
		}`), &collection); err != nil {
			return err
		}

		return app.Save(collection)
	})
}
