package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
	"github.com/tib-baseball-softball/skylarks-next/cronjobs"
	"github.com/tib-baseball-softball/skylarks-next/hooks"
	"github.com/tib-baseball-softball/skylarks-next/routes"
	"log"
	"net/http"
	"os"
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
		return e.Next()
	})

	app.OnRecordCreateRequest("users").BindFunc(func(event *core.RecordRequestEvent) error {
		return hooks.ValidateSignupKey(app, event)
	})

	app.OnRecordsListRequest("leaguegroups").BindFunc(func(event *core.RecordsListRequestEvent) error {
		return hooks.TriggerLeagueImport(app, event)
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/hello", func(e *core.RequestEvent) error {
			return e.String(http.StatusOK, "Hello ballplayers!")
		})
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/gamecount/{team}", routes.GetGamesCount(app))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/import/{club}/leagues", routes.StartLeagueGroupsImport(app))

		return se.Next()
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
			err := cronjobs.ImportLeagueGroups(app, nil, nil)
			if err != nil {
				log.Print("Error while running LeagueGroupImport: " + err.Error())
			}
		},
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {
		app.Cron().MustAdd("LeagueGroupImport", "0 * * * *", func() {
			err := cronjobs.ImportLeagueGroups(app, nil, nil)
			if err != nil {
				app.Logger().Error("Error while running cronjob LeagueGroupImport: " + err.Error())
			}
		})

		app.Cron().MustAdd("GamesImport", "0 * * * *", func() {
			cronjobs.ImportGames(app)
		})
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
