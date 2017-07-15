package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestInfoHandler(t *testing.T) {
	// Make a new request.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	// Create a new ResponseRecorder to record the response.
	rr := httptest.NewRecorder()

	// Make the http handler.
	handler := http.HandlerFunc(InfoHandler)

	/*
		The `handler` satisfies the `http.Handler` interface, so we
		call their ServeHTTP method directly, passing in the request
		and response recorder.
	*/
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Invalid status code. Got %v, expected %v", status, http.StatusOK)
	}
}

func TestInfoHandler2(t *testing.T) {
	// Build query params.
	queryArgs := url.Values{}
	queryArgs.Add("page", "1")
	// Alternative.
	// queryArgs := url.Values{
	//     "page": {"1"}, "sort", {"name"},
	// }

	// Make a new request.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}

	// Add query parameters to request url.
	req.URL.RawQuery = queryArgs.Encode()
	// Create a new ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	// Make the http handler.
	handler := http.HandlerFunc(InfoHandler)
	// Process handler
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Invalid status code. Got %v, expected %v", status, http.StatusOK)
	}

	// Make an arbitrary interface to
	// decode the response body into.
	var responseData map[string]interface{}

	// Decode the response body.
	json.NewDecoder(rr.Body).Decode(&responseData)

	// Check that the response has the 'page' key.
	if _, ok := responseData["page"]; !ok {
		t.Errorf("Expected 'page' in JSON response.")
	}
}
