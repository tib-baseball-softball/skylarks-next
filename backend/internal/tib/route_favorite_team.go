package tib

import (
	"net/http"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

// GetFavoriteTeamData returns a handler function for the favorite team data endpoint
func GetFavoriteTeamData() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		teamID := e.Request.URL.Query().Get("team")
		if teamID == "" {
			return e.JSON(http.StatusBadRequest, "Team ID not found in request")
		}

		convertedTeamID, err := strconv.Atoi(teamID)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "Team ID must be an integer")
		}

		season := e.Request.URL.Query().Get("season")
		if season == "" {
			return e.JSON(http.StatusBadRequest, "Season not found in request")
		}

		convertedSeason, err := strconv.Atoi(season)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "Season must be an integer")
		}

		response, err := GetCachedDatasetResponse(e.App, convertedTeamID, convertedSeason)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, "Internal error occurred")
		}
		return e.Blob(http.StatusOK, "application/json", response)
	}
}
