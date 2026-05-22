package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"git.berlinskylarks.de/tib-baseball/bsm-go"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/dp"
	"git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/internal/tib"
	_ "git.berlinskylarks.de/tib-baseball/skylarks-diamond-planner/migrations"
	"github.com/pocketbase/pocketbase/core"
	"github.com/subosito/gotenv"
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

// BindTiBHooks registers TiB-specific handlers that are a superset of core Diamond Planner functionality.
func BindTiBHooks(app core.App) {

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

func main() {
	client := bsm.NewAPIClient()
	pushService := dp.NewPushService()

	app := dp.NewDiamondPlanner(client, pushService)

	BindTiBHooks(app)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
