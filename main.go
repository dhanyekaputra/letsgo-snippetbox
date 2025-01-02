package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello From Snippetbox"))

}

// snippetview handleer function
func snippetview(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snipeet"))
}

// snippetcreate handler function
func snippetcreate(w http.ResponseWriter, r *http.Request) {
	//Use r.Method to check whether the request is using POST or not
	if r.Method != "POST" {
		// if not use WriteHeader() method to send a 405 status
		// code and Write() method to write message "Method Not Allowed"
		// response body. Return from the function so that the subsequent code
		// is not executed
		w.Header().Set("Allow", "POST")
		/// w.WriteHeader(405)
		/// w.Write([]byte("Method Not Allowed"))

		//Use the http.Error() function to send a 405 status code
		// and "Method not allowed" strings as the response body
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
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
