package main

import (
	"net/http"
	"testing"

	"ganigeorgiev.com/pocketbase/pocketbase/tests"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
)

func TestGetUserStatsEndpoint(t *testing.T) {
	validToken, err := generateToken(base.UserCollection, "karltestmichel@example.com")
	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := generateToken(base.UserCollection, "otheruser@example.com")
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
			URL:    "/api/stats/" + validUserID,
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
			URL:    "/api/stats/" + validUserID,
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
