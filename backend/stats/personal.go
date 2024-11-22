package stats

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"github.com/tib-baseball-softball/skylarks-next/utility"
	"strconv"
)

// LoadUserStats gets participation stats by event type and participation state
func LoadUserStats(user *core.Record, requestEvent *core.RequestEvent) (model.PersonalAttendanceStatsItem, error) {
	statsItem := model.PersonalAttendanceStatsItem{}
	seasonParam := requestEvent.Request.URL.Query().Get("seasonParam")

	filter := "user = {:user}"
	var start string
	var end string

	if seasonParam != "" {
		season, err := strconv.Atoi(seasonParam)
		if err != nil {
			return statsItem, err
		}
		statsItem.Season = season

		start, end = utility.YearStartAndEnd(season)
		filter = filter + " && created >= {:start} && created <= {:end}"
	}

	participations, err := requestEvent.App.FindRecordsByFilter(
		"participations",
		filter,
		"",
		0,
		0,
		dbx.Params{
			"user":  user.Id,
			"start": start,
			"end":   end,
		},
	)
	if err != nil {
		return statsItem, err
	}

	if errs := requestEvent.App.ExpandRecords(participations, []string{"event"}, nil); len(errs) > 0 {
		requestEvent.App.Logger().Error("failed to expand records: %v", errs)
	}
	statsItem.TotalPossibleEvents = len(participations)

	totalsIn := model.ParticipationTotal{
		Type:   enum.In,
		Amount: 0,
	}
	totalsMaybe := model.ParticipationTotal{
		Type:   enum.Maybe,
		Amount: 0,
	}
	totalsOut := model.ParticipationTotal{
		Type:   enum.Out,
		Amount: 0,
	}

	totalsGames := model.AttendanceTotal{
		Type:     enum.Game,
		Attended: 0,
		Total:    0,
	}
	totalsPractice := model.AttendanceTotal{
		Type:     enum.Practice,
		Attended: 0,
		Total:    0,
	}
	totalsMisc := model.AttendanceTotal{
		Type:     enum.Misc,
		Attended: 0,
		Total:    0,
	}

	for _, participation := range participations {
		// raw participation states
		state := participation.GetString("state")
		switch state {
		case string(enum.In):
			totalsIn.Amount++
		case string(enum.Maybe):
			totalsMaybe.Amount++
		case string(enum.Out):
			totalsOut.Amount++
		}
		// attendance by type
		event := participation.ExpandedOne("event")
		if event == nil {
			requestEvent.App.Logger().Error("failed to expand records: no event", "event", event)
			continue
		}
		eventType, err := enum.ParseEventType(event.GetString("type"))
		if err != nil {
			requestEvent.App.Logger().Error("failed to expand records: invalid event type", "eventType", eventType)
			continue
		}

		switch eventType {
		case enum.Game:
			totalsGames.Total++
			if state == string(enum.In) {
				totalsGames.Attended++
			}
		case enum.Practice:
			totalsPractice.Total++
			if state == string(enum.In) {
				totalsPractice.Attended++
			}
		case enum.Misc:
			totalsMisc.Total++
			if state == string(enum.In) {
				totalsMisc.Attended++
			}
		}
	}

	statsItem.ParticipationTotals = append(statsItem.ParticipationTotals, totalsIn, totalsMaybe, totalsOut)
	statsItem.AttendanceTotals = append(statsItem.AttendanceTotals, totalsPractice, totalsMisc, totalsGames)

	return statsItem, nil
}
