package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/cron"
	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
	"github.com/tib-baseball-softball/skylarks-next/cronjobs"
	"github.com/tib-baseball-softball/skylarks-next/hooks"
)

func init() {
	err := gotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables.")
	}
}

func main() {
	app := pocketbase.New()

	//------------------- Hooks -------------------------//

	app.OnRecordAuthRequest("users").BindFunc(func(e *core.RecordAuthRequestEvent) error {
		err := hooks.SetLastLogin(app, e.Record)
		if err != nil {
			app.Logger().Error(
				"Error upon setting auth record last login field",
				"error", err,
			)
		}
		return nil
	})

	app.OnRecordCreateRequest("users").BindFunc(func(e *core.RecordRequestEvent) error {
		// TODO: validate signup Key
		return nil
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/hello", func(e *core.RequestEvent) error {
			return e.String(http.StatusOK, "Hello ballplayers!")
		})

		return nil
	})

	//------------------- Commands -------------------------//

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:games",
		Run: func(cmd *cobra.Command, args []string) {
			cronjobs.ImportGames(app)
		},
	})

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:leagues",
		Run: func(cmd *cobra.Command, args []string) {
			err := cronjobs.ImportLeagueGroups(app)
			if err != nil {
				log.Print("Error while running LeagueGroupImport: " + err.Error())
			}
		},
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {
		app.OnServe().BindFunc(func(e *core.ServeEvent) error {
			scheduler := cron.New()

			scheduler.MustAdd("LeagueGroupImport", "0 * * * *", func() {
				err := cronjobs.ImportLeagueGroups(app)
				if err != nil {
					log.Print("Error while running cronjob LeagueGroupImport: " + err.Error())
				}
			})

			scheduler.MustAdd("GamesImport", "0 * * * *", func() {
				cronjobs.ImportGames(app)
			})

			scheduler.Start()

			return nil
		})
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
