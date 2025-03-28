package bsm

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGetAPIURL(t *testing.T) {
	err := os.Setenv("BSM_API_URL", "https://api.example.com")
	if err != nil {
		t.Fatalf("failed to set BSM_API_URL")
	}
	defer func() {
		err := os.Unsetenv("BSM_API_URL")
		if err != nil {
			t.Fatalf("failed to unset BSM_API_URL")
		}
	}()

	resource := "test-resource"
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	apiKey := "test-api-key"

	expectedURL := "https://api.example.com/test-resource?api_key=test-api-key&param1=value1&param2=value2"

	// Call the function
	result := GetAPIURL(resource, params, apiKey)

	// Verify the result
	if result.String() != expectedURL {
		t.Errorf("Expected %s, got %s", expectedURL, result.String())
	}
}

func TestFetchResource(t *testing.T) {
	type MockResponse struct {
		Message string `json:"message"`
	}

	mockData := MockResponse{
		Message: "Hello, World!",
	}

	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(mockData)
		if err != nil {
			t.Fatalf("failed to encode response")
		}
	}))
	defer server.Close()

	apiResponse, rawBody, err := FetchResource[MockResponse](server.URL)

	// Verify no error occurred
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Normalize and compare raw body
	var expectedBody, actualBody map[string]interface{}
	_ = json.Unmarshal([]byte(rawBody), &actualBody)
	_ = json.Unmarshal([]byte(`{"message":"Hello, World!"}`), &expectedBody)

	if !reflect.DeepEqual(expectedBody, actualBody) {
		t.Errorf("Expected raw body %+v, got %+v", expectedBody, actualBody)
	}

	// Verify the API response
	if !reflect.DeepEqual(apiResponse, mockData) {
		t.Errorf("Expected result %+v, got %+v", mockData, apiResponse)
	}
}
