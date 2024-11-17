package hooks

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/cronjobs"
	"strconv"
)

// TriggerLeagueImport called on every league group listing, if the result is empty and special query parameters `season`
// and `club` are sent, an import is automatically triggered.
func TriggerLeagueImport(app *pocketbase.PocketBase, event *core.RecordsListRequestEvent) error {
	s := event.Request.URL.Query().Get("season")
	club := event.Request.URL.Query().Get("club")

	if len(event.Records) != 0 || s == "" || club == "" {
		return event.Next()
	}

	season, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	err = cronjobs.ImportLeagueGroups(app, &club, &season)
	if err != nil {
		return apis.NewInternalServerError("league import failed", err)
	}

	return event.Next()
}