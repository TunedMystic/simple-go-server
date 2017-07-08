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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/key", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(bind, router))
}
