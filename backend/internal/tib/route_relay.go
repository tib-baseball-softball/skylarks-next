package tib

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

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

		stripped, err := stripResponseKeys(cachedResponseBody)
		if err != nil {
			e.App.Logger().Error("Failed to unmarshal cached BSM response", "error", err, "url", targetURL)
			return e.JSON(http.StatusInternalServerError, "Internal error occurred")
		}

		return e.Blob(http.StatusOK, "application/json", stripped)
	}
}

var keysToRemoveFromBSMResponse = []string{
	"current_player_list",
	"current_roster",
	"gamechanger",
}

func stripResponseKeys(cachedResponseBody string) ([]byte, error) {
	var arrayShapeData []map[string]any
	var objectShapeData map[string]any
	var unmarshalTypeError *json.UnmarshalTypeError
	var bytes []byte

	err := json.Unmarshal([]byte(cachedResponseBody), &arrayShapeData)
	if err != nil {
		if errors.As(err, &unmarshalTypeError) {
			// response is not JSON array, try to unmarshal as an object
			err = json.Unmarshal([]byte(cachedResponseBody), &objectShapeData)
			if err != nil {
				return nil, err
			}

			for _, key := range keysToRemoveFromBSMResponse {
				delete(objectShapeData, key)
			}
			bytes, err = json.Marshal(objectShapeData)
			if err != nil {
				return nil, err
			}
		} else {
			// all other errors can bubble up
			return nil, err
		}
	} else {
		// response is JSON array
		for _, item := range arrayShapeData {
			for _, key := range keysToRemoveFromBSMResponse {
				delete(item, key)
			}
		}
		bytes, err = json.Marshal(arrayShapeData)
		if err != nil {
			return nil, err
		}
	}

	return bytes, nil
}

func rewriteURLForProxying(url url.URL) url.URL {
	bsmHost := os.Getenv("BSM_API_HOST")
	targetURL := url
	targetURL.Scheme = "https"
	targetURL.Host = bsmHost
	targetURL.Path = strings.TrimPrefix(url.Path, "/api/bsm/relay")

	newQuery := url.Query()
	newQuery.Set("api_key", os.Getenv("BSM_API_KEY"))
	targetURL.RawQuery = newQuery.Encode()

	return targetURL
}
