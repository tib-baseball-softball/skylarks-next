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
