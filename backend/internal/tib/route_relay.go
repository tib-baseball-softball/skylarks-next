package tib

import (
	"errors"
	"net/http"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/pocketbase/pocketbase/core"
)

// GetRelayedBSMData relays BSM request so the client-side does not need to know API keys.
// Uses allowlist for URLs to make sure only public information is disclosed.
// Makes Pocketbase take the role of the SvelteKit Node.js server for server load functions.
func GetRelayedBSMData() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		targetURL := rewriteURLForProxying(*e.Request.URL)

		cachedResponseBody, err := GetCachedBSMResponse(e.App, &targetURL)
		if err != nil {
			var bsmErr *bsm.URLAllowlistError
			if errors.As(err, &bsmErr) {
				return e.JSON(http.StatusForbidden, err.Error())
			} else {
				e.App.Logger().Error("Failed to get cached BSM response", "error", err, "url", targetURL)
				return e.JSON(http.StatusInternalServerError, "Internal error occurred")
			}
		}

		// TODO: strip response

		return e.Blob(http.StatusOK, "application/json", []byte(cachedResponseBody))
	}
}
