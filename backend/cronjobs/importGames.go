package cronjobs

import (
    "encoding/json"
    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/tools/types"
    "github.com/tib-baseball-softball/skylarks-next/bsm"
    "github.com/tib-baseball-softball/skylarks-next/model"
    "io"
    "log"
    "net/http"
    "sync"
    "time"
)

func ImportGames(app *pocketbase.PocketBase) {
    teams, err := app.Dao().FindRecordsByFilter("teams", "bsm_league_group != 0", "", 0, 0)
    if err != nil {
        log.Print("Error fetching existing team records: ", err)
        return
    }

    var wg sync.WaitGroup

    for _, team := range teams {
        wg.Add(1)
        go func() {
            defer wg.Done()
            matches := fetchMatchesForLeagueGroup(team.GetString("bsm_league_group"))
            err := createOrUpdateEvents(app, matches, team.Id)
            if err != nil {
                log.Print(err)
                return
            }
        }()
    }

    wg.Wait()
    log.Println("Game Import sucessfully imported new game events")
}

func fetchMatchesForLeagueGroup(league string) []model.Match {

    params := make(map[string]string)
    params["filters[leagues][]"] = league
    params["search"] = "skylarks"

    url := bsm.GetAPIURL("matches.json", params)

    resp, err := http.Get(url.String())
    if err != nil {
        log.Print(err)
        return nil
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Print(err)
        return nil
    }

    var apiResponse []model.Match
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        log.Print(err)
        return nil
    }

    return apiResponse
}

func createOrUpdateEvents(app *pocketbase.PocketBase, matches []model.Match, teamID string) (err error) {
    for _, match := range matches {
        record, err := app.Dao().FindFirstRecordByData("events", "bsm_id", match.ID)

        // if not found, it throws an error, so create new record
        if err != nil {
            collection, err := app.Dao().FindCollectionByNameOrId("events")
            if err != nil {
                return err
            }

            record = models.NewRecord(collection)
            err = setRecordValues(record, match, teamID)
            if err != nil {
                return err
            }
        }
        // no error - update existing record
        err = setRecordValues(record, match, teamID)
        if err != nil {
            return err
        }

        if err := app.Dao().SaveRecord(record); err != nil {
            log.Print("Persisting event record failed: ", err)
            return err
        }
    }
    return
}

func setRecordValues(record *models.Record, match model.Match, teamID string) (err error) {
    starttime, err := types.ParseDateTime(match.Time)
    if err != nil {
        return err
    }
    endtime := starttime.Time().Add(time.Hour * 3)

    record.Set("title", match.AwayTeamName + " @ " + match.HomeTeamName)
    record.Set("bsm_id", match.ID)
    record.Set("starttime", starttime.String())
    record.Set("endtime", endtime.String())
    record.Set("type", "game")
    record.Set("team", teamID)

    return
}