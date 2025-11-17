package main

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
	"github.com/tib-baseball-softball/skylarks-next/internal/tib"
)

// Full integration scenarios inferred from the webhook handler.
// These require specific database fixtures (users with matching first/last names
// and predefined BSM ID values).
func TestWebhookEndpoint_IntegrationScenarios(t *testing.T) {
	identifier := "playerChangeTestIdentifier"
	secret := "superSecret123"

	// Expected DB states:
	// 1) Player with BSMID already set (should return 204)
	bodyAlreadySet := `{"playerUid":43,"bsmID":99999,"firstName":"Karl","lastName":"Testmichel"}`
	sigAlreadySet := tib.CalcWebhookSignature(identifier, []byte(bodyAlreadySet), secret)

	// 2) Player with BSMID = 0 in DB (should update and return 200)
	bodyNeedsUpdate := `{"playerUid":42,"bsmID":12345,"firstName":"Alex","lastName":"Sample"}`
	sigNeedsUpdate := tib.CalcWebhookSignature(identifier, []byte(bodyNeedsUpdate), secret)

	// 3) Player not found (should return 404)
	bodyNotFound := `{"playerUid":44,"bsmID":11111,"firstName":"DoesNot","lastName":"Exist"}`
	sigNotFound := tib.CalcWebhookSignature(identifier, []byte(bodyNotFound), secret)

	// 4) Malformed JSON with otherwise valid signature (should return 400)
	badJSON := `{`
	sigBadJSON := tib.CalcWebhookSignature(identifier, []byte(badJSON), secret)

	// Build scenarios
	scenarios := []tests.ApiScenario{
		{
			Name:            "Request with unset environment",
			Method:          http.MethodPost,
			URL:             "/webhooks/playerChange",
			ExpectedStatus:  http.StatusInternalServerError,
			ExpectedContent: []string{"Internal Server Error"},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", "")
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", "")
			},
		},
		{
			Name:            "Wrong request method",
			Method:          http.MethodGet,
			URL:             "/webhooks/playerChange",
			ExpectedStatus:  http.StatusNotFound, // TODO: Request never reaches my handler
			ExpectedContent: []string{"File not found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Missing signature header",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(bodyNeedsUpdate),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				// intentionally missing Webhook-Signature
			},
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"Invalid Webhook signature"},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
		},
		{
			Name:   "Invalid signature value",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(bodyNeedsUpdate),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				"Webhook-Signature":      "bogus-signature",
			},
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"Invalid Webhook signature"},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
		},
		{
			Name:   "Malformed JSON body",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(badJSON),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				"Webhook-Signature":      sigBadJSON,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Player Change Message does not match expected format"},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
		},
		{
			Name:   "Player not found",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(bodyNotFound),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				"Webhook-Signature":      sigNotFound,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"Player not found."},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
		},
		{
			Name:   "Player exists with BSMID already set -> 204",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(bodyAlreadySet),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				"Webhook-Signature":      sigAlreadySet,
			},
			ExpectedStatus: http.StatusNoContent,
			// no ExpectedContent to ensure empty body
			TestAppFactory: setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
		},
		{
			Name:   "Player exists with BSMID = 0 -> update and 200",
			Method: http.MethodPost,
			URL:    "/webhooks/playerChange",
			Body:   strings.NewReader(bodyNeedsUpdate),
			Headers: map[string]string{
				"Webhook-Signature-Algo": "sha256",
				"Webhook-Signature":      sigNeedsUpdate,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"Player BSM ID updated."},
			TestAppFactory:  setupTestApp,
			BeforeTestFunc: func(t testing.TB, app *tests.TestApp, e *core.ServeEvent) {
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER", identifier)
				_ = os.Setenv("PLAYER_CHANGE_WEBHOOK_SECRET", secret)
			},
			AfterTestFunc: func(t testing.TB, app *tests.TestApp, res *http.Response) {
				record, err := app.FindAuthRecordByEmail(dp.UserCollection, "alexsample@example.com")
				if err != nil {
					t.Fatalf("Could not find user %s", "alexsample@example.com")
				}
				player := &dp.User{}
				player.SetProxyRecord(record)

				player.SetBSMID(0)
				err = app.Save(player)
				if err != nil {
					t.Fatalf("Could not set BSM ID back to 0 for user %s", "alexsample@example.com")
				}
			},
		},
	}

	for _, s := range scenarios {
		s.Test(t)
	}
}
