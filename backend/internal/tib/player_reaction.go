package tib

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
)

type PlayerReactionPayload struct {
	PlayerId  string `json:"playerId"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	BSMID     int    `json:"bsmID"`
	Number    string `json:"number"`
	Position  []int  `json:"position"`
	Throwing  string `json:"throwing"`
	Batting   string `json:"batting"`
	Scorer    string `json:"scorer"`
	Umpire    string `json:"umpire"`
	ImageURL  string `json:"image_url"`
}

func SendUpdatedPlayerData(e *core.RecordEvent) error {
	player := &dp.User{}
	player.SetProxyRecord(e.Record)

	identifier := os.Getenv("PLAYER_UPDATE_REACTION_IDENTIFIER")
	secret := os.Getenv("PLAYER_UPDATE_REACTION_SECRET")
	typo3BaseURL := os.Getenv("PUBLIC_TYPO3_URL")

	if identifier == "" || secret == "" || typo3BaseURL == "" {
		e.App.Logger().Error("tried to send player change reaction, but identifier, secret or base URL were not set")
		return e.Next()
	}

	positions := player.GetStringSlice("position")
	var posInts []int
	for _, p := range positions {
		if n, err := strconv.Atoi(p); err == nil {
			posInts = append(posInts, n)
		}
	}

	payload := PlayerReactionPayload{
		PlayerId:  player.ID(),
		FirstName: player.FirstName(),
		LastName:  player.LastName(),
		BSMID:     player.BSMID(),
		Number:    player.Number(),
		Position:  posInts,
		Throwing:  player.Throws(),
		Batting:   player.Bats(),
		Scorer:    player.Scorer(),
		Umpire:    player.Umpire(),
		ImageURL:  player.Record.BaseFilesPath() + "/" + player.Image(),
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		e.App.Logger().Error("failed to marshal player change payload", "error", err)
		return e.Next()
	}

	url := typo3BaseURL + "/typo3/reaction/" + identifier

	// fire-and-forget the update call in a separate goroutine.
	// if it fails, we don't want to block the user update request.
	// the database transaction is already guaranteed to be committed at this point.
	go func(body []byte, app core.App) {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		defer func() {
			if r := recover(); r != nil {
				app.Logger().Error("panic in async player update", "recover", r)
			}
		}()

		request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
		if err != nil {
			e.App.Logger().Error("failed to create request for player change reaction", "error", err)
		}
		request.Header.Set("x-api-key", secret)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Accept", "application/json")
		// web server for CMS blocks at least `Go-http-client/2.0`
		request.Header.Set("User-Agent", "Skylarks Diamond Planner/1.0")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			e.App.Logger().Error("failed to send player change reaction", "error", err, "url", url, "payload", payload)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				e.App.Logger().Error("failed to close response body", "error", err)
				return
			}
		}(resp.Body)

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			e.App.Logger().Warn("player change reaction returned non-2xx status", "status", resp.Status, "url", url, "payload", payload)
		}

		e.App.Logger().Info("player change reaction sent", "status", resp.Status)
	}(bodyBytes, e.App)

	return e.Next()
}
