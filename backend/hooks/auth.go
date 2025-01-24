package hooks

import (
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
	"time"
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
	clubs, err := getValidSignupKeys(app)
	if err != nil {
		errorText := "failed to get valid signup clubs"

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
	for _, club := range clubs {
		if body.SignupKey != "" && club.SignupKey == body.SignupKey {
			isValid = true
			event.Record.Set("club", club.Id)
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

// Club Reduced record with just the relevant fields
type Club struct {
	Id        string `db:"id"`
	SignupKey string `db:"signup_key"`
}

func getValidSignupKeys(app core.App) ([]Club, error) {
	var clubs []Club

	err := app.DB().
		NewQuery("SELECT id, signup_key FROM clubs WHERE signup_key != '';").
		All(&clubs)

	if err != nil {
		return nil, err
	}
	return clubs, nil
}
