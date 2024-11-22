package routes

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/stats"
	"net/http"
	"slices"
)

// GetGamesCount returns the amount of official league games in the database for a given team and season.
// Requires valid auth record that belongs to the team queried.
func GetGamesCount(app *pocketbase.PocketBase) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		requireAuth := apis.RequireAuth()
		if err := requireAuth.Func(e); err != nil {
			return err
		}

		authTeams := e.Auth.GetStringSlice("teams")
		teamID := e.Request.PathValue("team")

		if !slices.Contains(authTeams, teamID) {
			return apis.NewUnauthorizedError("Only members can query team game count", nil)
		}

		season := e.Request.URL.Query().Get("season")

		if season == "" {
			return apis.NewBadRequestError("query param for season is required", nil)
		}

		result, err := LoadCount(app, teamID, season)
		if err != nil {
			return apis.NewInternalServerError("Error querying event", err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func LoadCount(app *pocketbase.PocketBase, teamID string, season string) (stats.SimpleCount, error) {
	result := stats.SimpleCount{}
	err := app.DB().
		Select("COUNT(events.id) AS count").
		From("events").
		Where(dbx.NewExp("bsm_id != 0")).
		AndWhere(dbx.NewExp("team = {:team}", dbx.Params{"team": teamID})).
		AndWhere(dbx.NewExp("strftime('%Y', events.starttime) = {:season}", dbx.Params{"season": season})).
		One(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}
