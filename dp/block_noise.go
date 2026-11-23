package dp

import (
	"net/http"
	"slices"
	"strings"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
)

// BlockKnownPaths acts as a simple blocker for some scriptkiddie "attacks" (every 404 response creates a log entry).
// Not intended to be an actual security measure.
func BlockKnownPaths(e *core.RequestEvent) error {
	if e.Request == nil || e.Request.URL == nil {
		return e.Next()
	}

	blockedPaths := []string{
		"/api/graphql",
		"/api/gql",
		"/api/gql",
		"/graphql/api",
		"/graphql",
	}

	if strings.Contains(e.Request.URL.Path, ".php") ||
		slices.Contains(blockedPaths, e.Request.URL.Path) ||
		(e.Request.URL.Path == "/" && e.Request.Method == http.MethodPost) {
		return e.String(http.StatusGone, "")
	}

	return e.Next()
}
