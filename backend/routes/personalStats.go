package routes

import (
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/middlewares"
	"github.com/tib-baseball-softball/skylarks-next/stats"
)

func GetUserStats() func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		requireAuth := apis.RequireAuth()
		if err := requireAuth.Func(event); err != nil {
			return err
		}

		userID := event.Request.PathValue("user")
		user, err := event.App.FindRecordById("users", userID)
		if err != nil {
			return err
		}

		requireUserAccess := middlewares.RequireUserAccess(user)
		if err := requireUserAccess.Func(event); err != nil {
			return err
		}

		statsItem, err := stats.LoadUserStats(user, event)
		if err != nil {
			return err
		}

		return event.JSON(200, statsItem)
	}
}
