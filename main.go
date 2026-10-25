package main

import (
	"log"
	"os"
	"time"
	_ "time/tzdata"

	"git.berlinskylarks.de/tib-baseball/bsm-go"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/tools/store"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/internal/tib"
	_ "git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/migrations"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/subosito/gotenv"
)

const (
	ContextDevelopment = "Development"
	ContextStaging     = "Staging"
	ContextProduction  = "Production"
)

// / Loads environment.
//   - Locally: from .env file
//   - Production: from container env
func init() {
	_ = gotenv.Load()

	if os.Getenv("APPLICATION_CONTEXT") == "" {
		log.Fatal("APPLICATION_CONTEXT not set, error loading environment variables")
	}

	if os.Getenv("BSM_API_URL") == "" || os.Getenv("BSM_API_HOST") == "" {
		log.Fatal("BSM API URL or host not set, error loading environment variables")
	}
}

func main() {
	client := bsm.NewAPIClient()
	pushService := dp.NewPushService()

	app := dp.NewDiamondPlanner(client, pushService)

	setupSentry(app)
	bindTiBHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// setupSentry performs required scaffolding for the Sentry SDK.
//
// If the key `SENTRY_DSN` is not found in the environment, setup is skipped.
func setupSentry(app dp.DiamondPlanner) {
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
	app.Flags().Set(store.FlagSentryEnabled, true)
}

// bindTiBHooks registers TiB-specific handlers that are a superset of core Diamond Planner functionality.
func bindTiBHooks(app core.App) {

	//------------------- Hooks -------------------------//

	app.OnRecordAfterCreateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	app.OnRecordAfterUpdateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/webhooks/playerChange", tib.HandlePlayerChangedMessage(app))
		return se.Next()
	})
}
