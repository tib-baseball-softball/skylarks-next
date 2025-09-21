package dp

import (
	"errors"
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func ValidateSignupKey(app core.App, event *core.RecordRequestEvent) error {
	user := &User{}
	user.SetProxyRecord(event.Record)

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
			user.SetTeams([]string{team.Id})
			user.SetClub([]string{team.Club})
			user.SetSignupKey(body.SignupKey)
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

// TeamWithSignupKey Reduced record with just the relevant fields
type TeamWithSignupKey struct {
	Id        string `db:"id"`
	Club      string `db:"club"`
	SignupKey string `db:"signup_key"`
}

func getValidSignupKeys(app core.App) ([]TeamWithSignupKey, error) {
	var teams []TeamWithSignupKey

	err := app.DB().
		NewQuery("SELECT id, club, signup_key FROM teams WHERE signup_key != '';").
		All(&teams)

	if err != nil {
		return nil, err
	}
	return teams, nil
}
