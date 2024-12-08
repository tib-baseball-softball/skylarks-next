package hooks

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/enum"
	"time"
)

func CreateEventsForSeries(e *core.RecordEvent) error {
	events, err := generateSeriesEvents(e.App, e)
	if err != nil {
		return err
	}

	err = e.App.RunInTransaction(func(txApp core.App) error {
		for _, record := range events {
			if err := txApp.Save(record); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return e.Next()
}

func generateSeriesEvents(app core.App, e *core.RecordEvent) ([]*core.Record, error) {
	eventSeries := e.Record

	startDate := eventSeries.GetDateTime("series_start").Time()
	endDate := eventSeries.GetDateTime("series_end").Time()
	interval := eventSeries.GetInt("interval")
	weekday := eventSeries.GetInt("weekday")

	startTimeString := eventSeries.GetString("starttime")
	endTimeString := eventSeries.GetString("endtime")

	// MARK: using hardcoded time zone for Germany here
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	startTimeOfDay, err := time.ParseInLocation("15:04", startTimeString, location)
	endTimeOfDay, err := time.ParseInLocation("15:04", endTimeString, location)
	if err != nil {
		return nil, err
	}

	// Adjust startDate to the next occurrence of the specified weekday
	for int(startDate.Weekday()) != weekday {
		startDate = startDate.AddDate(0, 0, 1)
	}

	eventCollection, err := app.FindCollectionByNameOrId("events")
	if err != nil {
		return nil, err
	}

	var events []*core.Record
	currentDate := startDate

	for !currentDate.After(endDate) {
		// Create start and end times for this specific event
		eventStart := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(),
			startTimeOfDay.Hour(), startTimeOfDay.Minute(), startTimeOfDay.Second(), 0, time.UTC)

		eventEnd := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(),
			endTimeOfDay.Hour(), endTimeOfDay.Minute(), endTimeOfDay.Second(), 0, time.UTC)

		eventStartDateTime, err := types.ParseDateTime(eventStart)
		if err != nil {
			return nil, err
		}
		eventEndDateTime, err := types.ParseDateTime(eventEnd)
		if err != nil {
			return nil, err
		}

		// TODO: fix timezone issues

		if eventStart.After(eventSeries.GetDateTime("series_start").Time()) && eventStart.Before(endDate) {
			event := core.NewRecord(eventCollection)
			event.Set("starttime", eventStartDateTime)
			event.Set("meetingtime", eventStartDateTime)
			event.Set("endtime", eventEndDateTime)
			event.Set("title", eventSeries.GetString("title"))
			event.Set("team", eventSeries.GetString("team"))
			event.Set("desc", eventSeries.GetString("desc"))
			event.Set("location", eventSeries.GetString("location"))
			event.Set("series", eventSeries.Id)
			event.Set("type", enum.Practice)

			events = append(events, event)
		}
		currentDate = currentDate.AddDate(0, 0, interval)
	}

	return events, nil
}
