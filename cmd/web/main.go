package main

import (
	"log"
	"net/http"
)

func main() {
	// Router
	mux := http.NewServeMux()

	// Handle that creates a file server that will serve the contents
	// of the directory ./ui/static. Note the relative path to the
	// project
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// http.FileServer will remove the leading "/" from the URL that's
	// passed along and then search for the requested file by concatenating
	// the URL and the directory. For example, if the request is:
	// http://localhost:4000/css/boom.css
	// The resulting search path will be ./ui/static/css/boom.css
	// If on the other hand, the URL is:
	// http://localhost:4000/static/css/boom.css
	// The resulting search path will be ./ui/static/static/css/boom.css
	// which does not exist. So we need to remove the extra "/static"
	// before it's handled, which http.StripPrefix does
	// http.StripPrefix returns another handler but one that has whatever
	// prefix is specified before handling its business
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	// Displays a specific snippet
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	// Shows form for creating a snipppet
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// When a snippet is created
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Secret zip file
	mux.HandleFunc("GET /secret/zip/file", downloadHandler)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
