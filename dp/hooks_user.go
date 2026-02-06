package dp

import (
	"fmt"
	"net/mail"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
	"github.com/spf13/cast"
)

func NotifyAdminsUserCreation(e *core.RecordEvent, ps PushService) error {
	err := notifyAdminsUserCreation(e.Record, e.App, ps)
	if err != nil {
		e.App.Logger().Warn(err.Error())
	}
	err = sendSuperUserNotification(e.Record, e.App)
	if err != nil {
		e.App.Logger().Warn(err.Error())
	}

	return e.Next()
}

func sendSuperUserNotification(record *core.Record, app core.App) error {
	user := &User{}
	user.SetProxyRecord(record)

	superusers, err := app.FindAllRecords(core.CollectionNameSuperusers)
	if err != nil {
		app.Logger().Warn("Couldn't fetch superuser records for email notification", "error", err)
		return err
	}

	client := app.NewMailClient()

	html, err := template.NewRegistry().LoadFiles(
		"templates/email/newUserCreated.gohtml",
	).Render(map[string]any{
		"username":  user.Username(),
		"firstName": user.FirstName(),
		"lastName":  user.LastName(),
		"email":     user.Email(),
		"signupKey": user.SignupKey(),
		"teams":     user.Teams(),
		"club":      user.Club(),
		"url":       app.Settings().Meta.AppURL,
	})
	if err != nil {
		app.Logger().Warn("error rendering superuser notification email template", "error", err)
		return err
	}

	for _, superuser := range superusers {
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
			return err
		}
	}
	return nil
}

// notifyAdminsUserCreation sends a push notification to club admins when a new user is created
func notifyAdminsUserCreation(record *core.Record, app core.App, ps PushService) error {
	user := &User{}
	user.SetProxyRecord(record)

	ids := make([]interface{}, len(user.Club()))
	for i, id := range user.Club() {
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

		ws := sub.ToWebPushSubscription()
		err := ps.SendPushMessage(msg, &ws)
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
