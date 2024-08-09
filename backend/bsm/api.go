package bsm

import (
    "log"
    "net/url"
    "os"
)

func GetAPIURL(resource string, params map[string]string) *url.URL {
	urlString := os.Getenv("BSM_API_URL")
	urlString = urlString + "/" + resource
	
    reqURL, err := url.Parse(urlString)
    if err != nil {
        log.Print(err)
    }

    query := reqURL.Query()
    query.Add("api_key", os.Getenv("BSM_API_KEY"))
    
    for key, value := range params {
        query.Add(key, value)
    }
    
    reqURL.RawQuery = query.Encode()
	
	return reqURL
}
