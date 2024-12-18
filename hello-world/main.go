package main

import (
	"fmt"
	"net/http"
)

// The func passed to HandleFunc is a request handler
// It receives all incoming HTTP connections from 
// browsers, HTTP clients, or API requests
// http.Response => where you write your text/html response to
// http.Request => contains all info about this HTTP request
// including things like the URL or header fields

// When we call HandleFunc and pass the request handler
// we've created, we register that request handler to
// the default HTTP server built into go's net/http lib

// To accpet any HTTP connections from the outside, an
// HTTP server must listen on a port to pass connections
// to the request handler.

// my laptop doesn't like port 80 for some reason
// localhost:8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}

