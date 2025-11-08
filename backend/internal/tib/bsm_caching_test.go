package tib

import (
	"errors"
	"io"
	"log/slog"
	"net/url"
	"testing"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
)

type MockApp struct {
	findFirstRecordByDataFunc    func(any, string, any) (*core.Record, error)
	findCollectionByNameOrIdFunc func(string) (*core.Collection, error)
	saveFunc                     func(core.Model) error
}

func (m *MockApp) Delete(model core.Model) error {
	return nil
}

func (m *MockApp) Logger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}

func (m *MockApp) FindFirstRecordByData(collectionModelOrIdentifier any, field string, value any) (*core.Record, error) {
	return m.findFirstRecordByDataFunc(collectionModelOrIdentifier, field, value)
}

func (m *MockApp) FindCollectionByNameOrId(nameOrId string) (*core.Collection, error) {
	return m.findCollectionByNameOrIdFunc(nameOrId)
}

func (m *MockApp) Save(record core.Model) error {
	return m.saveFunc(record)
}

func TestIsValidBSMURL(t *testing.T) {
	tests := []struct {
		rawURL string
		want   bool
	}{
		// Valid URLs
		{"https://bsm.baseball-softball.de/clubs/485/licenses.json", true},
		{"https://bsm.baseball-softball.de/clubs/123456/licenses.json", true},
		{"https://bsm.baseball-softball.de/clubs/987/licenses.json?param=value", true},
		{"https://bsm.baseball-softball.de/clubs/485.json", true},
		{"https://bsm.baseball-softball.de/clubs/999999.json?key=val", true},
		{"https://bsm.baseball-softball.de/league_groups/5732/details", true},
		{"https://bsm.baseball-softball.de/league_groups/5732/standings?season=2024", true},
		{"https://bsm.baseball-softball.de/league_groups/8888/stats", true},
		{"https://bsm.baseball-softball.de/league_groups/5732.json", true},
		{"https://bsm.baseball-softball.de/licenses/999999.json", true},
		{"https://bsm.baseball-softball.de/clubs/123456/club_functions.json?filter=active", true},
		{"https://bsm.baseball-softball.de/clubs/485/fields.json", true},
		{"https://bsm.baseball-softball.de/clubs/123456/fields.json?sort=desc", true},

		// Invalid URLs
		{"https://bsm.baseball-softball.de/clubs/abc/licenses.json", false},
		{"https://bsm.baseball-softball.de/clubs/485/licenses.xml", false},
		{"https://example.com/clubs/485/licenses.json", false},
		{"https://bsm.baseball-softball.de/clubs/485", false},
		{"https://bsm.baseball-softball.de/league_groups/abc/details", false},
		{"https://bsm.baseball-softball.de/clubs/485/player_lists.json", false},
		{"https://bsm.baseball-softball.de/people/123.json", false},
		{"https://bsm.baseball-softball.de/licenses/abcd.json", false},
		{"https://bsm.baseball-softball.de/licenses/6000.xml", false},
		{"https://bsm.baseball-softball.de/clubs/abc/club_functions.json", false},
		{"https://bsm.baseball-softball.de/clubs/485/club_function.json", false},
		{"https://bsm.baseball-softball.de/clubs/abc/fields.json", false},
		{"https://bsm.baseball-softball.de/clubs/485/fields.xml", false},
		{"https://bsm.baseball-softball.de/clubs/485/field.json", false},
	}

	for _, tt := range tests {
		t.Run(tt.rawURL, func(t *testing.T) {
			parsedURL, _ := url.Parse(tt.rawURL)
			got := isValidBSMURL(parsedURL)
			if got != tt.want {
				t.Errorf("isValidBSMURL(%q) = %v; want %v", tt.rawURL, got, tt.want)
			}
		})
	}
}

// helper that must never fail inside the test body
func mustParseDateTime(t *testing.T, tm time.Time) types.DateTime {
	t.Helper()
	parsed, err := types.ParseDateTime(tm)
	if err != nil {
		t.Fatalf("failed to parse time %v: %v", tm, err)
	}
	return parsed
}

