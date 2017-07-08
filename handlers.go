package main

import (
	"fmt"
	"net/http"
)

// RootEndpoint handler.
func RootEndpoint(w http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("Secret Key: %v\n", RandString(24))
	w.Write([]byte(message))
}
