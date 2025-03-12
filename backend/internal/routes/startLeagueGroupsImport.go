package routes

import (
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"net/http"
	"slices"
)

// StartLeagueGroupsImport runs import cron for league groups exactly once for one specific club.
func StartLeagueGroupsImport(app core.App) func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		requireAuth := apis.RequireAuth()
		if err := requireAuth.Func(event); err != nil {
			return event.UnauthorizedError("no access", err)
		}

		clubID := event.Request.PathValue("club")
		club, err := app.FindFirstRecordByData("clubs", "id", clubID)
		if err != nil {
			return event.NotFoundError("Club with ID "+clubID+" not found", err)
		}

		clubAdmin := club.GetStringSlice("admins")
		if !slices.Contains(clubAdmin, event.Auth.Id) {
			return event.ForbiddenError("only club admins can start league imports", nil)
		}

		err = bsm.ImportLeagueGroups(app, &club.Id, nil)
		if err != nil {
			return event.InternalServerError("error importing league groups", err)
		}

		response := struct{ message string }{
			"import cron for club ID " + club.Id + " has been run successfully",
		}
		return event.JSON(http.StatusOK, response)
	}
}
