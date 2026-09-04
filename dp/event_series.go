package dp

import (
	"container/list"
	"fmt"
	"time"

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

// AdditionalTeams returns the additional teams associated with the series.
func (s *EventSeries) AdditionalTeams() []string {
	return s.GetStringSlice("additional_teams")
}

func (s *EventSeries) SetAdditionalTeams(teams []string) {
	s.Set("additional_teams", teams)
}

func (s *EventSeries) AddAdditionalTeam(team string) {
	teams := append(s.AdditionalTeams(), team)
	s.SetAdditionalTeams(teams)
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

// Duration returns the duration of an event in the series (in minutes).
func (s *EventSeries) Duration() int {
	return s.GetInt("duration")
}
func (s *EventSeries) SetDuration(duration int) {
	s.Set("duration", duration)
}

func (s *EventSeries) SetState(state SeriesState) {
	s.WithCustomData(true)
	s.Set("series_state", string(state))
}

func (s *EventSeries) DetermineState() SeriesState {
	now := types.NowDateTime()

	if s.SeriesStart().After(now) {
		return SeriesStateFuture
	} else if s.SeriesEnd().Before(now) {
		return SeriesStatePast
	} else {
		return SeriesStateOngoing
	}
}

// PracticeSeason represents summer or winter in practice context.
type PracticeSeason uint8

// String returns a string representation of the underlying integer constant.
func (ps PracticeSeason) String() string {
	if ps == PracticeSeasonSummer {
		return "Summer"
	}
	return "Winter"
}

const (
	PracticeSeasonSummer PracticeSeason = 0
	PracticeSeasonWinter PracticeSeason = 1
)

// PracticeDTO represents a single event series in a calendaric format.
type PracticeDTO struct {
	TeamID         string         `json:"team_id"`
	Season         PracticeSeason `json:"season"`
	HumanSeason    string         `json:"human_season"`
	DayOfWeek      time.Weekday   `json:"day_of_week"`
	HumanDayOfWeek string         `json:"human_day_of_week"`
	StartTime      string         `json:"start_time"`
	EndTime        string         `json:"end_time"`
	Location       *LocationDTO   `json:"location"`
	Desc           string         `json:"desc"`
}

// ToPracticeDTO converts an event series to a PracticeDTO.
//
// The returned pointer is never nil.
func (s *EventSeries) ToPracticeDTO(loc *time.Location) *PracticeDTO {
	seriesStart := s.SeriesStart().Time() //.In(loc)
	endTimeFirstOccurence := seriesStart.Add((time.Duration(s.Duration()) * time.Minute))

	season := PracticeSeasonWinter
	if seriesStart.Month() >= time.April && seriesStart.Month() <= time.October {
		season = PracticeSeasonSummer
	}

	dto := &PracticeDTO{
		TeamID:         s.Team(),
		Season:         season,
		HumanSeason:    season.String(),
		DayOfWeek:      seriesStart.Weekday(),
		HumanDayOfWeek: seriesStart.Weekday().String(),
		StartTime:      seriesStart.Format(time.TimeOnly),
		EndTime:        endTimeFirstOccurence.Format(time.TimeOnly),
		Location:       nil,
		Desc:           s.Desc(),
	}
	locationRecord := s.ExpandedOne("location")
	if locationRecord != nil {
		eventLoc := &Location{}
		eventLoc.SetProxyRecord(locationRecord)
		dto.Location = eventLoc.ToDTO()
	}

	return dto
}

// findEventRecordsForSeries fetches all events associated with a given eventSeries.
func findEventRecordsForSeries(app core.App, eventSeries *EventSeries) (events []*Event, err error) {
	err = app.RecordQuery(EventsCollection).
		AndWhere(dbx.NewExp("series = {:seriesID}", dbx.Params{"seriesID": eventSeries.Id})).
		OrderBy("starttime ASC").
		All(&events)
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return events, nil
	}

	return events, nil
}

// createEventSeriesLinkedListFromDatabase creates a linked list of events for given event records.
//
// If an empty slice is given, the returned list will be empty as well.
// records needs to be already ordered.
//
// list will always point to a non-nil value unless an error occurs.
func createEventSeriesLinkedListFromSlice(records []*Event) (*list.List, error) {
	list := &list.List{}
	list.Init()

	if len(records) == 0 {
		return list, nil
	}

	currentListElement := list.PushFront(records[0])
	for i, event := range records {
		if i == 0 {
			continue
		}
		currentListElement = list.InsertAfter(event, currentListElement)
	}
	return list, nil
}

// EventSeriesLinkedListToSlice adds persistence info to an event series and returns a slice of events.
func EventSeriesLinkedListToSlice(list *list.List) (events []*Event, err error) {
	for element := list.Front(); element != nil; element = element.Next() {
		event, ok := element.Value.(*Event)
		if !ok {
			return nil, fmt.Errorf("data corrupted: event %v is not an Event pointer", element.Value)
		}

		events = append(events, event)
	}
	return events, nil
}
