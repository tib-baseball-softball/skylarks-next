package main

import (
	"net/http"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
}

func TestHelloEndpoint(t *testing.T) {
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
