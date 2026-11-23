package dp

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"ganigeorgiev.com/pocketbase/pocketbase"
	"ganigeorgiev.com/pocketbase/pocketbase/apis"
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"ganigeorgiev.com/pocketbase/pocketbase/plugins/migratecmd"
	"ganigeorgiev.com/pocketbase/pocketbase/tools/hook"
	"ganigeorgiev.com/pocketbase/pocketbase/tools/osutils"
	"git.berlinskylarks.de/tib-baseball/bsm-go"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/push"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/tools/store"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/spf13/cobra"
)

const (
	ContextDevelopment = "Development"
	ContextStaging     = "Staging"
	ContextProduction  = "Production"
)

// DiamondPlanner defines custom extensions to the Pocketbase app.
type DiamondPlanner interface {
	core.App
}

// BaseDiamondPlanner represents the main DiamondPlanner implementation.
type BaseDiamondPlanner struct {
	*pocketbase.PocketBase
}

func NewDiamondPlanner(client bsm.APIClient, pushService push.PushService) *BaseDiamondPlanner {
	dp := &BaseDiamondPlanner{
		PocketBase: pocketbase.New(),
	}

	if os.Getenv("APPLICATION_SECRET") == "" {
		log.Fatal("APPLICATION_SECRET not set, error loading environment variables")
	}

	if osutils.IsProbablyGoRun() {
		slog.Info("App was started with `go run`, automigrations are active.")
	}

	migratecmd.MustRegister(dp, dp.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: osutils.IsProbablyGoRun(),
	})

	BindDPHooks(dp, client, pushService)

	//------------------- Commands -------------------------//

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "import:games",
		Run: func(cmd *cobra.Command, args []string) {
			// app is a pointer here because commands only work after app.Start()
			gamesImportService := GameImportService{App: *dp}
			gamesImportService.ImportGames()
		},
	})

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "import:leagues",
		Run: func(cmd *cobra.Command, args []string) {
			err := ImportLeagueGroups(dp, client, nil, nil)
			if err != nil {
				log.Print("Error while running LeagueGroupImport: " + err.Error())
			}
		},
	})

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "push:generateVAPID",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
			if err != nil {
				log.Print("Error while generating VAPID keys: " + err.Error())
			}
			fmt.Printf("Private key: %s\nPublic key: %s\n", privateKey, publicKey)
		},
	})

	return dp
}

