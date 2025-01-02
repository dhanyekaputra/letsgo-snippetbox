package main

import (
	"log"
	"net/http"
)

func main() {
	// Register the two new handler function to corresponding URL patterns
	// with servemux

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetcreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
