package dp

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// AppRecordProxy defines the interface for all Diamond Planner Record Proxies
// used for type-safe access to record model fields.
type AppRecordProxy interface {
	core.RecordProxy
	ID() string
	CollectionName() string
}

// FindRecordProxyByID retrieves a record proxy by its ID from the database. Simple Wrapper for `app.FindRecordById()` to
// prevent having to load the result manually into a proxy object.
// Caution: `proxy` needs to be a pointer to a struct that implements AppRecordProxy, not a value.
func FindRecordProxyByID(app core.App, proxy AppRecordProxy, id string) error {
	err := app.RecordQuery(proxy.CollectionName()).
		AndWhere(dbx.HashExp{"id": id}).
		Limit(1).
		One(proxy)
	if err != nil {
		return err
	}
	return nil
}
