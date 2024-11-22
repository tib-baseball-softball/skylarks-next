package stats

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"strconv"
)

// LoadUserStats gets participation stats by event type and participation state
func LoadUserStats(user *core.Record, requestEvent *core.RequestEvent) (PersonalAttendanceStatsItem, error) {
	statsItem := PersonalAttendanceStatsItem{}
	seasonParam := requestEvent.Request.URL.Query().Get("season")
	eventTypeParam := requestEvent.Request.URL.Query().Get("eventType")

	eventCounts, err := GetEventCounts(requestEvent.App, user)
	if err != nil {
		requestEvent.App.Logger().Error("failed to get eventCounts: %v", err)
		return statsItem, err
	}
	statsItem.TotalPossibleEvents = eventCounts.AllEvents

	// data structure for frontend
	totalsIn := ParticipationTotal{
		Type:   enum.In,
		Amount: 0,
	}
	totalsMaybe := ParticipationTotal{
		Type:   enum.Maybe,
		Amount: 0,
	}
	totalsOut := ParticipationTotal{
		Type:   enum.Out,
		Amount: 0,
	}

	totalsGames := AttendanceTotal{
		Type:     enum.Game,
		Attended: 0,
		Total:    eventCounts.Games,
	}
	totalsPractice := AttendanceTotal{
		Type:     enum.Practice,
		Attended: 0,
		Total:    eventCounts.Practices,
	}
	totalsMisc := AttendanceTotal{
		Type:     enum.Misc,
		Attended: 0,
		Total:    eventCounts.Misc,
	}

	var participations []ParticipationByPerson

	query := requestEvent.App.DB().
		Select(
			"users.id",
			"users.last_name",
			"users.first_name",
			"events.type",
			"COUNT(CASE WHEN participations.state = 'in' THEN participations.id END)    AS inCount",
			"COUNT(CASE WHEN participations.state = 'out' THEN participations.id END)   AS outCount",
			"COUNT(CASE WHEN participations.state = 'maybe' THEN participations.id END) AS maybeCount",
			"COUNT(participations.id)",
		).
		From("users").
		LeftJoin("participations", dbx.NewExp("participations.user = users.id")).
		InnerJoin("events", dbx.NewExp("participations.event = events.id")).
		GroupBy("users.id", "users.first_name", "users.last_name", "events.type")

	if user.Id != "" {
		query.AndWhere(dbx.NewExp("users.id = {:id}", dbx.Params{"id": user.Id}))
	}

	if seasonParam != "" {
		season, err := strconv.Atoi(seasonParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse season: %v", err)
			return statsItem, err
		}
		statsItem.Season = season

		query.AndWhere(dbx.NewExp("strftime('%Y', events.starttime) = {:season}", dbx.Params{"season": seasonParam}))
	}

	if eventTypeParam != "" {
		eventType, err := enum.ParseEventType(eventTypeParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse eventType: %v", err)
			return statsItem, err
		}
		statsItem.Type = eventType

		query.AndWhere(dbx.NewExp("events.type = {:eventType}", dbx.Params{"eventType": eventTypeParam}))
	}

	err = query.All(&participations)
	if err != nil {
		requestEvent.App.Logger().Error("failed to load participations: %v", err)
		return statsItem, err
	}
	// query can only get the totals for which participation data exists,
	// add raw possible totals from previous query
	for i, participation := range participations {
		eventType, err := enum.ParseEventType(participation.Type)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse eventType: %v", err)
			return statsItem, err
		}
		switch eventType {
		case enum.Game:
			participations[i].TotalCount = eventCounts.Games
			totalsGames.Attended = participation.InCount
		case enum.Practice:
			participations[i].TotalCount = eventCounts.Practices
			totalsPractice.Attended = participation.InCount
		case enum.Misc:
			participations[i].TotalCount = eventCounts.Misc
			totalsMisc.Attended = participation.InCount
		}

		totalsIn.Amount = totalsIn.Amount + participation.InCount
		totalsMaybe.Amount = totalsMaybe.Amount + participation.MaybeCount
		totalsOut.Amount = totalsOut.Amount + participation.OutCount
	}

	statsItem.ParticipationTotals = append(statsItem.ParticipationTotals, totalsIn, totalsMaybe, totalsOut)
	statsItem.AttendanceTotals = append(statsItem.AttendanceTotals, totalsPractice, totalsMisc, totalsGames)
	statsItem.Values = append(statsItem.Values, participations...)

	return statsItem, nil
}
