package main

import (
	"fmt"
	"net/http"
	"time"
)

// DateMiddleware struct.
type DateMiddleware struct {
	handler http.Handler
}

func (d *DateMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.UnixDate)
	fmt.Printf("[date-middleware] %v\n", now)
	d.handler.ServeHTTP(w, r)
}

// NewDateMiddleware is a help function which wraps
// a given Http handler in a DateMiddleware type.
func NewDateMiddleware(handler func(http.ResponseWriter, *http.Request)) *DateMiddleware {
	return &DateMiddleware{http.HandlerFunc(handler)}
}
