package stats

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type SimpleCount struct {
	Count int `json:"count" db:"count"`
}

type EventCounts struct {
	Games     int `db:"games"`
	Practices int `db:"practices"`
	Misc      int `db:"misc"`
	AllEvents int `db:"all_events"`
}

func GetEventCounts(app core.App, user *core.Record) (EventCounts, error) {
	result := EventCounts{}
	query := app.DB().
		Select(
			"COUNT(CASE WHEN events.type = 'game' THEN events.id END) AS games",
			"COUNT(CASE WHEN events.type = 'practice' THEN events.id END) AS practices",
			"COUNT(CASE WHEN events.type = 'misc' THEN events.id END) AS misc",
			"COUNT(events.id) AS all_events",
		).
		From("events")

	if user != nil {
		teams := user.GetStringSlice("teams")
		var expressions []dbx.Expression
		for _, team := range teams {
			expressions = append(expressions, dbx.HashExp{"team": team})
		}
		query.Where(dbx.Or(expressions...))
	}
	err := query.One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
