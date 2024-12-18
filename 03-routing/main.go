package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// We use gorilla/mux to create routes with named parameters,
// GET/POST handlers, and domain restrictions
// because net/http doesn't do segmenting request urls
// into single parameters very well

// The request router is the main router for your web app
// It will receive all http connections and pass it to the
// request handlers you register on it

// To register the request handler, we call HandleFunc from
// the newly created request router

// We then use mux to interact with the dynamic segments
// of our url, to get data from these segments.
// muxVars(r) takes the http.Request as a parameter and
// returns a map of the dynamic segments that we can access

// Then we pass out request router to  ListenAndServe
// telling it to use that instead of the default router
// of the net/http package

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]  // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8080", r)
}
