package hooks

import (
	"errors"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"net/http"
	"time"
)

func SetLastLogin(app *pocketbase.PocketBase, record *core.Record) (err error) {
	currentTime := time.Now()
	record.Set("last_login", currentTime)

	if err := app.Save(record); err != nil {
		app.Logger().Error("Persisting user last login time failed", "error", err)

		return err
	}

	return nil
}

func ValidateSignupKey(app *pocketbase.PocketBase, event *core.RecordRequestEvent) error {
	// TODO: not working yet
	clubs, err := getValidSignupKeys(app)
	if err != nil {
		errorText := "failed to get valid signup clubs"

		app.Logger().Error(errorText)
		return errors.New(errorText)
	}

	signupKeyUsed := event.Record.GetString("signup_key")

	isValid := false
	for _, club := range clubs {
		if signupKeyUsed != "" && club.SignupKey == signupKeyUsed {
			isValid = true
			event.Record.Set("club", club.Id)
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
		return errors.New("signup key invalid")
	}

	return event.Next()
}

// Club Reduced record with just the relevant fields
type Club struct {
	Id        string `db:"id"`
	SignupKey string `db:"signup_key"`
}

func getValidSignupKeys(app *pocketbase.PocketBase) ([]Club, error) {
	var clubs []Club

	err := app.DB().
		NewQuery("SELECT id, signup_key FROM clubs WHERE signup_key != '';").
		All(&clubs)

	if err != nil {
		return nil, err
	}
	return clubs, nil
}
