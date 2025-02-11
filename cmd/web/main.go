package main

import (
	"log"
	"net/http"
)

func main() {
	// Router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	// Displays a specific snippet
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	// Shows form for creating a snipppet
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// When a snippet is created
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