func BindDPHooks(app core.App, client bsm.APIClient, pushService push.PushService) {

	//------------------- Hooks -------------------------//

	app.OnRecordAuthRequest(base.UserCollection).BindFunc(func(e *core.RecordAuthRequestEvent) error {
		err := SetLastLogin(e.App, e.Record)
		if err != nil {
			app.Logger().Error(
				"Error upon setting auth record last login field",
				"error", err, "user", e.Record.Id,
			)
		}
		return e.Next()
	})

	app.OnRecordAuthWithOAuth2Request(base.UserCollection).BindFunc(OAuthUpdateUserData)

	app.OnRecordCreateRequest(base.UserCollection).BindFunc(ValidateSignupKey)

	app.OnRecordAfterCreateSuccess(base.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return NotifyAdminsUserCreation(e, pushService)
	})

	app.OnRecordEnrich(base.UserCollection).BindFunc(AddUserICalLink)

	app.OnRecordUpdateExecute(base.UserCollection).BindFunc(SetDisplayName)

	app.OnRecordsListRequest(base.LeagueGroupsCollection).BindFunc(func(event *core.RecordsListRequestEvent) error {
		return TriggerLeagueImport(event.App, client, event)
	})

	app.OnRecordCreateRequest(base.EventsCollection).BindFunc(ValidateEventTimes)

	app.OnRecordUpdateRequest(base.EventsCollection).BindFunc(ValidateEventTimes)

	app.OnRecordValidate(base.EventsCollection).BindFunc(ValidateEvent)

	app.OnRecordEnrich(base.EventsCollection).BindFunc(func(event *core.RecordEnrichEvent) error {
		return AddEventParticipationData(event.App, event)
	})

	app.OnRecordEnrich(base.EventSeriesCollection).BindFunc(AddSeriesState)

	app.OnRecordEnrich(base.ServiceEntryCollection).BindFunc(HideCommunityServiceAdminFields)

	app.OnRecordCreateRequest(base.EventSeriesCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return CreateOrUpdateEventsForSeries(e, CreateEventSeriesMode)
	})

	app.OnRecordUpdateRequest(base.EventSeriesCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return CreateOrUpdateEventsForSeries(e, UpdateEventSeriesMode)
	})

	app.OnRecordDelete(base.EventSeriesCollection).BindFunc(DeleteEventsForSeries)

	app.OnRecordEnrich(base.TeamsCollection).BindFunc(HideTeamSignupKey)

	//------------------- Hooks - Web Push -------------------------//

	app.OnRecordAfterCreateSuccess(base.AnnouncementsCollection).BindFunc(func(e *core.RecordEvent) error {
		return NotifyNewAnnouncement(e, pushService)
	})

	app.OnRecordAfterCreateSuccess(base.EventsCollection).BindFunc(func(e *core.RecordEvent) error {
		return NotifyNewEvent(e, pushService)
	})

	//------------------- Serve static dir -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if it exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), true))

		return se.Next()
	})

	//------------------- Middlewares -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.Bind(&hook.Handler[*core.RequestEvent]{
			Id:       "block_noise",
			Func:     BlockKnownPaths,
			Priority: 100,
		})

		return se.Next()
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/hello", func(e *core.RequestEvent) error {
			return e.String(http.StatusOK, "Hello ballplayers!")
		})
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/import/{club}/leagues", StartLeagueGroupsImport(app, client))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/stats/{user}",
			GetUserStats()).
			Bind(apis.RequireAuth())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/communityservice/{user}/{season}",
			GetUserCommunityService()).
			Bind(apis.RequireAuth())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		icalGroup := se.Router.Group("/api/dp/ical")

		icalGroup.GET("/{user}/{hash}/calendar.ics", getUserCalendar(se.App))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		teamsGroup := se.Router.Group("/api/dp/teams")

		teamsGroup.Bind(apis.RequireAuth())
		individualTeamGroup := teamsGroup.Group("/{team}")

		individualTeamGroup.GET("/gamecount", GetGamesCount(app))
		individualTeamGroup.POST("/join", JoinTeam(app))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		clubAdminGroup := se.Router.Group("/api/clubs/{club}/admin")

		clubAdminGroup.Bind(RequireClubAdminAccess())
		clubAdminGroup.GET("/communityservice", GetClubCommunityService())
		clubAdminGroup.GET("/practices", HandleClubPractices())
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		pushGroup := se.Router.Group("/api/webpush")
		pushGroup.Bind(apis.RequireAuth())

		pushGroup.POST("/{user}/notify-me", push.NotifyWithTestPushMessage(pushService, app))

		pushGroup.DELETE("/unsubscribe", push.DeleteSubscription(app))

		return se.Next()
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {
		app.Cron().MustAdd("LeagueGroupImport", "0 * * * *", func() {
			err := ImportLeagueGroups(app, client, nil, nil)
			if err != nil {
				app.Logger().Error("Error while running cronjob LeagueGroupImport: " + err.Error())
			}
		})

		app.Cron().MustAdd("GamesImport", "0 * * * *", func() {
			gamesImportService := GameImportService{App: app}
			gamesImportService.ImportGames()
		})
	}
}

// SetupSentry performs required scaffolding for the Sentry SDK.
//
// If the key `SENTRY_DSN` is not found in the environment, setup is skipped.
//
// The resulting sentry state is written to the app store with key `FlagSentryEnabled`
func SetupSentry(app DiamondPlanner) {
	applicationContext := os.Getenv("APPLICATION_CONTEXT")
	isDevelopment := applicationContext == ContextDevelopment

	sentryDSN, dsnIsSet := os.LookupEnv("SENTRY_DSN")

	if !dsnIsSet {
		app.Logger().Info("Sentry DSN is not set, skipping Sentry init", "DSN", sentryDSN)
		return
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:                  sentryDSN,
		Debug:                isDevelopment,
		SendDefaultPII:       false,
		DisableClientReports: true,
		Environment:          applicationContext,
		TracesSampleRate:     0.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentryMiddleware := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.BindFunc(apis.WrapStdMiddleware(sentryMiddleware.Handle))

		return se.Next()
	})
	app.Logger().Info("Sentry Middleware set up", "DSN", sentryDSN)
	app.Store().Set(string(store.FlagSentryEnabled), true)
}
