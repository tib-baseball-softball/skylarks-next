package routes

import (
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"net/http"
	"net/url"
)

// GetRelayedBSMData relays BSM request so the client-side does not need to know API keys.
// Uses allowlist for URLs to make sure only public information is disclosed.
func GetRelayedBSMData() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		// request checks
		bsmURL := e.Request.URL.Query().Get("url")
		if bsmURL == "" {
			return e.JSON(http.StatusBadRequest, "BSM URL not found in request")
		}
		clubID := e.Request.URL.Query().Get("club")
		if clubID == "" {
			return e.JSON(http.StatusBadRequest, "Club ID not found in request")
		}
		// actual operations
		parsedURL, err := url.Parse(bsmURL)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "Failed to parse given URL string into a valid URL")
		}

		urlWithKey, err := bsm.AppendAPIKey(e.App, *parsedURL, clubID)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, "Internal error occurred")
		}

		cachedResponseBody, err := bsm.GetCachedBSMResponse(e.App, &urlWithKey)
		if err != nil {
			var bsmErr *bsm.URLWhitelistError
			if errors.As(err, &bsmErr) {
				return e.JSON(http.StatusForbidden, err.Error())
			} else {
				e.App.Logger().Error("Failed to get cached BSM response", "error", err)
				return e.JSON(http.StatusInternalServerError, "Internal error occurred")
			}
		}

		return e.Blob(http.StatusOK, "application/json", []byte(cachedResponseBody))
	}
}
