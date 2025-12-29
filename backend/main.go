package main

import (
	"log"
	"os"
	_ "time/tzdata"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/diamond-planner/diamond-planner/dp"
	_ "github.com/diamond-planner/diamond-planner/migrations"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
	"github.com/tib-baseball-softball/skylarks-next/internal/tib"
	_ "github.com/tib-baseball-softball/skylarks-next/migrations"
)

// / Loads environment.
//   - Locally: from .env file
//   - Production: from container env
func init() {
	_ = gotenv.Load()

	if os.Getenv("APPLICATION_CONTEXT") == "" {
		log.Fatal("APPLICATION_CONTEXT not set, error loading environment variables")
	}

	if os.Getenv("BSM_API_KEY") == "" || os.Getenv("BSM_API_HOST") == "" {
		log.Fatal("BSM API key or host not set, error loading environment variables")
	}
}

// BindTiBHooks registers TiB-specific handlers that are a superset of core Diamond Planner functionality.
func BindTiBHooks(app core.App, client bsm.APIClient) {

	//------------------- Hooks -------------------------//

	app.OnRecordAfterCreateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	app.OnRecordAfterUpdateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e, os.Getenv("PUBLIC_TYPO3_URL"))
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/bsm/relay/top10/{league}", tib.GetLeagueLeaders(client))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/bsm/relay/{path...}", tib.GetRelayedBSMData())
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/team/favorite", tib.GetFavoriteTeamData(client))
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/webhooks/playerChange", tib.HandlePlayerChangedMessage(app))
		return se.Next()
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {

		app.Cron().MustAdd("TeamDatasetImport", "30 * * * *", func() {
			tib.ImportTeamDatasets(app, client)
		})

		app.Cron().MustAdd("CacheCleanup", "0 4 * * *", func() {
			err := tib.CleanupOutdatedCaches(app)
			if err != nil {
				return
			}
		})
	}
}

func main() {
	client := bsm.NewAPIClient()
	app := dp.NewDiamondPlanner(client)

	BindTiBHooks(app, client)

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:teamdatasets",
		Run: func(cmd *cobra.Command, args []string) {
			tib.ImportTeamDatasets(app, client)
		},
	})

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "cache:cleanup",
		Run: func(cmd *cobra.Command, args []string) {
			err := tib.CleanupOutdatedCaches(app)
			if err != nil {
				log.Print("Error while running CacheCleanup: " + err.Error())
			}
		},
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
