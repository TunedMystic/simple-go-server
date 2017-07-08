package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// RootHandler returns some basic information in JSON format.
func RootHandler(w http.ResponseWriter, request *http.Request) {
	// Process form data.
	request.ParseForm()
	// Make response map.
	responseData := make(map[string]string)

	// Set server name.
	responseData["server"] = "Simple Server"

	// Get the current date and format it.
	dateTime := time.Now().Format(time.UnixDate)
	responseData["date"] = dateTime

	// Access the `page` query param using `request.Form.Get()`
	// This returns the first value associated with the key.
	if page := request.Form.Get("page"); page != "" {
		responseData["page"] = page
	}

	// Access the `sort` query param using the `request.Form` map.
	// This map has all values associated with the key.
	if sort, ok := request.Form["sort"]; ok {
		responseData["sort"] = strings.Join(sort, ", ")
	}

	// Encode response.
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(responseData)
}

// GenerateKeyHandler returns a randomly generated string.
func GenerateKeyHandler(w http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("Random Key: %v\n", RandString(24))
	w.Write([]byte(message))
}
