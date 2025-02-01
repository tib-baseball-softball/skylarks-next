package bsm

import (
	"net/url"
	"testing"
)

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
