package dp

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// GetGamesCount returns the number of official league games in the database for a given team and season.
// Requires a valid auth record that belongs to the team queried.
func GetGamesCount(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		teamID := e.Request.PathValue("team")
		team, err := app.FindRecordById("teams", teamID)
		if err != nil {
			return e.BadRequestError("Failed to load teams record", err)
		}

		info, err := e.RequestInfo()
		if err != nil {
			return e.BadRequestError("Failed to retrieve request info", err)
		}

		canAccess, err := e.App.CanAccessRecord(team, info, team.Collection().ViewRule)
		if !canAccess {
			return e.ForbiddenError("Only members and admins can query team game count", err)
		}

		// do actual work
		season := e.Request.URL.Query().Get("season")

		if season == "" {
			return apis.NewBadRequestError("Query param for season is required", nil)
		}

		result, err := LoadCount(app, teamID, season)
		if err != nil {
			return apis.NewInternalServerError("Error querying event", err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

type SimpleCount struct {
	Count int `json:"count" db:"count"`
}

func LoadCount(app core.App, teamID string, season string) (SimpleCount, error) {
	result := SimpleCount{}
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

type JoinTeamPayload struct {
	UserID    string `json:"userID"`
	SignupKey string `json:"signupKey"`
}

func JoinTeam(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		teamID := e.Request.PathValue("team")
		team := &Team{}
		err := FindRecordProxyByID(app, team, teamID)
		if err != nil {
			return e.BadRequestError("Failed to load teams record", err)
		}

		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(e.Request.Body)

		body, err := io.ReadAll(e.Request.Body)
		if err != nil {
			app.Logger().Error("Failed to read request body", "file", "routes_teams.go", "err", err, "team", team)
			return e.InternalServerError("Failed to read request body", err)
		}

		payload := JoinTeamPayload{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			app.Logger().Error("Failed to parse request body", "file", "routes_teams.go", "err", err)
			return e.BadRequestError("Supplied data is malformed", err)
		}

		if payload.UserID != e.Auth.Id {
			return e.ForbiddenError("User ID does not match authenticated user", nil)
		}

		user := &User{}
		err = FindRecordProxyByID(app, user, payload.UserID)
		if err != nil {
			return e.BadRequestError("Failed to load user record", err)
		}

		if slices.Contains(user.Teams(), teamID) {
			return e.BadRequestError("User is already a member of this team", nil)
		}

		if payload.SignupKey != team.SignupKey() {
			return e.ForbiddenError("Provided signup key does not match team signup key", nil)
		}

		user.AddTeam(teamID)
		err = app.Save(user)
		if err != nil {
			return e.InternalServerError("Failed to save user record", err)
		}

		return e.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "User joined team successfully",
		})
	}
}
