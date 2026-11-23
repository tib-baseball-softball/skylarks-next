package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"ganigeorgiev.com/pocketbase/pocketbase/core"
	"git.berlinskylarks.de/tib-baseball/bsm-go"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/base"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp/push"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/internal/tib"
	_ "git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/migrations"
	"github.com/subosito/gotenv"
)

// / Loads environment.
//   - Locally: from .env file
//   - Production: from container env
func init() {
	// TODO: expand validation, maybe set app settings
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
	pushService := push.NewPushService()

	app := dp.NewDiamondPlanner(client, pushService)

	dp.SetupSentry(app)
	bindTiBHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// bindTiBHooks registers TiB-specific handlers that are a superset of core Diamond Planner functionality.
func bindTiBHooks(app core.App) {

	//------------------- Hooks -------------------------//

	app.OnRecordAfterCreateSuccess(base.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	app.OnRecordAfterUpdateSuccess(base.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/webhooks/playerChange", tib.HandlePlayerChangedMessage(app))
		return se.Next()
	})
}
