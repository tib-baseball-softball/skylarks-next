package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
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

	app.OnRecordAuthRequest("users").Add(func(e *core.RecordAuthEvent) error {
		hooks.SetLastLogin(app, e.Record)
		return nil
	})

	//------------------- Custom Routes -------------------------//

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// add new "GET /hello" route to the app router (echo)
		_, err := e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/hello",
			Handler: func(c echo.Context) error {
				return c.String(200, "Hello ballplayers!")
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		if err != nil {
			log.Fatal("Could not register /hello route")
		}

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
		app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
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
