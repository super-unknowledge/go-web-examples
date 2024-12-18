package main

import(
	"fmt"
	"net/http"
)

// Jobs of a basic http server:
// process dynamic requests => requests from users like login,
// posting images
// serve static assets => like javascript, css, and images
// accept connections => the http server must listen on a 
// specific port to be able to accpet connections from the
// internet

// http.HandleFunc says
// "when this path is matched, execute this function"
// the http.Request r contains all info about the request
// GET with r.URL.Query().Get("token")
// POST with r.FormValue("email")

// the inbuilt http.FileServer points to a url path
// which is where it will serve files from
// We need to strip part of the url path, usually
// the name of the directory the files live in
// i don't know if this part works based on testing
// when i go to localhost:8080/static I get a 404
// is it supposed to be like that?

// then we ListenAndServe to accept connections

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
