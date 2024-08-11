package bsm

import (
    "encoding/json"
    "io"
    "log"
    "net/http"
    "net/url"
    "os"
)

// GetAPIURL Construct a BSM API URL from a given path, appending all query parameters and the API key automatically
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