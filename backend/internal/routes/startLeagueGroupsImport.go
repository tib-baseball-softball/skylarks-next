package routes

import (
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/cronjobs"
	"net/http"
	"slices"
)

// StartLeagueGroupsImport runs import cron for league groups exactly once for one specific club.
func StartLeagueGroupsImport(app core.App) func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		requireAuth := apis.RequireAuth()
		if err := requireAuth.Func(event); err != nil {
			return err
		}
		clubID := event.Request.PathValue("club")
		club, err := app.FindFirstRecordByData("clubs", "id", clubID)
		if err != nil {
			return err
		}

		clubAdmin := club.GetStringSlice("admins")
		if !slices.Contains(clubAdmin, event.Auth.Id) {
			return apis.NewUnauthorizedError("only club admins can start league imports", nil)
		}

		err = cronjobs.ImportLeagueGroups(app, &club.Id, nil)
		if err != nil {
			return err
		}

		response := struct{ message string }{
			"import cron for club ID " + club.Id + " has been run successfully",
		}
		return event.JSON(http.StatusOK, response)
	}
}
