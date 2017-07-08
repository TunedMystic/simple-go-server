package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var bind = "localhost:8000"

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// Create router.
	router := mux.NewRouter().StrictSlash(true)

	// Configure routes.
	router.HandleFunc("/", RootHandler).Methods("Get")
	router.HandleFunc("/key", GenerateKeyHandler).Methods("GET")

	// Run server.
	log.Fatal(http.ListenAndServe(bind, router))
}
