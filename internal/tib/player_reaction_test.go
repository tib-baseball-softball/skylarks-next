package tib

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/core"
)

type mockLogOnlyApp struct{ logger *slog.Logger }

func (m *mockLogOnlyApp) Logger() *slog.Logger { return m.logger }

func buildTestRecord() *core.Record {
	collection := core.NewAuthCollection(dp.UserCollection, "usersColTest")
	record := core.NewRecord(collection)
	player := &dp.User{}
	player.SetProxyRecord(record)

	player.Id = "user123"
	player.SetFirstName("Alice")
	player.SetLastName("Doe")
	player.SetBSMID(777)
	player.SetNumber("42")
	player.SetThrows("R")
	player.SetBats("L")
	player.SetPosition([]string{"1", "2", "9"})
	player.SetScorer("C")
	player.SetUmpire("B")
	player.SetImage("pic.png")
	player.SetDisplayOnWebsite(true)

	return player.Record
}

func Test_sendUpdatedPlayerData_MockHTTPStatuses(t *testing.T) {
	type captured struct {
		Method    string
		Path      string
		Headers   http.Header
		Payload   PlayerReactionPayload
		WasCalled bool
	}

	scenarios := []struct {
		name       string
		status     int
		expectWarn bool
	}{
		{name: "200 OK", status: http.StatusOK, expectWarn: false},
		{name: "202 Accepted (unchanged)", status: http.StatusAccepted, expectWarn: false},
		{name: "400 Bad Request (invalid payload)", status: http.StatusBadRequest, expectWarn: true},
		// 500 could also occur but is explicitly out of scope/responsibility; behavior would be same as other non-2xx.
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			logBuf := bytes.Buffer{}
			logger := slog.New(slog.NewTextHandler(&logBuf, &slog.HandlerOptions{Level: slog.LevelInfo}))
			app := &mockLogOnlyApp{logger: logger}

			// channel to synchronize with the async goroutine
			done := make(chan struct{}, 1)

			captured := &captured{}
			identifier := "playerChangeTest"
			secret := "s3cr3t"

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					done <- struct{}{}
				}()

				captured.WasCalled = true
				captured.Method = r.Method
				captured.Path = r.URL.Path
				captured.Headers = r.Header.Clone()

				// basic path and headers checks
				if r.Method != http.MethodPost {
					t.Errorf("expected POST, got %s", r.Method)
				}
				expectedPath := "/typo3/reaction/" + identifier
				if r.URL.Path != expectedPath {
					t.Errorf("unexpected request path: got %s want %s", r.URL.Path, expectedPath)
				}
				if got := r.Header.Get("x-api-key"); got != secret {
					t.Errorf("missing or wrong x-api-key header: got %q", got)
				}
				if ct := r.Header.Get("Content-Type"); ct != "application/json" {
					t.Errorf("unexpected Content-Type: %q", ct)
				}
				if acc := r.Header.Get("Accept"); acc != "application/json" {
					t.Errorf("unexpected Accept: %q", acc)
				}
				if ua := r.Header.Get("User-Agent"); ua != "Skylarks Diamond Planner/1.0" {
					t.Errorf("unexpected User-Agent: %q", ua)
				}

				// decode body
				var p PlayerReactionPayload
				if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
					t.Errorf("failed to decode payload: %v", err)
				} else {
					captured.Payload = p
				}

				w.WriteHeader(sc.status)
			}))
			defer server.Close()

			// trigger the async sender
			rec := buildTestRecord()
			if err := sendUpdatedPlayerData(rec, app, server.URL, identifier, secret); err != nil {
				t.Fatalf("sendUpdatedPlayerData returned error: %v", err)
			}

			// wait for the goroutine to hit the server
			select {
			case <-done:
				// continue
			case <-time.After(2 * time.Second):
				t.Fatalf("timeout waiting for async HTTP call")
			}

			if !captured.WasCalled {
				t.Fatalf("mock server was not called")
			}

			// verify payload mapping
			if captured.Payload.PlayerId == "" || captured.Payload.PlayerId != rec.Id {
				t.Errorf("unexpected player id: %q", captured.Payload.PlayerId)
			}
			if captured.Payload.FirstName != "Alice" || captured.Payload.LastName != "Doe" {
				t.Errorf("unexpected name: %s %s", captured.Payload.FirstName, captured.Payload.LastName)
			}
			if captured.Payload.BSMID != 777 {
				t.Errorf("unexpected BSMID: %d", captured.Payload.BSMID)
			}
			if captured.Payload.Number != "42" {
				t.Errorf("unexpected number: %s", captured.Payload.Number)
			}
			// positions should be converted from []string{"1","2","9"} to []int{1,2,9}
			expectedPos := []int{1, 2, 9}
			if len(captured.Payload.Position) != len(expectedPos) {
				t.Fatalf("unexpected position length: got %d want %d", len(captured.Payload.Position), len(expectedPos))
			}
			for i, v := range expectedPos {
				if captured.Payload.Position[i] != v {
					t.Errorf("unexpected position[%d]: got %d want %d", i, captured.Payload.Position[i], v)
				}
			}
			if captured.Payload.Throwing != "R" || captured.Payload.Batting != "L" {
				t.Errorf("unexpected throwing/batting: %s/%s", captured.Payload.Throwing, captured.Payload.Batting)
			}
			if captured.Payload.Scorer != "C" || captured.Payload.Umpire != "B" {
				t.Errorf("unexpected scorer/umpire: %s/%s", captured.Payload.Scorer, captured.Payload.Umpire)
			}
			// image url should be collectionId + "/" + recordId + "/" + image
			expectedImageURL := rec.Collection().BaseFilesPath() + "/" + rec.Id + "/" + "pic.png"
			if captured.Payload.ImageURL != expectedImageURL {
				t.Errorf("unexpected ImageURL: got %q want %q", captured.Payload.ImageURL, expectedImageURL)
			}

			// wait for logger
			time.Sleep(1 * time.Second)

			logs := logBuf.String()
			if sc.expectWarn {
				if !strings.Contains(logs, "player change reaction returned non-2xx status") {
					t.Errorf("expected warning log for non-2xx status; logs: %s", logs)
				}
			} else {
				if strings.Contains(logs, "player change reaction returned non-2xx status") {
					t.Errorf("did not expect warning log for 2xx status; logs: %s", logs)
				}
			}
			if !strings.Contains(logs, "player change reaction sent") {
				t.Errorf("expected info log indicating request sent; logs: %s", logs)
			}
		})
	}
}
