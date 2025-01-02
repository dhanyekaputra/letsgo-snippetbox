package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Snippetbox"))
}

// snippetview handleer function
func snippetview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snipeet"))
}

// snippetcreate handler function
func snippetcreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet"))
}

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
