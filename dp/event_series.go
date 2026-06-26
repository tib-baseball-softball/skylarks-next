package dp

import (
	"container/list"
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type SeriesState string

const (
	SeriesStateFuture  SeriesState = "future"
	SeriesStatePast    SeriesState = "past"
	SeriesStateOngoing SeriesState = "ongoing"
)

const (
	EventSeriesCollection = "eventseries"
)

var _ core.RecordProxy = (*EventSeries)(nil)

// EventSeries is a RecordProxy for the `eventseries` collection.
// It offers type-safe getters and setters for every field.
//
// Pure database model, relations are not resolved automatically.
type EventSeries struct {
	core.BaseRecordProxy
}

func (s *EventSeries) CollectionName() string {
	return EventSeriesCollection
}

// ID returns the record id.
func (s *EventSeries) ID() string {
	return s.Id
}

// Title returns the series title.
func (s *EventSeries) Title() string {
	return s.GetString("title")
}
func (s *EventSeries) SetTitle(title string) {
	s.Set("title", title)
}

// Team returns the id of the team the series belongs to.
func (s *EventSeries) Team() string {
	return s.GetString("team")
}
func (s *EventSeries) SetTeam(team string) {
	s.Set("team", team)
}

// Desc returns the optional description.
func (s *EventSeries) Desc() string {
	return s.GetString("desc")
}
func (s *EventSeries) SetDesc(desc string) {
	s.Set("desc", desc)
}

// Location returns the location string.
func (s *EventSeries) Location() string {
	return s.GetString("location")
}
func (s *EventSeries) SetLocation(location string) {
	s.Set("location", location)
}

// SeriesStart returns the first day/time of the series.
func (s *EventSeries) SeriesStart() types.DateTime {
	return s.GetDateTime("series_start")
}
func (s *EventSeries) SetSeriesStart(start types.DateTime) {
	s.Set("series_start", start)
}

// SeriesEnd returns the last day/time of the series.
func (s *EventSeries) SeriesEnd() types.DateTime {
	return s.GetDateTime("series_end")
}
func (s *EventSeries) SetSeriesEnd(end types.DateTime) {
	s.Set("series_end", end)
}

// Interval returns the number of days between occurrences of events in the series.
func (s *EventSeries) Interval() int {
	return s.GetInt("interval")
}
func (s *EventSeries) SetInterval(interval int) {
	s.Set("interval", interval)
}

// Duration returns the duration of an avent in the series (in minutes).
func (s *EventSeries) Duration() int {
	return s.GetInt("duration")
}
func (s *EventSeries) SetDuration(duration int) {
	s.Set("duration", duration)
}

func (s *EventSeries) SetState(state SeriesState) {
	s.Set("series_state", string(state))
}

// findEventRecordsForSeries fetches all events associated with a given eventSeries.
func findEventRecordsForSeries(app core.App, eventSeries *EventSeries) (events []*Event, firstEvent *Event, err error) {
	err = app.RecordQuery(EventsCollection).
		AndWhere(dbx.NewExp("series = {:seriesID}", dbx.Params{"seriesID": eventSeries.Id})).
		All(&events)
	if err != nil {
		return nil, nil, err
	}

	verifyCount := 0
	for _, event := range events {
		if event.Prev() == "" {
			verifyCount++
			firstEvent = event
		}
	}

	if verifyCount != 1 {
		return nil, nil, fmt.Errorf("logic: expected 1 first event, got %d", verifyCount)
	}

	return events, firstEvent, nil
}

// createEventSeriesLinkedList creates a linked list of events for given event records.
// 
// If an empty slice is given, the returned list will be empty as well.
// firstEvent will be set as first element of the list.
func createEventSeriesLinkedList(records []*Event, firstEvent *Event) (list list.List, err error) {
	list.Init()

	if len(records) == 0 {
		return list, nil
	}

	eventsMap := make(map[string]*Event)

	for _, event := range records {
		eventsMap[event.Id] = event
	}

	currentListElement := list.PushFront(firstEvent)
	for {
		currentEvent, ok := currentListElement.Value.(*Event)
		if !ok {
			return list, fmt.Errorf("data corrupted: event %v is not an Event", currentListElement.Value)
		}

		if currentEvent.Next() == "" {
			break
		}

		toInsert := eventsMap[currentEvent.Next()]
		if toInsert == nil {
			return list, fmt.Errorf("data corrupted: next event ID %s not found in map when creating linked list", currentEvent.Next())
		}

		currentListElement = list.InsertAfter(toInsert, currentListElement)
	}

	return list, nil
}
