package tib

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/pb"
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
	player := &pb.User{}
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

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		e.App.Logger().Error("failed to create request for player change reaction", "error", err)
		return e.Next()
	}
	request.Header.Set("x-api-key", secret)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		e.App.Logger().Error("failed to send player change reaction", "error", err)
		return e.Next()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			e.App.Logger().Error("failed to close response body", "error", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		e.App.Logger().Warn("player change reaction returned non-2xx status", "status", resp.Status)
	}

	return e.Next()
}
