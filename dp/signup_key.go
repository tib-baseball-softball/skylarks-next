package dp

import (
	validation "github.com/pocketbase/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
)

// ValidateSignupKey checks that a valid signup key is used on user creation.
// A user is assigned a team and club based on the signup key used.
//
// Superuser auth bypasses this check (so creating users from PocketBase admin panel is always possible).
func ValidateSignupKey(e *core.RecordRequestEvent) error {
	if e.HasSuperuserAuth() {
		return e.Next()
	}

	user := &User{}
	user.SetProxyRecord(e.Record)

	teams, err := getValidSignupKeys(e.App)
	if err != nil {
		errorText := "failed to get valid signup teams"

		e.App.Logger().Error(errorText, "error", err)
		return e.InternalServerError(errorText, err)
	}

	// read signup key from request body as it's not present in the user record
	body := struct {
		SignupKey string `json:"signup_key"`
	}{}
	if err := e.BindBody(&body); err != nil {
		msg := "Malformed request body"
		return e.BadRequestError("Failed to read request body", map[string]validation.Error{
			"body": validation.NewError("malformed_request_body", msg),
		})
	}

	isValid := false
	for _, team := range teams {
		if body.SignupKey != "" && team.SignupKey == body.SignupKey {
			isValid = true
			user.SetTeams([]string{team.Id})
			user.SetClubs([]string{team.Club})
			user.SetSignupKey(body.SignupKey)
			break
		}
	}

	if !isValid {
		msg := "The entered signup key is not valid for any team."
		return e.BadRequestError(msg, map[string]validation.Error{
			"signup_key": validation.NewError("signup_key_invalid", msg),
		})
	}

	return e.Next()
}

// TeamWithSignupKey represents a reduced record with just the relevant fields.
type TeamWithSignupKey struct {
	Id        string `db:"id"`
	Club      string `db:"club"`
	SignupKey string `db:"signup_key"`
}

// getValidSignupKeys collects all possible signup keys for any team in the database.
//
// The query is static, so if this returns an error, something is really wrong.
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
