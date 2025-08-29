package tib

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/internal/pb"
)

type PlayerChangedMessage struct {
	PlayerUid int    `json:"playerUid"`
	BSMID     int    `json:"bsmID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func HandlePlayerChangedMessage(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		webhookIdentifier := os.Getenv("PLAYER_CHANGE_WEBHOOK_IDENTIFIER")
		webhookSecret := os.Getenv("PLAYER_CHANGE_WEBHOOK_SECRET")

		if webhookIdentifier == "" || webhookSecret == "" {
			app.Logger().Error("Tried to receive player change webhook, but identifier or secret were not set.")
			return e.JSON(http.StatusInternalServerError, "Internal Server Error")
		}

		body, err := io.ReadAll(e.Request.Body)
		if err != nil {
			app.Logger().Error("Failed to read request body: %v", err)
			return e.JSON(http.StatusInternalServerError, "Internal Server Error")
		}

		if !verifyWebhookSignature(e.Request, webhookIdentifier, webhookSecret, body) {
			app.Logger().Warn("Invalid webhook signature for player change", "identifier", webhookIdentifier)
			return e.JSON(http.StatusUnauthorized, "Invalid Webhook signature")
		}

		var message PlayerChangedMessage
		err = json.Unmarshal(body, &message)
		if err != nil {
			return e.JSON(http.StatusBadRequest, "Player Change Message does not match expected format")
		}

		record, err := app.FindFirstRecordByFilter(
			pb.UserCollection,
			"first_name = {:firstName} && last_name = {:lastName}",
			dbx.Params{"firstName": message.FirstName, "lastName": message.LastName},
		)
		if err != nil {
			app.Logger().Warn("Failed to find player for webhook message", "err", err, "message", message)
			return e.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
		if record == nil {
			return e.JSON(http.StatusNotFound, "Player not found.")
		}

		player := &pb.User{}
		player.SetProxyRecord(record)

		if player.BSMID() != 0 && player.BSMID() != message.BSMID {
			return e.JSON(http.StatusNoContent, "Player already has a BSM ID, changing it is not supported.")
		} else {
			player.SetBSMID(message.BSMID)
			err := app.Save(player.Record)
			if err != nil {
				app.Logger().Error("Failed to save player", "err", err, "player", player)
				return e.JSON(http.StatusInternalServerError, "Internal Server Error")
			}
			app.Logger().Info("Updated player BSM ID", "message", message, "player", player)
			return e.JSON(http.StatusOK, "Player BSM ID updated.")
		}
	}
}
