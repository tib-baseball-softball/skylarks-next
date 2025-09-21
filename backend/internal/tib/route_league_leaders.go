package tib

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

// GetLeagueLeaders returns a nicely formatted response for top 10 stats in the given league to avoid making multiple
// direct BSM requests on the frontend side.
func GetLeagueLeaders() func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		leagueID := event.Request.PathValue("league")
		if leagueID == "" {
			return event.JSON(http.StatusBadRequest, "league ID missing")
		}
		result, err := strconv.Atoi(leagueID)
		if err != nil || result == 0 {
			return event.JSON(http.StatusBadRequest, "league ID must be an integer")
		}

		statsType := event.Request.URL.Query().Get("statsType")
		if statsType == "" {
			statsType = "batting"
		}

		leagueLeaderboard, err := GetLeagueTop10Data(event.App, leagueID, statsType)
		if err != nil {
			return errors.New("could not get league leaderboard")
		}

		return event.JSON(http.StatusOK, leagueLeaderboard)
	}
}