// builds an in-memory RequestCache record for a given raw URL and body
func buildCacheRecord(t *testing.T, rawURL, body string, updatedAt time.Time) *core.Record {
	t.Helper()
	coll := core.NewBaseCollection(dp.RequestCacheCollection)
	rec := core.NewRecord(coll)
	cache := &dp.RequestCache{}
	cache.SetProxyRecord(rec)
	cache.SetURL(rawURL)
	cache.SetResponseBody(body)
	cache.SetHash(dp.GetMD5Hash(rawURL))
	cache.SetUpdated(mustParseDateTime(t, updatedAt))
	return rec
}

func TestGetCachedBSMResponse_CacheHit_ReturnsCachedBody(t *testing.T) {
	t.Setenv("BSM_API_HOST", "bsm.baseball-softball.de")

	raw := "https://bsm.baseball-softball.de/clubs/485/licenses.json"
	u, _ := url.Parse(raw)

	// fresh record (not outdated)
	rec := buildCacheRecord(t, raw, "cached-body", time.Now().Add(-5*time.Minute))

	app := &MockApp{
		findFirstRecordByDataFunc:    func(any, string, any) (*core.Record, error) { return rec, nil },
		findCollectionByNameOrIdFunc: func(string) (*core.Collection, error) { return core.NewBaseCollection(dp.RequestCacheCollection), nil },
		saveFunc:                     func(core.Model) error { return nil },
	}

	got, err := GetCachedBSMResponse(app, u)
	if err != nil {
		t.Fatalf("GetCachedBSMResponse returned error: %v", err)
	}
	if got != "cached-body" {
		t.Fatalf("unexpected cached body: got %q want %q", got, "cached-body")
	}
}

func TestGetCachedBSMResponse_InvalidHost_ReturnsAllowlistError(t *testing.T) {
	t.Setenv("BSM_API_HOST", "bsm.baseball-softball.de")
	raw := "https://example.com/clubs/485/licenses.json"
	u, _ := url.Parse(raw)

	app := &MockApp{
		findFirstRecordByDataFunc:    func(any, string, any) (*core.Record, error) { return nil, nil },
		findCollectionByNameOrIdFunc: func(string) (*core.Collection, error) { return core.NewBaseCollection(dp.RequestCacheCollection), nil },
		saveFunc:                     func(core.Model) error { return nil },
	}

	_, err := GetCachedBSMResponse(app, u)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var URLAllowlistError *bsm.URLAllowlistError
	if !errors.As(err, &URLAllowlistError) {
		t.Fatalf("expected URLAllowlistError, got %T: %v", err, err)
	}
}

func TestGetCachedBSMResponse_InvalidPath_ReturnsAllowlistError(t *testing.T) {
	t.Setenv("BSM_API_HOST", "bsm.baseball-softball.de")
	raw := "https://bsm.baseball-softball.de/people/123.json" // not allowlisted
	u, _ := url.Parse(raw)

	app := &MockApp{
		findFirstRecordByDataFunc:    func(any, string, any) (*core.Record, error) { return nil, nil },
		findCollectionByNameOrIdFunc: func(string) (*core.Collection, error) { return core.NewBaseCollection(dp.RequestCacheCollection), nil },
		saveFunc:                     func(core.Model) error { return nil },
	}

	_, err := GetCachedBSMResponse(app, u)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var URLAllowlistError *bsm.URLAllowlistError
	if !errors.As(err, &URLAllowlistError) {
		t.Fatalf("expected URLAllowlistError, got %T: %v", err, err)
	}
}

func TestIsOutdated(t *testing.T) {
	fresh := mustParseDateTime(t, time.Now().Add(-30*time.Minute))
	if isOutdated(fresh) {
		t.Errorf("fresh cache reported as outdated")
	}

	old := mustParseDateTime(t, time.Now().Add(-(time.Duration(cacheLifetimeMinutes)+1)*time.Minute))
	if !isOutdated(old) {
		t.Errorf("old cache not reported as outdated")
	}
}
