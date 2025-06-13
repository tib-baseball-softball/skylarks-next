package bsm

import (
	"encoding/json"
	"errors"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// GetAPIURL Construct a BSM API URL from a given path, appending all query parameters and the API key automatically.
// Resource should be provided without the preceding slash.
func GetAPIURL(resource string, params map[string]string, apiKey string) *url.URL {
	urlString := os.Getenv("BSM_API_URL")
	urlString = urlString + "/" + resource

	reqURL, err := url.Parse(urlString)
	if err != nil {
		log.Print(err)
	}

	query := reqURL.Query()
	query.Add("api_key", apiKey)

	for key, value := range params {
		query.Add(key, value)
	}

	reqURL.RawQuery = query.Encode()

	return reqURL
}

// AppendAPIKey is used in cases where the full URL already exists, but without the API key.
// Returns input URL with the API key for the given club appended.
func AppendAPIKey(app core.App, url url.URL, clubID string) (url.URL, error) {
	club, err := app.FindFirstRecordByData("clubs", "bsm_id", clubID)
	if err != nil {
		app.Logger().Error("Failed to get club record for ID", "club", clubID, "error", err)
		return url, err
	}

	apiKey := club.GetString("bsm_api_key")
	if apiKey == "" {
		app.Logger().Error("API key not found for club", "club", clubID)
		return url, errors.New("API key not found for club")
	}

	query := url.Query()
	query.Add("api_key", apiKey)
	url.RawQuery = query.Encode()

	return url, nil
}

// FetchResource thin wrapper over http + json funcs to prevent repetition in code
func FetchResource[T any](url string) (T, string, error) {
	var apiResponse T
	var responseBody string

	resp, err := http.Get(url)
	if err != nil {
		return apiResponse, responseBody, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return apiResponse, responseBody, err
	}

	responseBody = string(body)
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return apiResponse, responseBody, err
	}

	return apiResponse, responseBody, nil
}
