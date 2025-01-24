package main

import (
	"github.com/pocketbase/pocketbase/tests"
	"net/http"
	"testing"
)

const testDataDir = "./test_pb_data"

func generateToken(collectionNameOrId string, email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.FindAuthRecordByEmail(collectionNameOrId, email)
	if err != nil {
		return "", err
	}

	return record.NewAuthToken()
}

func TestHelloEndpoint(t *testing.T) {
	setupTestApp := func(t testing.TB) *tests.TestApp {
		testApp, err := tests.NewTestApp(testDataDir)
		if err != nil {
			t.Fatal(err)
		}
		// no need to clean up since scenario.Test() will do that for us

		bindAppHooks(testApp)

		return testApp
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "try with different http method, e.g. POST",
			Method:          http.MethodPost,
			URL:             "/api/hello",
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"wasn't found"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "any auth status with GET",
			Method:          http.MethodGet,
			URL:             "/api/hello",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"Hello ballplayers!"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
