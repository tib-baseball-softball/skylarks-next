package main

import (
    "github.com/labstack/echo/v5"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/tools/cron"
    "github.com/subosito/gotenv"
    "github.com/tib-baseball-softball/skylarks-next/cronjobs"
    "log"
    "net/http"
)

func init() {
    err := gotenv.Load()

    if err != nil {
        log.Fatal("Error loading environment variables.")
    }
}

func main() {
    app := pocketbase.New()

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

    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        scheduler := cron.New()

        scheduler.MustAdd("TeamsImport", "* * * * *", cronjobs.ImportTeams)
        scheduler.MustAdd("GamesImport", "* * * * *", func() {
            cronjobs.ImportGames(app)
        })

        scheduler.Start()

        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
