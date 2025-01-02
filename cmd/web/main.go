package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//define command line with name addr, default 4000
	// and some explaining what the flag control
	// flag will be stored in the addr cariable at runtime
	//VALUE IN POINTER
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	//importantly we use flag.Parse func to parse command-line flag
	//this read the command line flag and assign to addr
	//this should be called before using addr var
	//otherwise the value will always be ":4000"
	//error will be terminated
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetcreate)

	//value is a pointer, so we need to dereference the pointer
	log.Println("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}
