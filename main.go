package main

import (
	"log"
	"net/http"
)

// Handler. This is in charge of writing a response when a client request wants "/" and everything else
// Turns out "/" in the servermux world is a catch-all: it will catch every request that
// doesn't have a handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet.."))
}

func main() {
	// Request router, which is call servemux in Go Land. This component is in charge
	// of routing requests to the correct handler based on which resource is requested
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// HTTP Server. Listens on TCP port 4000, using the router
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
