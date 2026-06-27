package dp

import (
	"container/list"
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/security"
	"github.com/pocketbase/pocketbase/tools/types"
)

type EventSeriesMode string

const (
	Create EventSeriesMode = "create"
	Update EventSeriesMode = "update"
)

// CreateOrUpdateEventsForSeries is the main entry point for event series logic.
// Triggered on all DB operations AFTER successful PocketBase validation of the series record, but before persistence.
func CreateOrUpdateEventsForSeries(e *core.RecordEvent, mode EventSeriesMode) error {
	events, err := generateSeriesEvents(e.App, e.Record, mode)
	if err != nil {
		return err
	}

	err = e.App.RunInTransaction(func(txApp core.App) error {
		for _, record := range events {
			// validations are skipped: other events in the same series might not be persisted yet,
			// so relation validation fails
			if err := txApp.SaveNoValidate(record); err != nil {
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

// DeleteEventsForSeries deletes all events associated with a given event series record.
func DeleteEventsForSeries(e *core.RecordEvent) error {
	eventSeries := &EventSeries{}
	eventSeries.SetProxyRecord(e.Record)

	eventsToBeDeleted, _, err := findEventRecordsForSeries(e.App, eventSeries)
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

// generateSeriesEvents contains the main logic to handle single events belonging to a series.
func generateSeriesEvents(app core.App, record *core.Record, mode EventSeriesMode) ([]*Event, error) {
	eventSeries := &EventSeries{}
	eventSeries.SetProxyRecord(record)

	startDateSeries := eventSeries.SeriesStart()
	endDateSeries := eventSeries.SeriesEnd()
	seriesInterval := eventSeries.Interval()
	seriesEventDuration := eventSeries.Duration()

	eventCollection, err := app.FindCollectionByNameOrId(EventsCollection)
	if err != nil {
		return nil, err
	}

	errorContext := &ErrorContext{
		Key: "local",
		Values: map[string]any{
			"seriesID":   eventSeries.Id,
			"seriesName": eventSeries.Title(),
		},
	}

	existingEvents, first, err := findEventRecordsForSeries(app, eventSeries)
	if err != nil {
		LogErrorInternalExternal(app, err, errorContext, nil)
		return nil, err
	}

	eventLinkedList, err := createEventSeriesLinkedList(existingEvents, first)
	if err != nil {
		LogErrorInternalExternal(app, err, errorContext, nil)
		return nil, err
	}

	location, err := LoadAppTimeZone()
	if err != nil {
		return nil, err
	}

	// reads timezone information to ensure that the later call to `AddDate()` accounts for daylight savings time traversal
	currentDate, err := types.ParseDateTime(startDateSeries.Time().In(location))
	if err != nil {
		return nil, err
	}

	switch mode {
	case Create:
		eventLinkedList.Init()
		var currentElement *list.Element

		for currentDate.Before(endDateSeries) {
			eventStart := currentDate
			eventEnd := currentDate.Add(time.Duration(seriesEventDuration) * time.Minute)

			event := &Event{}
			event.SetProxyRecord(core.NewRecord(eventCollection))

			setValuesForSeriesEvent(event, eventStart, eventEnd, eventSeries)

			if eventLinkedList.Len() == 0 {
				currentElement = eventLinkedList.PushFront(event)
			} else {
				currentElement = eventLinkedList.InsertAfter(event, currentElement)
			}

			currentDate = currentDate.AddDate(0, 0, seriesInterval)
		}
	case Update:
		// available because the hook runs before record persistence
		existingSeries := &EventSeries{}
		err = FindRecordProxyByID(app, existingSeries, eventSeries.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to find existing series: %w", err)
		}

		// TODO: does that capture all use cases?
		timeDiff := currentDate.Sub(existingSeries.SeriesStart())

		eventsToDelete := make(map[string]*Event)

		for element := eventLinkedList.Front(); element != nil; element = element.Next() {
			// MARK: no extra check for type assertion => delegated to list creation func
			event := element.Value.(*Event)

			var eventStart types.DateTime
			if event.StartTime().IsZero() {
				// just created from the previous loop, new event
				eventStart = currentDate
			} else {
				// has an existing DB value
				eventStart = event.StartTime().Add(timeDiff)
			}
			eventEnd := eventStart.Add(time.Duration(seriesEventDuration) * time.Minute)

			if eventStart.After(endDateSeries) {
				eventsToDelete[event.Id] = event
				eventLinkedList.Remove(element)
				continue
			}
			setValuesForSeriesEvent(event, eventStart, eventEnd, eventSeries)

			element.Value = event

			// series has been extended over the original end, append new event to handle in next loop
			if eventStart.Before(endDateSeries) {
				event := &Event{}
				event.SetProxyRecord(core.NewRecord(eventCollection))

				eventLinkedList.InsertAfter(event, element)
			}

			currentDate = currentDate.AddDate(0, 0, seriesInterval)
		}

		// Delete any leftover events that are no longer part of the updated series
		for _, staleEvent := range eventsToDelete {
			err := app.Delete(staleEvent)
			if err != nil {
				return nil, fmt.Errorf("failed to delete stale event: %w", err)
			}
		}
	}
	events, err := eventSeriesLinkedListToSlice(eventLinkedList)
	if err != nil {
		LogErrorInternalExternal(app, err, errorContext, nil)
	}

	return events, nil
}

func setValuesForSeriesEvent(event *Event, eventStart types.DateTime, eventEnd types.DateTime, eventSeries *EventSeries) {
	if event.Id == "" {
		event.Id = security.PseudorandomString(15)
	}
	event.SetStartTime(eventStart)
	event.SetEndTime(eventEnd)
	event.SetTitle(eventSeries.Title())
	event.SetTeam(eventSeries.Team())
	event.SetDesc(eventSeries.Desc())
	event.SetLocation(eventSeries.Location())
	event.SetSeries(eventSeries.Id)
	event.SetType(Practice.String())
}

// AddSeriesState is a hook that sets the state of the event series based on the current date
func AddSeriesState(e *core.RecordEnrichEvent) error {
	addSeriesState(e.Record)

	return e.Next()
}

// addSeriesState sets the state of the event series based on the current date
func addSeriesState(record *core.Record) {
	record.WithCustomData(true)
	eventSeries := &EventSeries{}
	eventSeries.SetProxyRecord(record)

	now := types.NowDateTime()

	if eventSeries.SeriesStart().After(now) {
		eventSeries.SetState(SeriesStateFuture)
	} else if eventSeries.SeriesEnd().Before(now) {
		eventSeries.SetState(SeriesStatePast)
	} else {
		eventSeries.SetState(SeriesStateOngoing)
	}
}
