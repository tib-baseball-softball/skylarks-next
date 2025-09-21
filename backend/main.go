package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	_ "time/tzdata"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/spf13/cobra"
	"github.com/subosito/gotenv"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
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
}

func bindAppHooks(app core.App) {

	//------------------- Hooks -------------------------//

	app.OnRecordAuthRequest(dp.UserCollection).BindFunc(func(e *core.RecordAuthRequestEvent) error {
		err := dp.SetLastLogin(e.App, e.Record)
		if err != nil {
			app.Logger().Error(
				"Error upon setting auth record last login field",
				"error", err, "user", e.Record.Id,
			)
		}
		return e.Next()
	})

	app.OnRecordCreateRequest(dp.UserCollection).BindFunc(func(event *core.RecordRequestEvent) error {
		return dp.ValidateSignupKey(event)
	})

	app.OnRecordsListRequest("leaguegroups").BindFunc(func(event *core.RecordsListRequestEvent) error {
		return dp.TriggerLeagueImport(event.App, event)
	})

	app.OnRecordCreateRequest(dp.EventsCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return dp.ValidateEventTimes(e)
	})

	app.OnRecordUpdateRequest(dp.EventsCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return dp.ValidateEventTimes(e)
	})

	app.OnRecordEnrich(dp.EventsCollection).BindFunc(func(event *core.RecordEnrichEvent) error {
		return dp.AddEventParticipationData(event.App, event)
	})

	app.OnRecordAfterCreateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e)
	})

	app.OnRecordAfterUpdateSuccess(dp.UserCollection).BindFunc(func(e *core.RecordEvent) error {
		return tib.SendUpdatedPlayerData(e)
	})

	app.OnRecordAfterCreateSuccess(dp.EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return dp.CreateOrUpdateEventsForSeries(e)
	})

	app.OnRecordAfterUpdateSuccess(dp.EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return dp.CreateOrUpdateEventsForSeries(e)
	})

	app.OnRecordDelete(dp.EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return dp.DeleteEventsForSeries(e)
	})

	//------------------- Serve static dir -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if it exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

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
		se.Router.GET("/api/gamecount/{team}", dp.GetGamesCount(app))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/import/{club}/leagues", dp.StartLeagueGroupsImport(app))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/stats/{user}", dp.GetUserStats())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/bsm/relay/top10/{league}", tib.GetLeagueLeaders())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/bsm/relay", tib.GetRelayedBSMData())
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/team/favorite", tib.GetFavoriteTeamData())
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/webhooks/playerChange", tib.HandlePlayerChangedMessage(app))
		return se.Next()
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {
		app.Cron().MustAdd("LeagueGroupImport", "0 * * * *", func() {
			err := dp.ImportLeagueGroups(app, nil, nil)
			if err != nil {
				app.Logger().Error("Error while running cronjob LeagueGroupImport: " + err.Error())
			}
		})

		app.Cron().MustAdd("GamesImport", "0 * * * *", func() {
			gamesImportService := dp.GameImportService{App: app}
			gamesImportService.ImportGames()
		})

		app.Cron().MustAdd("TeamDatasetImport", "30 * * * *", func() {
			tib.ImportTeamDatasets(app)
		})
	}
}

func main() {
	app := pocketbase.New()

	bindAppHooks(app)

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	if isGoRun {
		log.Print("App was started with `go run`, automigrations are active.")
	}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	//------------------- Commands -------------------------//

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:games",
		Run: func(cmd *cobra.Command, args []string) {
			// app is a pointer here because commands only work after app.Start()
			gamesImportService := dp.GameImportService{App: *app}
			gamesImportService.ImportGames()
		},
	})

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:leagues",
		Run: func(cmd *cobra.Command, args []string) {
			err := dp.ImportLeagueGroups(app, nil, nil)
			if err != nil {
				log.Print("Error while running LeagueGroupImport: " + err.Error())
			}
		},
	})

	app.RootCmd.AddCommand(&cobra.Command{
		Use: "import:teamdatasets",
		Run: func(cmd *cobra.Command, args []string) {
			tib.ImportTeamDatasets(app)
		},
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
