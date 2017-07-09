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

func testHandler(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("This is not the route you are looking for. **Waves hand**\n"))
}

// JediHandler struct
type JediHandler struct {
	Name string
}

func (j *JediHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("My name is %v and I am a Jedi\n", j.Name)))
}

// DateMiddleware struct.
type DateMiddleware struct {
	handler http.Handler
}

func (d *DateMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.UnixDate)
	fmt.Printf("[date-middleware] %v\n", now)
	d.handler.ServeHTTP(w, r)
}

func main() {
	// Create router.
	router := mux.NewRouter().StrictSlash(true)

	// Configure routes.
	router.HandleFunc("/", RootHandler).Methods("Get")
	router.HandleFunc("/key", GenerateKeyHandler).Methods("Get")
	router.Handle("/test", http.HandlerFunc(testHandler)).Methods("Get")

	// Create a new custom handler.
	jh := &JediHandler{"Anakin"}
	router.Handle("/jedi", jh).Methods("Get")

	// Connect a route to a handler, with custom middleware.
	router.Handle("/test2", &DateMiddleware{http.HandlerFunc(testHandler)}).Methods("Get")

	// Run server.
	fmt.Printf("Running server on %v...\n", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}

// Http middleware:
// https://stackoverflow.com/a/33403252
// http://www.alexedwards.net/blog/making-and-using-middleware
