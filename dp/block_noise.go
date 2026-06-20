package dp

import (
	"net/http"
	"slices"
	"strings"

	"github.com/pocketbase/pocketbase/core"
)

// BlockKnownPaths acts as a simple blocker for some scriptkiddie "attacks" (every 404 response creates a log entry).
// Not intended to be an actual security measure.
func BlockKnownPaths(e *core.RequestEvent) error {
	blockedPaths := []string{
		"/api/graphql",
		"/api/gql",
		"/api/gql",
		"/graphql/api",
		"/graphql",
	}

	if strings.Contains(e.Request.URL.Path, ".php") ||
		slices.Contains(blockedPaths, e.Request.URL.Path) {
		return e.String(http.StatusGone, "")
	}

	return e.Next()
}
