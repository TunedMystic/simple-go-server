package main

import (
	"net/http"
	"net/http/httptest"
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
