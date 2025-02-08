package main

import (
	"log"
	"net/http"
)

// Handler. This is in charge of writing a response when a client request wants "/"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Request router, which is call servemux in Go Land. This component is in charge
	// of routing requests to the correct handler based on which resource is requested
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// HTTP Server. Listens on TCP port 4000, using the router
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
