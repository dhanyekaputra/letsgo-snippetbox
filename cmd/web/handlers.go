package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	// extract the value of the id param from the query string and try to\
	// convert it into an integer susing strconv.Atoi() function
	// if the value cant be converted and less than 1 then we return
	// a 404 not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display Specific Snippet With ID %d...", id)
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
