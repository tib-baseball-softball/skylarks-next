package main

import (
	"net/http"
	"testing"

	"ganigeorgiev.com/pocketbase/pocketbase/tests"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func TestPracticeExport(t *testing.T) {
	validToken, err := generateToken(base.UserCollection, bClubAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := generateToken(base.UserCollection, bertaBeaverAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:               "Unauthenticated",
			Method:             http.MethodGet,
			URL:                "/api/clubs/" + clubB.ID + "/admin/practices",
			ExpectedStatus:     http.StatusUnauthorized,
			NotExpectedContent: []string{"[]"},
			TestAppFactory:     setupTestApp,
		},
		{
			Name:   "Invalid token",
			Method: http.MethodGet,
			URL:    "/api/clubs/" + clubB.ID + "/admin/practices",
			Headers: map[string]string{
				"Authorization": invalidToken,
			},
			ExpectedStatus:     http.StatusUnauthorized,
			NotExpectedContent: []string{"[]"},
			TestAppFactory:     setupTestApp,
		},
		{
			Name:   "Not a club",
			Method: http.MethodGet,
			URL:    "/api/clubs/mostCertainlyNotATeam/admin/practices",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:     http.StatusNotFound,
			NotExpectedContent: []string{"[]"},
			TestAppFactory:     setupTestApp,
		},
		{
			Name:   "Valid",
			Method: http.MethodGet,
			URL:    "/api/clubs/" + clubB.ID + "/admin/practices",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"[]"}, // test DB has no current series
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
