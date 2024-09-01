package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containig
// "Hello from Snippetbox" as response body
// *http.Request parameter is a pointer to a struct which holds information
// about the current request (like the HTTP method and the URL being requested)
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// more routes
// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display spific snippt ..*"))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	//Go’s servemux has different matching rules depending on
	//whether a route pattern ends with a trailing slash or not.
	//Our two new route patterns — "/snippet/view" and "/snippet/create" — don’t end in a
	//trailing slash. When a pattern doesn’t have a trailing slash, it will only be matched (and the
	//corresponding handler called) when the request URL path exactly matches the pattern in
	//full.
	//When a route pattern ends with a trailing slash — like "/" or "/static/" — it is known as a
	//subtree path pattern. Subtree path patterns are matched (and the corresponding handler
	//called) whenever the start of a request URL path matches the subtree path
	
	// mux.HandleFunc("/", home)
	mux.HandleFunc("/{$}", home) // Restrict this route to exact matches on / only.
	mux.HandleFunc("/snippt/view", snippetView)
	mux.HandleFunc("/snippt/create", snippetCreate)

	// print a log message to say the the server is starting
	log.Print("starting server on :4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	// Go’s servemux treats the route
	// pattern "/" like a catch-all. So at the moment all HTTP requests to our server will be
	//handled by the home function, regardless of their URL path. For instance, you can visit
	//a different URL path like http://localhost:4000/foo/bar and you’ll receive exactly
	//the same response.
}