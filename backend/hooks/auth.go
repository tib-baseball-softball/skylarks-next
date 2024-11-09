package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"time"
)

func SetLastLogin(app *pocketbase.PocketBase, record *core.Record) (err error) {
	currentTime := time.Now()
	record.Set("last_login", currentTime)

	if err := app.Save(record); err != nil {
		app.Logger().Error("Persisting user last login time failed", "error", err)

		return err
	}

	return nil
}
