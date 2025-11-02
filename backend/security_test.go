package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
)

type userData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TeamsPlus string `json:"teams+"`
	ClubsPlus string `json:"clubs+"`
}

// Those users need to exist in the testing database
var alice = userData{
	ID:        "wvizeo6lht334n9",
	Email:     "alice_anderson@example.com",
	Username:  "alice_anderson",
	FirstName: "Alice",
	LastName:  "Anderson",
}
var bob = userData{
	ID:        "t63c8isheazpytq",
	Email:     "bob_brown@example.com",
	Username:  "bob_brown",
	FirstName: "Bob",
	LastName:  "Brown",
}

func TestAPIRules(t *testing.T) {
	aliceToken, err := generateToken(dp.UserCollection, alice.Email)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "Change user without authorization token",
			Method:          http.MethodPatch,
			URL:             "/api/collections/users/records/" + alice.ID,
			ExpectedStatus:  http.StatusNotFound, // PB does not disclose whether a user exists
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Change other user with own authorization token",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + bob.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusNotFound, // PB does not disclose whether a user exists
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Change user(self) with own valid authorization token",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + alice.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:           strings.NewReader(`{"last_name":"Anderson_CHANGED"}`),
			ExpectedStatus: http.StatusOK,
			ExpectedContent: []string{
				`"collectionName":"users"`,
				`"email":"alice_anderson@example.com"`,
			},
			TestAppFactory: setupTestApp,
		},
	}

	for _, s := range scenarios {
		s.Test(t)
	}
}
