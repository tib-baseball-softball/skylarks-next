package tib

import (
	"log/slog"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type CacheCleanupApp interface {
	FindRecordsByFilter(collectionModelOrIdentifier any, filter string, sort string, limit int, offset int, params ...dbx.Params) ([]*core.Record, error)
	Logger() *slog.Logger
	Delete(model core.Model) error
}

func CleanupOutdatedCaches(app CacheCleanupApp) error {
	outdatedCaches, err := app.FindRecordsByFilter(dp.RequestCacheCollection, "updated <= @yesterday", "", 0, 0)
	if err != nil {
		app.Logger().Error("Error fetching outdated caches", "error", err)
		return err
	}
	for _, cache := range outdatedCaches {
		if err := app.Delete(cache); err != nil {
			app.Logger().Error("Error deleting outdated cache", "error", err, "cache", cache)
			return err
		}
	}
	app.Logger().Info("Deleted outdated caches", "number of caches", len(outdatedCaches))
	return nil
}
