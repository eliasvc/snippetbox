package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler. This is in charge of writing a response when a client request wants "/" and everything else
// Turns out "/" in the servermux world is a catch-all: it will catch every request that
// doesn't have a handler
func home(w http.ResponseWriter, r *http.Request) {
	// Here we are writing the custom header 'Server':'Go' to the response header map
	// So, 'Server': 'Go' should be included in the response
	// Note that custom headers have to be done before calls to w.Write() and w.WriteHeader() as any
	// as any header added after will not be received by the client
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	// id should be a positive integer greater than 0
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with id %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet.."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// HTTP Status Code 201 is frequently used as a response for a POST operation
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet..."))
}

func main() {
	// Request router, which is call servemux in Go Land. This component is in charge
	// of routing requests to the correct handler based on which resource is requested
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// HTTP Server. Listens on TCP port 4000, using the router
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
