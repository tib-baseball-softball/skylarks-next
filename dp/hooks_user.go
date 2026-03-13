package dp

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
	"github.com/spf13/cast"
)

// NotifyAdminsUserCreation is a hook that handles push notifications to club admins when a new user is created
// as well as email notifications to superusers.
func NotifyAdminsUserCreation(e *core.RecordEvent, ps PushService) error {
	err := notifyAdminsUserCreation(e.Record, e.App, ps)
	if err != nil {
		e.App.Logger().Warn(err.Error())
	}
	sendSuperUserNotificationForNewUsers(e.Record, e.App)

	return e.Next()
}

// sendSuperUserNotificationForNewUsers sends an email notification to superusers when a new user is created.
func sendSuperUserNotificationForNewUsers(record *core.Record, app core.App) {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		user := &User{}
		user.SetProxyRecord(record)

		superusers, err := app.FindAllRecords(core.CollectionNameSuperusers)
		if err != nil {
			app.Logger().Warn("Couldn't fetch superuser records for email notification", "error", err)
			return
		}

		client := app.NewMailClient()
		if client == nil {
			app.Logger().Warn("Failed to create mail client")
			return
		}

		html, err := template.NewRegistry().LoadFiles(
			"templates/email/newUserCreated.gohtml",
		).Render(map[string]any{
			"username":  user.Username(),
			"firstName": user.FirstName(),
			"lastName":  user.LastName(),
			"email":     user.Email(),
			"signupKey": user.SignupKey(),
			"teams":     user.Teams(),
			"club":      user.Clubs(),
			"url":       app.Settings().Meta.AppURL,
		})
		if err != nil {
			app.Logger().Warn("error rendering superuser notification email template", "error", err)
			return
		}

		for _, superuser := range superusers {
			select {
			case <-ctx.Done():
				app.Logger().Warn("Context cancelled during email sending")
				return
			default:
			}

			message := &mailer.Message{
				From: mail.Address{
					Address: app.Settings().Meta.SenderAddress,
					Name:    app.Settings().Meta.SenderName,
				},
				To:      []mail.Address{{Address: superuser.Email()}},
				Subject: "New user created in Diamond Planner",
				HTML:    html,
				Text:    fmt.Sprintf("New user created in Diamond Planner: %s %s (%s). Email: %s, Signup Key: %s", user.FirstName(), user.LastName(), user.Username(), user.Email(), user.SignupKey()),
			}

			err = client.Send(message)
			if err != nil {
				app.Logger().Warn("Couldn't send email notification to superuser", "error", err, "email", superuser.Email())
			}
		}
	}()
}

// notifyAdminsUserCreation sends a push notification to club admins when a new user is created.
func notifyAdminsUserCreation(record *core.Record, app core.App, ps PushService) error {
	user := &User{}
	user.SetProxyRecord(record)

	ids := make([]interface{}, len(user.Clubs()))
	for i, id := range user.Clubs() {
		ids[i] = id
	}

	var clubs []Club
	err := app.RecordQuery(ClubsCollection).
		AndWhere(dbx.In("id", ids...)).
		All(&clubs)
	if err != nil {
		app.Logger().Warn("Could not query clubs for user", "userID", user.ID(), "error", err)
		return err
	}

	adminIDs := make([]string, 0)
	for _, club := range clubs {
		admins, err := GetAdminsForClub(club, app)
		if err != nil {
			app.Logger().Warn("Could not get admins for club", "club", club.Id, "error", err)
			continue
		}
		for _, admin := range admins {
			adminIDs = append(adminIDs, admin.Id)
		}
	}

	subs, err := GetSubscriptionsForUserIDs(adminIDs, app)
	if err != nil {
		app.Logger().Warn("Could not get subscriptions for users", "userIDs", adminIDs, "error", err)
		return err
	}

	msg := &PushMessage{
		Title: "New User Created in Diamond Planner",
		Body:  fmt.Sprintf("A new user has been created for your club: %s %s.", user.FirstName(), user.LastName()),
		Tag:   "user_created",
	}

	for _, sub := range subs {
		app.Logger().Debug("Sending notification to admin", "admin", sub.User())

		err := ps.SendPushMessage(msg, new(sub.ToWebPushSubscription()))
		if err != nil {
			return err
		}
	}
	return nil
}

// OAuthUpdateUserData updates user record with information found in the OAuth2 response.
// This could also be used to validate the signup key, but does not do that at the moment
// (see generic record creation hook that also works for password auth)
func OAuthUpdateUserData(e *core.RecordAuthWithOAuth2RequestEvent) error {
	if e.Record == nil {
		return e.Next()
	}
	user := &User{}
	user.SetProxyRecord(e.Record)

	rawUser := e.OAuth2User.RawUser
	firstName, firstNameExists := rawUser["given_name"]
	lastName, lastNameExists := rawUser["family_name"]

	if firstNameExists {
		user.SetFirstName(cast.ToString(firstName))
	}
	if lastNameExists {
		user.SetLastName(cast.ToString(lastName))
	}

	err := e.App.Save(user)
	if err != nil {
		e.App.Logger().Warn("Could not persist new OAuth2 user data", "rawUser", rawUser, "record", user.ID())
		return e.Next()
	}
	return e.Next()
}

// SetDisplayName sets the display name for the user if it is not already set.
// Default display name format: first name + first letter of last name + dot
func SetDisplayName(e *core.RecordEvent) error {
	user := &User{}
	user.SetProxyRecord(e.Record)

	if user.DisplayName() == "" {
		user.SetDisplayName(user.FirstName() + " " + fmt.Sprintf("%.*s", 1, user.LastName()) + ".")
	}

	return e.Next()
}

// AddUserICalLink adds the user's custom iCal link to the JSON response.
func AddUserICalLink(e *core.RecordEnrichEvent) error {
	user := &User{}
	user.SetProxyRecord(e.Record)

	auth := e.RequestInfo.Auth

	if auth.Id == user.Id {
	    user.WithCustomData(true)
					
	    appSecret := os.Getenv("APPLICATION_SECRET")
		hash := GetSHA3Hash(user.Id + appSecret)
		user.SetICalLink(e.App.Settings().Meta.AppURL + "/api/dp/ical/" + user.Id + "/" + hash)
	}

	return e.Next()
}
