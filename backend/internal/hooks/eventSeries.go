package hooks

import (
	"fmt"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/pb"
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
		pb.EventsCollection,
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
	eventSeries := &pb.EventSeries{}
	eventSeries.SetProxyRecord(e.Record)

	startDate := eventSeries.SeriesStart()
	endDate := eventSeries.SeriesEnd()
	interval := eventSeries.Interval()
	duration := eventSeries.Duration()

	eventCollection, err := app.FindCollectionByNameOrId(pb.EventsCollection)
	if err != nil {
		return nil, err
	}

	existingEvents, err := app.FindRecordsByFilter(
		pb.EventsCollection,
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
	existingEventsMap := make(map[string]*pb.Event)

	// Create a map of existing events for easy lookup
	for _, event := range existingEvents {
		eventProxy := &pb.Event{}
		eventProxy.SetProxyRecord(event)

		key := fmt.Sprintf("%s-%s", eventProxy.StartTime().Time().Format(time.RFC3339), eventProxy.EndTime().Time().Format(time.RFC3339))
		existingEventsMap[key] = eventProxy
	}

	currentDate := startDate

	for currentDate.Before(endDate) {
		// Create start and end times for this specific event
		eventStart := currentDate
		eventEnd := currentDate.Add(time.Duration(duration) * time.Minute)

		// Check if an event already exists for this time slot
		key := fmt.Sprintf("%s-%s", eventStart.Time().Format(time.RFC3339), eventEnd.Time().Format(time.RFC3339))
		event, exists := existingEventsMap[key]

		if !exists {
			// Not found - create a new event
			event = &pb.Event{}
			event.SetProxyRecord(core.NewRecord(eventCollection))
		}

		event.SetStartTime(eventStart)
		event.SetEndTime(eventEnd)
		event.SetTitle(eventSeries.Title())
		event.SetTeam(eventSeries.Team())
		event.SetDesc(eventSeries.Desc())
		event.SetLocation(eventSeries.Location())
		event.SetSeries(eventSeries.Id)
		event.SetType(stats.Practice.String())

		events = append(events, event.Record)

		// Mark this event as processed
		delete(existingEventsMap, key)

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
