package dp

import (
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

type EventType string

const (
	EventTypeGame     EventType = "game"
	EventTypePractice EventType = "practice"
	EventTypeMisc     EventType = "misc"
	EventTypeAll      EventType = "all"
)

func (et EventType) String() string {
	return string(et)
}

func ParseEventType(s string) (EventType, error) {
	switch EventType(s) {
	case EventTypeGame, EventTypePractice, EventTypeMisc, EventTypeAll:
		return EventType(s), nil
	default:
		return "", fmt.Errorf("invalid EventType: %s", s)
	}
}

const (
	EventsCollection = "events"
)

var _ core.RecordProxy = (*Event)(nil)

// Event RecordProxy for collection `events`.
// Provides type-safe struct access to event records.
type Event struct {
	core.BaseRecordProxy
}

func (e *Event) CollectionName() string {
	return EventsCollection
}

func (e *Event) ID() string {
	return e.Id
}

func (e *Event) Created() types.DateTime {
	return e.GetDateTime("created")
}

func (e *Event) Updated() types.DateTime {
	return e.GetDateTime("updated")
}

func (e *Event) BSMID() int {
	return e.GetInt("bsm_id")
}

func (e *Event) SetBSMID(bsmID int) {
	e.Set("bsm_id", bsmID)
}

func (e *Event) Title() string {
	return e.GetString("title")
}

func (e *Event) SetTitle(title string) {
	e.Set("title", title)
}

func (e *Event) Desc() string {
	return e.GetString("desc")
}

func (e *Event) SetDesc(desc string) {
	e.Set("desc", desc)
}

func (e *Event) StartTime() types.DateTime {
	return e.GetDateTime("starttime")
}

func (e *Event) SetStartTime(startTime types.DateTime) {
	e.Set("starttime", startTime)
}

func (e *Event) EndTime() types.DateTime {
	return e.GetDateTime("endtime")
}

func (e *Event) SetEndTime(endTime types.DateTime) {
	e.Set("endtime", endTime)
}

func (e *Event) MeetingTime() types.DateTime {
	return e.GetDateTime("meetingtime")
}

func (e *Event) SetMeetingTime(meetingTime types.DateTime) {
	e.Set("meetingtime", meetingTime)
}

// Type returns the event type enum:
// - Game
// - Practice
// - Misc
// 
// Returns `Misc` if parsing the DB value fails.
func (e *Event) Type() EventType {
	t, err := ParseEventType(e.GetString("type"))
	if err != nil {
	    return EventTypeMisc
	}
	return t
}

// SetType sets the event type.
// 
// Caution: `All` is not a valid value in database context, will be replaced by `Misc`.
func (e *Event) SetType(eventType EventType) {
    if eventType == EventTypeAll {
        e.Set("type", EventTypeMisc.String())
        return
    }
	e.Set("type", eventType.String())
}

func (e *Event) Team() string {
	return e.GetString("team")
}

func (e *Event) SetTeam(team string) {
	e.Set("team", team)
}

func (e *Event) AdditionalTeams() []string {
	return e.GetStringSlice("additional_teams")
}

func (e *Event) SetAdditionalTeams(additionalTeams []string) {
	e.Set("additional_teams", additionalTeams)
}

func (e *Event) AddAdditionalTeam(team string) {
	e.SetAdditionalTeams(append(e.AdditionalTeams(), team))
}

func (e *Event) HasAdditionalTeams() bool {
	return len(e.AdditionalTeams()) > 0
}

func (e *Event) Club() string {
	return e.GetString("club")
}

func (e *Event) SetClub(club string) {
	e.Set("club", club)
}

func (e *Event) Location() string {
	return e.GetString("location")
}

func (e *Event) SetLocation(location string) {
	e.Set("location", location)
}

func (e *Event) MatchJSON() string {
	return e.GetString("match_json")
}

func (e *Event) SetMatchJSON(matchJSON string) {
	e.Set("match_json", matchJSON)
}

func (e *Event) Guests() string {
	return e.GetString("guests")
}

func (e *Event) SetGuests(guests string) {
	e.Set("guests", guests)
}

func (e *Event) IsCancelled() bool {
	return e.GetBool("cancelled")
}

func (e *Event) SetCancelled(cancelled bool) {
	e.Set("cancelled", cancelled)
}

func (e *Event) Attire() string {
	return e.GetString("attire")
}

func (e *Event) SetAttire(attire string) {
	e.Set("attire", attire)
}

func (e *Event) Series() string {
	return e.GetString("series")
}

func (e *Event) SetSeries(series string) {
	e.Set("series", series)
}

// SetParticipations sets a virtual (non-DB) field
// to include participation data (back-relation).
func (e *Event) SetParticipations(participations ParticipationsByType) {
	e.WithCustomData(true)
	e.Set("participations", participations)
}

// SetUserParticipation sets a virtual (non-DB) field to include
// participation data for the currently-logged-in user.
func (e *Event) SetUserParticipation(participation *core.Record) {
	e.WithCustomData(true)
	e.Set("userParticipation", participation)
}
