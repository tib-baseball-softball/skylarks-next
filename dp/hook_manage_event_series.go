package dp

import (
	"container/list"
	"fmt"
	"time"

	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/utils"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type EventSeriesMode string

const (
	CreateEventSeriesMode EventSeriesMode = "create"
	UpdateEventSeriesMode EventSeriesMode = "update"
)

type EventSeriesAction string

const (
	CreateEventsAction       EventSeriesAction = "create"
	UpdateFutureEventsAction EventSeriesAction = "update-future"
	UpdateAllEventsAction    EventSeriesAction = "update-all"
)

// CreateOrUpdateEventsForSeries is the main entry point for event series logic.
// Triggered on all Create and Update requests (regular API endpoints)
func CreateOrUpdateEventsForSeries(e *core.RecordRequestEvent, mode EventSeriesMode) error {
	action := EventSeriesAction(e.Request.URL.Query().Get("action"))
	if action == "" {
		action = UpdateAllEventsAction
	}

	events, err := generateSeriesEvents(e.App, e.Record, mode, action)
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

	eventsToBeDeleted, err := findEventRecordsForSeries(e.App, eventSeries)
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
func generateSeriesEvents(app core.App, record *core.Record, mode EventSeriesMode, action EventSeriesAction) ([]*Event, error) {
	var events []*Event
	eventSeries := &EventSeries{}
	eventSeries.SetProxyRecord(record)

	location, err := LoadAppTimeZone()
	if err != nil {
		return nil, err
	}

	// reads timezone information to ensure that the later call to `AddDate()` accounts for daylight savings time traversal
	// both calls ignore error: underlying method can only error on string casting, not when using time.Time arguments
	startDateSeries, _ := types.ParseDateTime(eventSeries.SeriesStart().Time().In(location))
	endDateSeries, _ := types.ParseDateTime(eventSeries.SeriesEnd().Time().In(location))

	if startDateSeries.After(endDateSeries) {
		return nil, fmt.Errorf("series start date is after series end date")
	}

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

	existingEvents := []*Event{}

	if mode == UpdateEventSeriesMode {
		existingEvents, err = findEventRecordsForSeries(app, eventSeries)
		if err != nil {
			LogErrorInternalExternal(app, err, errorContext, nil)
			return nil, err
		}
	}

	eventLinkedList, err := createEventSeriesLinkedListFromSlice(existingEvents)
	if err != nil {
		LogErrorInternalExternal(app, err, errorContext, nil)
		return nil, err
	}

	currentDate := startDateSeries
	today, _ := types.ParseDateTime(time.Now().In(location))

	switch mode {
	case CreateEventSeriesMode:
		eventLinkedList.Init()
		var currentElement *list.Element

		for currentDate.Before(endDateSeries) {
			eventStart := currentDate
			eventEnd := currentDate.Add(time.Duration(eventSeries.Duration()) * time.Minute)

			event := &Event{}
			event.SetProxyRecord(core.NewRecord(eventCollection))

			setValuesForSeriesEvent(event, eventStart, eventEnd, eventSeries)

			if eventLinkedList.Len() == 0 {
				currentElement = eventLinkedList.PushFront(event)
			} else {
				currentElement = eventLinkedList.InsertAfter(event, currentElement)
			}

			currentDate = currentDate.AddDate(0, 0, eventSeries.Interval())
		}
	case UpdateEventSeriesMode:
		// available because the hook runs before record persistence
		existingSeries := &EventSeries{}
		err = FindRecordProxyByID(app, existingSeries, eventSeries.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to find existing series: %w", err)
		}

		eventsToDelete := make(map[string]*Event)

		for element := eventLinkedList.Front(); element != nil; element = element.Next() {
			app.Logger().Debug("processing event", "element", element)

			// process the event for this loop
			event, ok := element.Value.(*Event)
			if !ok {
				// this really should not happen, but better be safe than sorry
				err = fmt.Errorf("data corrupted: event %v is not an Event pointer", element.Value)
				LogErrorInternalExternal(app, err, errorContext, nil)
				return nil, err
			}

			var eventStart types.DateTime
			if event.StartTime().IsZero() {
				// just created from the previous loop, new event
				eventStart = currentDate
			} else {
				// has an existing DB value
				timeDiff := currentDate.Sub(event.StartTime())
				eventStart = event.StartTime().Add(timeDiff)
			}
			eventEnd := eventStart.Add(time.Duration(eventSeries.Duration()) * time.Minute)

			if eventStart.After(endDateSeries) {
				eventsToDelete[event.Id] = event
				deleting := eventLinkedList.Remove(element)
				app.Logger().Debug("deleting event", "event", deleting)
				app.Logger().Debug("end of update loop", "list length", eventLinkedList.Len())
				continue
			}

			if action == UpdateAllEventsAction || eventStart.After(today) {
				// only update future events if set
				setValuesForSeriesEvent(event, eventStart, eventEnd, eventSeries)
				element.Value = event
			}

			currentDate = currentDate.AddDate(0, 0, eventSeries.Interval())
			
			// series has been extended over the original end, append new event to handle in next loop
			if element.Next() == nil && currentDate.Before(endDateSeries) {
				app.Logger().Debug("appending new event to extend series", "currentDate", currentDate, "endDateSeries", endDateSeries)

				event := &Event{}
				event.SetProxyRecord(core.NewRecord(eventCollection))
				eventLinkedList.InsertAfter(event, element)
			}
			app.Logger().Debug("end of update loop", "list length", eventLinkedList.Len())
		}

		// Delete any leftover events that are no longer part of the updated series
		for _, staleEvent := range eventsToDelete {
			err := app.Delete(staleEvent)
			if err != nil {
				return nil, fmt.Errorf("failed to delete stale event: %w", err)
			}
		}
	}
	events, err = EventSeriesLinkedListToSlice(eventLinkedList)
	if err != nil {
		LogErrorInternalExternal(app, err, errorContext, nil)
	}
	eventLinkedList = nil // explicit pointer reset after use

	return events, nil
}

func setValuesForSeriesEvent(event *Event, eventStart types.DateTime, eventEnd types.DateTime, eventSeries *EventSeries) {
	if event.Id == "" {
		event.Id = utils.CreatePocketBaseIDString()
	}
	if eventSeries.Id == "" {
		eventSeries.Id = utils.CreatePocketBaseIDString()
	}
	event.SetStartTime(eventStart)
	event.SetEndTime(eventEnd)
	event.SetTitle(eventSeries.Title())
	event.SetTeam(eventSeries.Team())
	event.SetAdditionalTeams(eventSeries.AdditionalTeams())
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
