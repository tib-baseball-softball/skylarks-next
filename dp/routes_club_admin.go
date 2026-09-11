package dp

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/pocketbase/pocketbase/tools/types"
)

const ClubAdminAccessMiddlewareId = "clubAdminAccess"
const EventStoreKeyClub = "club"

// RequireClubAdminAccess verifies that the current auth is an admin for the club.
// The club is expected to be provided via the path parameter `club`.
func RequireClubAdminAccess() *hook.Handler[*core.RequestEvent] {
	return &hook.Handler[*core.RequestEvent]{
		Id:   ClubAdminAccessMiddlewareId,
		Func: requireClubAdminAccess(),
	}
}

func requireClubAdminAccess() func(*core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		if e.Auth == nil {
			return e.UnauthorizedError("", nil)
		}

		clubID := e.Request.PathValue("club")
		if clubID == "" {
			return e.BadRequestError("club ID is required", nil)
		}

		clubRecord, err := e.App.FindRecordById(ClubsCollection, clubID)
		if err != nil {
			return e.NotFoundError("", err)
		}
		club := Club{}
		club.SetProxyRecord(clubRecord)

		if slices.Contains(club.Admins(), e.Auth.Id) {
			e.Set(EventStoreKeyClub, club)
			return e.Next()
		}
		return e.UnauthorizedError("", nil)
	}
}

func GetClubCommunityService() func(event *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		club := e.Get(EventStoreKeyClub).(Club)

		var season int
		var err error

		seasonParam := e.Request.URL.Query().Get("season")
		if seasonParam != "" {
			season, err = strconv.Atoi(seasonParam)
			if err != nil {
				e.App.Logger().Error("failed to parse season: %v", "error", err)
				return err
			}
		} else {
			season = types.NowDateTime().Time().Year()
		}

		rows, err := getCommunityServiceDataForClub(e.App, club.Id, strconv.Itoa(season))
		if err != nil {
			return e.InternalServerError("", err)
		}
		return e.JSON(http.StatusOK, rows)
	}
}

// HandleClubPractices is a handler func that returns all event series
// for a single club in a format suitable for external systems.
//
// Authentication is expected to be handled by a preceding middleware.
func HandleClubPractices() func(e *core.RequestEvent) error {
	return handleClubPractices
}

func handleClubPractices(e *core.RequestEvent) error {
	club, ok := e.Get(EventStoreKeyClub).(Club)
	ctx := &ErrorContext{
		Key: "local",
		Values: map[string]any{
			"clubID":   club.Id,
			"clubName": club.Name(),
		},
	}
	if !ok {
		err := fmt.Errorf("Club not properly set in request context")
		ForwardErrorToExternalSystem(err, ctx, nil)
		return e.InternalServerError("", err)
	}

	seriesRecords, err := e.App.FindRecordsByFilter(
		EventSeriesCollection,
		"series_start <= @now && series_end >= @now && team.club.id = {:clubID}",
		"+series_start",
		0,
		0,
		dbx.Params{
			"clubID": club.Id,
		},
	)
	if err != nil {
		ForwardErrorToExternalSystem(err, ctx, nil)
		return e.InternalServerError("", err)
	}
	errs := e.App.ExpandRecords(seriesRecords, []string{"location"}, nil)
	if len(errs) > 0 {
		return fmt.Errorf("failed to expand: %v", errs)
	}

	location, err := LoadAppTimeZone()
	if err != nil {
		ForwardErrorToExternalSystem(err, ctx, nil)
		return e.InternalServerError("", err)
	}

	var dtos []*PracticeDTO

	for _, record := range seriesRecords {
		s := &EventSeries{}
		s.SetProxyRecord(record)
		dtos = append(dtos, s.ToPracticeDTO(location))
	}

	return e.JSON(http.StatusOK, dtos)
}
