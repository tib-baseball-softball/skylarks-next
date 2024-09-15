package hooks

import (
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func SetLastLogin(app *pocketbase.PocketBase, record *models.Record) (err error) {
	currentTime := time.Now()
	record.Set("last_login", currentTime)

	if err := app.Dao().SaveRecord(record); err != nil {
		app.Logger().Error("Persisting user last login time failed", "error", err)

		return err
	}

	return nil
}
