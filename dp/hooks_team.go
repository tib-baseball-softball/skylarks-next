package dp

import (
	"slices"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

// HideTeamSignupKey hides the signup key for a team record if the current user is not an admin.
func HideTeamSignupKey(e *core.RecordEnrichEvent) error {
	team := &base.Team{}
	team.SetProxyRecord(e.Record)

	team.HideSignupKey()

	auth := e.RequestInfo.Auth

	if auth != nil && slices.Contains(team.Admins(), e.RequestInfo.Auth.Id) {
		team.UnhideSignupKey()
	}

	return e.Next()
}
