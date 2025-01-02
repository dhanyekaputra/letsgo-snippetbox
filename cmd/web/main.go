package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// create files server which serves files out of the directory
	// relative to the directory root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetcreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
