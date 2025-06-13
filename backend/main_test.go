package main

import (
	"github.com/pocketbase/pocketbase/tests"
	"net/http"
	"testing"
)

const (
	testDataDir      = "./test_pb_data"
	testTeamID       = "847u6a1af1loox5"
	seasonQueryParam = "2025"
	validUserId      = "jm866jti5piq9g9"
)

func generateToken(collectionNameOrId string, email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.FindAuthRecordByEmail(collectionNameOrId, email)
	if err != nil {
		return "", err
	}

	return record.NewAuthToken()
}

func setupTestApp(t testing.TB) *tests.TestApp {
	testApp, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}
	// no need to clean up since scenario.Test() will do that for us

	bindAppHooks(testApp)

	return testApp
}

func TestHelloEndpoint(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "try with different http method, e.g. POST",
			Method:          http.MethodPost,
			URL:             "/api/hello",
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"wasn't found"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "any auth status with GET",
			Method:          http.MethodGet,
			URL:             "/api/hello",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"Hello ballplayers!"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGamesCountEndpoint(t *testing.T) {
	validToken, err := generateToken("users", "karltestmichel@example.com")
	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := generateToken("users", "otheruser@example.com")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "Unauthenticated",
			Method:          http.MethodGet,
			URL:             "/api/gamecount/" + testTeamID,
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"data\":{},\"message\":\"The request requires valid record authorization token."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Invalid token",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID,
			Headers: map[string]string{
				"Authorization": invalidToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"Only members and admins can query team game count"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Bogus Team",
			Method: http.MethodGet,
			URL:    "/api/gamecount/doesnotexist",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to load teams record"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid user and team without query",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Query param for season is required"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid user and team with query",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID + "?season=" + seasonQueryParam,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"count\":0"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetLeagueLeadersEndpoint(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "Invalid league ID",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/abc",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"league ID must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Valid request with default statsType",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/5732",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_id", "stats_type", "data"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Valid request with specified statsType",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/5732?statsType=pitching",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_id", "stats_type", "data"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestGetUserStatsEndpoint(t *testing.T) {
	validToken, err := generateToken("users", "karltestmichel@example.com")
	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := generateToken("users", "otheruser@example.com")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "Unauthenticated request",
			Method:          http.MethodGet,
			URL:             "/api/stats/123",
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Invalid token",
			Method: http.MethodGet,
			URL:    "/api/stats/" + validUserId,
			Headers: map[string]string{
				"Authorization": invalidToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User not found",
			Method: http.MethodGet,
			URL:    "/api/stats/doesnotexist",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid request",
			Method: http.MethodGet,
			URL:    "/api/stats/" + validUserId,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"season", "totalPossibleEvents", "attendanceTotals", "participationTotals", "teamName"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}

func TestFavoriteTeamEndpoint(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "try with different http method, e.g. POST",
			Method:          http.MethodPost,
			URL:             "/api/team/favorite",
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"wasn't found"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "missing team parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?season=2025",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Team ID not found in request"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "invalid team parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=invalid&season=2025",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Team ID must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "missing season parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Season not found in request"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "invalid season parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219&season=invalid",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Season must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "valid request",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219&season=2025",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_group", "playoff_series", "table", "table_row"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
