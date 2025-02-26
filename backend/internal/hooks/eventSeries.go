package hooks

import (
	"fmt"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/internal/stats"
	"time"
)

func CreateOrUpdateEventsForSeries(e *core.RecordEvent) error {
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

func DeleteEventsForSeries(e *core.RecordEvent) error {
	eventSeries := e.Record

	eventsToBeDeleted, err := e.App.FindRecordsByFilter(
		"events",
		"series = {:seriesID}",
		"",
		0,
		0,
		dbx.Params{"seriesID": eventSeries.Id},
	)
	if err != nil {
		return err
	}

	err = e.App.RunInTransaction(func(txApp core.App) error {
		for _, record := range eventsToBeDeleted {
			if err := txApp.Delete(record); err != nil {
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

	location, err := time.LoadLocation("UTC")
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

	existingEvents, err := app.FindRecordsByFilter(
		"events",
		"series = {:seriesID}",
		"",
		0,
		0,
		dbx.Params{"seriesID": eventSeries.Id},
	)
	if err != nil {
		return nil, err
	}

	var events []*core.Record
	existingEventsMap := make(map[string]*core.Record)

	// Create a map of existing events for easy lookup
	for _, event := range existingEvents {
		key := fmt.Sprintf("%s-%s", event.GetDateTime("starttime").Time().Format(time.RFC3339), event.GetDateTime("endtime").Time().Format(time.RFC3339))
		existingEventsMap[key] = event
	}

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

		if eventStart.After(eventSeries.GetDateTime("series_start").Time()) && eventStart.Before(endDate) {
			// Check if an event already exists for this time slot
			key := fmt.Sprintf("%s-%s", eventStart.Format(time.RFC3339), eventEnd.Format(time.RFC3339))
			event, exists := existingEventsMap[key]

			if !exists {
				// Not found - create a new event
				event = core.NewRecord(eventCollection)
			}

			event.Set("starttime", eventStartDateTime)
			event.Set("meetingtime", eventStartDateTime)
			event.Set("endtime", eventEndDateTime)
			event.Set("title", eventSeries.GetString("title"))
			event.Set("team", eventSeries.GetString("team"))
			event.Set("desc", eventSeries.GetString("desc"))
			event.Set("location", eventSeries.GetString("location"))
			event.Set("series", eventSeries.Id)
			event.Set("type", stats.Practice)

			events = append(events, event)

			// Mark this event as processed
			delete(existingEventsMap, key)
		}
		currentDate = currentDate.AddDate(0, 0, interval)
	}

	// Delete any leftover events that are no longer part of the updated series
	for _, staleEvent := range existingEventsMap {
		err := app.Delete(staleEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to delete stale event: %w", err)
		}
	}

	return events, nil
}
