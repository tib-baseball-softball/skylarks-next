package pb

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	EventsCollection = "events"
)

var _ core.RecordProxy = (*Events)(nil)

// Events RecordProxy for collection `events`.
// Provides type-safe struct access to event records.
type Events struct {
	core.BaseRecordProxy
}

func (e *Events) ID() string {
	return e.Id
}

func (e *Events) BSMID() int {
	return e.GetInt("bsm_id")
}

func (e *Events) SetBSMID(bsmID int) {
	e.Set("bsm_id", bsmID)
}

func (e *Events) Title() string {
	return e.GetString("title")
}

func (e *Events) SetTitle(title string) {
	e.Set("title", title)
}

func (e *Events) StartTime() types.DateTime {
	return e.GetDateTime("starttime")
}

func (e *Events) SetStartTime(startTime types.DateTime) {
	e.Set("starttime", startTime)
}

func (e *Events) EndTime() types.DateTime {
	return e.GetDateTime("endtime")
}

func (e *Events) SetEndTime(endTime types.DateTime) {
	e.Set("endtime", endTime)
}

func (e *Events) MeetingTime() types.DateTime {
	return e.GetDateTime("meetingtime")
}

func (e *Events) SetMeetingTime(meetingTime types.DateTime) {
	e.Set("meetingtime", meetingTime)
}

func (e *Events) Type() string {
	return e.GetString("type")
}

func (e *Events) SetType(eventType string) {
	e.Set("type", eventType)
}

func (e *Events) Team() string {
	return e.GetString("team")
}

func (e *Events) SetTeam(team string) {
	e.Set("team", team)
}

func (e *Events) Location() string {
	return e.GetString("location")
}

func (e *Events) SetLocation(location string) {
	e.Set("location", location)
}

func (e *Events) MatchJSON() string {
	return e.GetString("match_json")
}

func (e *Events) SetMatchJSON(matchJSON string) {
	e.Set("match_json", matchJSON)
}

func (e *Events) Guests() string {
	return e.GetString("guests")
}

func (e *Events) SetGuests(guests string) {
	e.Set("guests", guests)
}

func (e *Events) IsCancelled() bool {
	return e.GetBool("cancelled")
}

func (e *Events) SetCancelled(cancelled bool) {
	e.Set("cancelled", cancelled)
}

func (e *Events) Attire() string {
	return e.GetString("attire")
}

func (e *Events) SetAttire(attire string) {
	e.Set("attire", attire)
}

func (e *Events) Series() string {
	return e.GetString("series")
}

func (e *Events) SetSeries(series string) {
	e.Set("series", series)
}
