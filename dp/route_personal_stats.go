package dp

import (
	"net/http"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func GetUserStats() func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		userID := event.Request.PathValue("user")
		user, err := event.App.FindRecordById(base.UserCollection, userID)
		if err != nil {
			return event.NotFoundError("invalid user provided", err)
		}

		requireUserAccess := RequireUserAccess(user)
		if err := requireUserAccess.Func(event); err != nil {
			return event.ForbiddenError("no access", err)
		}

		statsItem, err := LoadUserStats(user, event)
		if err != nil {
			return event.InternalServerError("", err)
		}

		return event.JSON(http.StatusOK, statsItem)
	}
}
