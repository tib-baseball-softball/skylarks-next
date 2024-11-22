package stats

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"strconv"
)

// LoadUserStats gets participation stats by event type and participation state
func LoadUserStats(user *core.Record, requestEvent *core.RequestEvent) (PersonalAttendanceStatsItem, error) {
	statsItem := PersonalAttendanceStatsItem{}
	seasonParam := requestEvent.Request.URL.Query().Get("season")
	eventTypeParam := requestEvent.Request.URL.Query().Get("eventType")

	if seasonParam != "" {
		season, err := strconv.Atoi(seasonParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse season: %v", err)
			return statsItem, err
		}
		statsItem.Season = season
	}

	if eventTypeParam != "" {
		eventType, err := enum.ParseEventType(eventTypeParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse eventType: %v", err)
			return statsItem, err
		}
		statsItem.Type = eventType
	}

	eventCounts, err := GetEventCounts(requestEvent.App, user, seasonParam)
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

	participations, err := GetParticipationStats(user, requestEvent.App, seasonParam, eventTypeParam, err)
	if err != nil {
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
