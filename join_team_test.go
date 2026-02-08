package main

import (
	"net/http"
	"strings"
	"testing"

	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
)

func TestJoinTeam_Integration(t *testing.T) {
	testApp := setupTestApp(t)
	defer testApp.Cleanup()

	userToken, err := generateToken(dp.UserCollection, alice.Email)
	if err != nil {
		t.Fatal(err)
	}

	signupKey := "random-gators"

	scenarios := []tests.ApiScenario{
		{
			Name:   "User joins team with valid signup key",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAlligators.ID + "/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "wvizeo6lht334n9",
				"signupKey": "` + signupKey + `"
			}`),
			ExpectedStatus: http.StatusOK,
			ExpectedContent: []string{
				`"message":"User joined team successfully"`,
			},
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "User tries to join team with invalid signup key",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAlligators.ID + "/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "wvizeo6lht334n9",
				"signupKey": "wrong_key"
			}`),
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"Provided signup key does not match team signup key"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User tries to join team with mismatched user ID",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAlligators.ID + "/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "t63c8isheazpytq",
				"signupKey": "` + signupKey + `"
			}`),
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"User ID does not match authenticated user"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User tries to join team they're already a member of",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAntelopes.ID + "/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "wvizeo6lht334n9",
				"signupKey": "` + signupKey + `"
			}`),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"User is already a member of this team"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User joins team but is not authenticated",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAlligators.ID + "/join",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "wvizeo6lht334n9",
				"signupKey": "` + signupKey + `"
			}`),
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"The request requires valid record authorization token."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User joins team with malformed JSON",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/" + teamAlligators.ID + "/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body:            strings.NewReader(`{`),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Supplied data is malformed"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User joins non-existent team",
			Method: http.MethodPost,
			URL:    "/api/dp/teams/nonexistentteam/join",
			Headers: map[string]string{
				"Authorization": userToken,
				"Content-Type":  "application/json",
			},
			Body: strings.NewReader(`{
				"userID": "wvizeo6lht334n9",
				"signupKey": "` + signupKey + `"
			}`),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to load teams record"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, s := range scenarios {
		s.Test(t)
	}
}
