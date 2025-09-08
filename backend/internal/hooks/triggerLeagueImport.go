package hooks

import (
	"strconv"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
)

// TriggerLeagueImport called on every league group listing, if the result is empty and special query parameters `season`
// and `club` are sent, an import is automatically triggered.
func TriggerLeagueImport(app core.App, event *core.RecordsListRequestEvent) error {
	s := event.Request.URL.Query().Get("season")
	club := event.Request.URL.Query().Get("club")

	if len(event.Records) != 0 || s == "" || club == "" {
		return event.Next()
	}

	season, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	err = bsm.ImportLeagueGroups(app, &club, &season)
	if err != nil {
		return apis.NewInternalServerError("league import failed", err)
	}

	return event.Next()
}
