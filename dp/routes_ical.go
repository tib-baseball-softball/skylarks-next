package dp

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/security"
)

type NotAuthorizedError struct {
	Message string
}

func (e *NotAuthorizedError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func getUserCalendar(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		user := &User{}
		err := FindRecordProxyByID(app, user, e.Request.PathValue("user"))
		if err != nil {
			return e.BadRequestError("No valid user given.", "")
		}

		appSecret := os.Getenv("APPLICATION_SECRET")
		givenHash := e.Request.PathValue("hash")
		calculatedHash := GetSHA3Hash(user.Id + appSecret)

		if !security.Equal(calculatedHash, givenHash) {
			return e.BadRequestError("Invalid hash.", "")
		}

		teamQuery := e.Request.URL.Query().Get("team")

		events, err := getUserCalendarEvents(app, user, teamQuery)
		if err != nil {
			if _, ok := errors.AsType[*NotAuthorizedError](err); ok {
				return e.UnauthorizedError("", "")
			} else {
				app.Logger().Error("Failed to fetch events for ical", "err", err, "user", user.Id)
				return e.InternalServerError("Failed to fetch events.", "")
			}
		}

		calendar := createICalendarFromEventRecords(events)

		e.Response.Header().Set("Content-Disposition", "inline")
		e.Response.Header().Set("filename", "calendar.ics")
		return e.Blob(http.StatusOK, "text/calendar", []byte(calendar))
	}
}

func getUserCalendarEvents(app core.App, user *User, teamID string) ([]Event, error) {
	var events []Event
	expressions := []dbx.Expression{}

	if teamID != "" {
	    // filter by specific team
		if !slices.Contains(user.Teams(), teamID) {
			return events, &NotAuthorizedError{"Attempting to filter by team not in user teams"}
		}
		expressions = append(expressions, dbx.HashExp{"team": teamID})
	} else {
	    // filter by all user teams
		for _, userTeam := range user.Teams() {
			expressions = append(expressions, dbx.HashExp{"team": userTeam})
		}
	}

	// strftime() requires strings for correct comparisons
	prevYear := strconv.Itoa(time.Now().AddDate(-1, 0, 0).Year())

	err := app.RecordQuery(EventsCollection).
		AndWhere(dbx.Or(expressions...)).
		AndWhere(dbx.NewExp("strftime('%Y', starttime) >= {:prevYear}", dbx.Params{"prevYear": prevYear})).
		All(&events)

	return events, err
}

func createICalendarFromEventRecords(events []Event) string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, eventRecord := range events {
		calEvent := cal.AddEvent(eventRecord.Id)
		calEvent.SetCreatedTime(eventRecord.Created().Time())
		calEvent.SetDtStampTime(time.Now())
		calEvent.SetModifiedAt(eventRecord.Updated().Time())

		calEvent.SetStartAt(eventRecord.StartTime().Time())

		endTime := eventRecord.EndTime().Time()
		if endTime.IsZero() {
			// end time is not a required field in our database, but calendar events need one
			endTime = eventRecord.StartTime().Time().Add(2 * time.Hour)
		}
		calEvent.SetEndAt(endTime)

		calEvent.SetSummary(eventRecord.Title())
		calEvent.SetDescription(eventRecord.Desc())
		calEvent.SetLocation(eventRecord.Location()) // TODO location expand
	}

	return cal.Serialize()
}
