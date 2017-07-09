package main

import (
	"fmt"
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

	// Configure routes with 'HandleFunc'.
	router.HandleFunc("/", InfoHandler).Methods("Get")
	router.HandleFunc("/key", GenerateKeyHandler).Methods("Post")

	// Configure routes with 'Handle'.
	router.Handle("/team", http.HandlerFunc(ExampleHandler)).Methods("Get")

	// Connect a route to a handler, with custom middleware.
	router.Handle("/about", &DateMiddleware{
		http.HandlerFunc(ExampleHandler),
	}).Methods("Get")

	// Connect a route to a handler, with a middleware helper function.
	router.Handle("/pricing", NewDateMiddleware(ExampleHandler))

	// 404 NotFound handler
	nf := &NotFound{"404"}
	router.NotFoundHandler = nf

	// Run server.
	fmt.Printf("Running server on %v...\n", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}

// Http middleware:
// https://stackoverflow.com/a/33403252
// http://www.alexedwards.net/blog/making-and-using-middleware
// https://github.com/gorilla/handlers
