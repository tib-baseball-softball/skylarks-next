package main

import (
	"net/http"
	"strings"
	"testing"

	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
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

type recordData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// These datasets need to exist in the testing database
var clubA = recordData{
	ID:   "vogf6j5vwmwrl6g",
	Name: "Club AAA",
}
var clubB = recordData{
	ID:   "zw96qr1u3fnz8kv",
	Name: "Club BBB",
}
var teamAlligators = recordData{
	ID:   "mydnx3d2bunjup2",
	Name: "Team Alligators",
}
var teamAntelopes = recordData{
	ID:   "ud80ykstikukfmo",
	Name: "Team Antelopes",
}
var teamBeavers = recordData{
	ID:   "3i7532nj13msbjg",
	Name: "Team Beavers",
}
var teamBees = recordData{
	ID:   "66prup8tpz8zekn",
	Name: "Team Bees",
}

var antelopeInternalEvent = recordData{
	ID:   "3h963ldke46yx7q",
	Name: "Antelope Internal",
}

var antelopeEventWithAdditional = recordData{
	ID:   "h8uai6gsbm951vk",
	Name: "Antelope External",
}

var alligatorEventForAntelopes = recordData{
	ID:   "8yxcf45qf1kzfbl",
	Name: "Alligator external",
}

var clubAEvent = recordData{
	ID:   "dvqfwl12270ol5l",
	Name: "Club A Event",
}

var beeInternalEvent = recordData{
	ID:   "i0kpukfmchkcyno",
	Name: "Bee Internal",
}

var beeEventWithAdditional = recordData{
	ID:   "pv5xv8jatjjq01k",
	Name: "Bee External",
}

var clubBEvent = recordData{
	ID:   "idwgcxk68p99o2s",
	Name: "Club B Event",
}

// Antelopes member
var alice = userData{
	ID:        "wvizeo6lht334n9",
	Email:     "alice_anderson@example.com",
	Username:  "alice_anderson",
	FirstName: "Alice",
	LastName:  "Anderson",
}

// Bees member
var bob = userData{
	ID:        "t63c8isheazpytq",
	Email:     "bob_brown@example.com",
	Username:  "bob_brown",
	FirstName: "Bob",
	LastName:  "Brown",
}
var bClubAdmin = userData{
	ID:        "54byrqk799uj8pj",
	Email:     "bernie_b_admin@example.com",
	Username:  "bernie_b_club_admin",
	FirstName: "Bernie",
	LastName:  "B Admin",
}

// team admin for the Beavers, but not the Bees
var bertaBeaverAdmin = userData{
	ID:        "edraypynmzd36gf",
	Email:     "berta_beaver@example.com",
	Username:  "berta_beaver",
	FirstName: "Berta",
	LastName:  "Beaver",
}

var aClubAdmin = userData{
	ID:        "8qjhl2zp4rt625r",
	Email:     "andy_a_admin@example.com",
	Username:  "andy_a_club_admin",
	FirstName: "Andy",
	LastName:  "A Admin",
}

// team admin for the Antelopes, but not the Alligators
var andreaAntelopeAdmin = userData{
	ID:        "mve4v09v47bdwqz",
	Email:     "andrea_antelope@example.com",
	Username:  "andrea_antelope",
	FirstName: "Andrea",
	LastName:  "Antelope",
}

func TestAPIRules(t *testing.T) {
	aliceToken, err := generateToken(dp.UserCollection, alice.Email)
	if err != nil {
		t.Fatal(err)
	}

	bobToken, err := generateToken(dp.UserCollection, bob.Email)
	if err != nil {
		t.Fatal(err)
	}

	aClubAdminToken, err := generateToken(dp.UserCollection, aClubAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	bClubAdminToken, err := generateToken(dp.UserCollection, bClubAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	andreaToken, err := generateToken(dp.UserCollection, andreaAntelopeAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	bertaToken, err := generateToken(dp.UserCollection, bertaBeaverAdmin.Email)
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
			ExpectedStatus:  http.StatusNotFound,
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
		{
			Name:   "User cannot self-assign to a team via teams+",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + alice.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:            strings.NewReader(`{"teams+":["` + teamAlligators.ID + `"]}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"message", "data"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User cannot self-assign to a club via clubs+",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + alice.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:            strings.NewReader(`{"club+" : "` + clubB.ID + `"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"message", "data"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin can rename own team (Antelopes)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamAntelopes.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body:            strings.NewReader(`{"name":"Team Antilopes_CHANGED"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"teams"`, `"name":"Team Antilopes_CHANGED"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin cannot rename other team in same club (Alligators)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamAlligators.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body:            strings.NewReader(`{"name":"Team Alligators_CHANGED"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin cannot rename team in another club (Bees)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamBees.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body:            strings.NewReader(`{"name":"Team Bees_CHANGED"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club A admin can rename teams in own club (Alligators)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamAlligators.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"name":"Team Alligators_CHANGED_BY_CLUB_A"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"teams"`, `"name":"Team Alligators_CHANGED_BY_CLUB_A"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club A admin can rename other team in own club (Antelopes)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamAntelopes.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"name":"Team Antelopes_CHANGED_BY_CLUB_A"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"teams"`, `"name":"Team Antelopes_CHANGED_BY_CLUB_A"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club A admin cannot rename team in Club B (Beavers)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamBeavers.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"name":"Team Beavers_CHANGED_BY_CLUB_A"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club B admin can rename team in own club (Bees)",
			Method: http.MethodPatch,
			URL:    "/api/collections/teams/records/" + teamBees.ID,
			Headers: map[string]string{
				"Authorization": bClubAdminToken,
			},
			Body:            strings.NewReader(`{"name":"Team Bees_CHANGED_BY_CLUB_B"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"teams"`, `"name":"Team Bees_CHANGED_BY_CLUB_B"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin (Andrea) can add Club A member (Club A Admin) to Antelopes",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + aClubAdmin.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body:            strings.NewReader(`{"teams+":["` + teamAntelopes.ID + `"]}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"users"`, `"teams":["` + teamAntelopes.ID + `"]`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin (Andrea) cannot add Club B member (Club B Admin) to Antelopes",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + bClubAdmin.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body:            strings.NewReader(`{"teams+":["` + teamAntelopes.ID + `"]}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin (Berta) can add Club B member (Club B Admin) to Beavers",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + bClubAdmin.ID,
			Headers: map[string]string{
				"Authorization": bertaToken,
			},
			Body:            strings.NewReader(`{"teams+":["` + teamBeavers.ID + `"]}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"users"`, `"teams":["` + teamBeavers.ID + `"]`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club A admin can edit user inside own club (Andrea)",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + andreaAntelopeAdmin.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"last_name":"Antelope_CHANGED"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"users"`, `"id":"` + andreaAntelopeAdmin.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club A admin cannot edit user in other club (Berta)",
			Method: http.MethodPatch,
			URL:    "/api/collections/users/records/" + bertaBeaverAdmin.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"last_name":"Beaver_CHANGED_BY_CLUB_A"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Nobody can delete clubs (superadmin only)",
			Method: http.MethodDelete,
			URL:    "/api/collections/clubs/records/" + clubA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"Only superusers can perform this action."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Member can see their own team event (Antelopes)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + antelopeInternalEvent.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"events"`, `"id":"` + antelopeInternalEvent.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Member can see an event with additional teams (Antelopes)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + antelopeEventWithAdditional.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"events"`, `"id":"` + antelopeEventWithAdditional.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Member can see another team's event that has additional_teams set",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + alligatorEventForAntelopes.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"events"`, `"id":"` + alligatorEventForAntelopes.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team admin can see their own team's event (Antelopes)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + antelopeInternalEvent.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"events"`, `"id":"` + antelopeInternalEvent.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club admin can see a club event (Club A)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + clubAEvent.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"collectionName":"events"`, `"id":"` + clubAEvent.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Member cannot see another team's event (Bob seeing Antelopes)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + antelopeInternalEvent.ID,
			Headers: map[string]string{
				"Authorization": bobToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club admin cannot see another club's event (Bernie seeing Club A)",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + clubAEvent.ID,
			Headers: map[string]string{
				"Authorization": bClubAdminToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Member cannot see an event from a different club",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + beeEventWithAdditional.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Team Admin cannot see an event from a different club",
			Method: http.MethodGet,
			URL:    "/api/collections/events/records/" + beeEventWithAdditional.ID,
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, s := range scenarios {
		s.Test(t)
	}
}
