package hooks

import (
	"errors"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

func SetLastLogin(app core.App, record *core.Record) (err error) {
	currentTime := time.Now()
	record.Set("last_login", currentTime)

	if err := app.Save(record); err != nil {
		app.Logger().Error("Persisting user last login time failed", "error", err)

		return err
	}

	return nil
}

func ValidateSignupKey(app core.App, event *core.RecordRequestEvent) error {
	teams, err := getValidSignupKeys(app)
	if err != nil {
		errorText := "failed to get valid signup teams"

		app.Logger().Error(errorText, "error", err)
		return errors.New(errorText)
	}

	// read signup key from request body as it's not present in the user record
	body := struct {
		SignupKey string `json:"signup_key"`
	}{}
	if err := event.BindBody(&body); err != nil {
		return event.BadRequestError("Failed to read request body", err)
	}

	isValid := false
	for _, team := range teams {
		if body.SignupKey != "" && team.SignupKey == body.SignupKey {
			isValid = true
			event.Record.Set("teams", team.Id)
			event.Record.Set("club", team.Club)
			event.Record.Set("signup_key", body.SignupKey)
			break
		}
	}

	if !isValid {
		err := event.JSON(http.StatusBadRequest, map[string]string{
			"message": "Signup key invalid",
		})
		if err != nil {
			return err
		}
		return event.BadRequestError("signup key invalid", err)
	}

	return event.Next()
}

// Team Reduced record with just the relevant fields
type Team struct {
	Id        string `db:"id"`
	Club      string `db:"club"`
	SignupKey string `db:"signup_key"`
}

func getValidSignupKeys(app core.App) ([]Team, error) {
	var teams []Team

	err := app.DB().
		NewQuery("SELECT id, club, signup_key FROM teams WHERE signup_key != '';").
		All(&teams)

	if err != nil {
		return nil, err
	}
	return teams, nil
}
