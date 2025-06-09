package routes

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"net/http"
)

// GetFavoriteTeamData returns a handler function for the favorite team data endpoint
func GetFavoriteTeamData() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		teamID := e.Request.URL.Query().Get("team")
		if teamID == "" {
			return e.JSON(http.StatusBadRequest, "Team ID not found in request")
		}

		leagueGroupID := e.Request.URL.Query().Get("leagueGroup")
		if leagueGroupID == "" {
			return e.JSON(http.StatusBadRequest, "League Group ID not found in request")
		}

		homeDatasets, err := bsm.LoadHomeData(e.App, teamID, leagueGroupID)
		if err != nil {
			e.App.Logger().Error("Failed to load home data", "error", err)
			return e.JSON(http.StatusInternalServerError, "Internal error occurred")
		}

		return e.JSON(http.StatusOK, homeDatasets)
	}
}
