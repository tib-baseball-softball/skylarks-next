package routes

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"github.com/tib-baseball-softball/skylarks-next/middlewares"
	"github.com/tib-baseball-softball/skylarks-next/model"
	"github.com/tib-baseball-softball/skylarks-next/utility"
	"strconv"
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

		statsItem, err := loadUserStats(user, event)
		if err != nil {
			return err
		}

		return event.JSON(200, statsItem)
	}
}

func loadUserStats(user *core.Record, event *core.RequestEvent) (model.PersonalAttendanceStatsItem, error) {
	statsItem := model.PersonalAttendanceStatsItem{}
	seasonParam := event.Request.URL.Query().Get("seasonParam")

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

	participations, err := event.App.FindRecordsByFilter(
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

	for _, participation := range participations {
		// raw participation states
		switch participation.GetString("state") {
		case string(enum.In):
			totalsIn.Amount++
		case string(enum.Maybe):
			totalsMaybe.Amount++
		case string(enum.Out):
			totalsOut.Amount++
		}
		// attendance by type
		// TODO: implement
	}

	statsItem.ParticipationTotals = append(statsItem.ParticipationTotals, totalsIn, totalsMaybe, totalsOut)

	return statsItem, nil
}
