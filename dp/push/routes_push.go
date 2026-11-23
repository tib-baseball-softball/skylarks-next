package push

import (
	"net/http"

	"ganigeorgiev.com/pocketbase/dbx"
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func NotifyWithTestPushMessage(ps PushService, app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		userID := e.Request.PathValue("user")

		if userID != e.Auth.Id {
			return e.ForbiddenError("", nil)
		}

		subRecords, err := app.FindRecordsByFilter(
			base.PushSubscriptionsCollection,
			"user.id = {:userID}",
			"",
			0,
			0,
			dbx.Params{"userID": userID},
		)
		if err != nil {
			return err
		}

		for _, subRecord := range subRecords {
			sub := &base.PushSubscription{}
			sub.SetProxyRecord(subRecord)
			err := ps.handleTestPush(new(sub.ToWebPushSubscription()))
			if err != nil {
				app.Logger().Warn("failed to send push message", "error", err, "sub", sub)
				continue
			}
		}

		return e.JSON(http.StatusOK, "Test push sent")
	}
}

func DeletePushSubscription(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return e.InternalServerError("", err)
		}
		endpoint := info.Body["endpoint"].(string)

		record, err := app.FindFirstRecordByFilter(
			base.PushSubscriptionsCollection,
			"endpoint = {:endpoint}",
			dbx.Params{"endpoint": endpoint},
		)
		if err != nil {
			app.Logger().Error("failed to find subscription to delete", "error", err, "endpoint", endpoint)
			return e.BadRequestError("", err)
		}
		var sub base.PushSubscription
		sub.SetProxyRecord(record)

		if sub.User() != e.Auth.Id {
			return e.ForbiddenError("", nil)
		}

		err = app.Delete(sub)
		if err != nil {
			app.Logger().Error("failed to delete subscription", "error", err, "sub", sub)
			return e.InternalServerError("", err)
		}

		return e.NoContent(http.StatusNoContent)
	}
}
