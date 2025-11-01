package tib

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/tib-baseball-softball/skylarks-next/bsm"
	"github.com/tib-baseball-softball/skylarks-next/internal/dp"
)

type teamDatasetCacheIdentifier struct {
	TeamID int
	Season string
}

func ImportTeamDatasets(app core.App, client bsm.APIClient) {
	currentYear := time.Now().Year()
	cacheRecords, err := app.FindRecordsByFilter(
		dp.RequestCacheCollection,
		"identifier ~ {:currentYear} && identifier ~ {:cacheTag}",
		"",
		0,
		0,
		dbx.Params{
			"currentYear": currentYear,
			"cacheTag":    homeDatasetCacheTag,
		},
	)
	if err != nil {
		app.Logger().Error("ImportTeamDatasets error", "error", err)
		return
	}

	processedRecords := 0
	var wg sync.WaitGroup

	for _, cacheRecord := range cacheRecords {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestCache := &dp.RequestCache{}
			requestCache.SetProxyRecord(cacheRecord)

			if isOutdated(requestCache.Updated()) {
				parsedIdentifier, err := parseTeamDatasetCacheIdentifier(requestCache.Identifier())
				if err != nil {
					app.Logger().Error("ImportTeamDatasets error", "error", err, "cacheIdentifier", requestCache.Identifier())
					return
				}

				datasets, err := LoadHomeData(app, client, parsedIdentifier.TeamID, currentYear) // we have established earlier that only the current year is of interest
				responseBody, err := json.Marshal(datasets)
				if err != nil {
					app.Logger().Error("ImportTeamDatasets error", "error", err, "cacheIdentifier", requestCache.Identifier())
					return
				}

				requestCache.SetResponseBody(string(responseBody))
				if err = app.Save(requestCache); err != nil {
					app.Logger().Error("Failed saving TeamDatasets cache record", "error", err, "cacheIdentifier", requestCache.Identifier(), "record", requestCache)
					return
				}
				processedRecords++
			}
		}()
	}
	wg.Wait()
	app.Logger().Info("TeamDatasets Import successfully imported new team datasets.", "Number of records processed", processedRecords)
}

const teamDatasetCacheIdentifierLength = 3

func parseTeamDatasetCacheIdentifier(identifier string) (teamDatasetCacheIdentifier, error) {
	var ret teamDatasetCacheIdentifier
	slice := strings.Split(identifier, "_")
	if len(slice) != teamDatasetCacheIdentifierLength {
		return ret, errors.New("invalid cache identifier")
	}

	if slice[0] == "" || slice[1] == "" {
		return ret, errors.New("invalid cache identifier")
	}

	convertedTeamID, err := strconv.Atoi(slice[1])
	if err != nil {
		return ret, errors.New("invalid cache identifier, could not convert team ID to integer")
	}

	ret.Season = slice[0]
	ret.TeamID = convertedTeamID

	return ret, nil
}
