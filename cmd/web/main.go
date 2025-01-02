package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//define command line with name addr, default 4000
	// and some explaining what the flag control
	// flag will be stored in the addr cariable at runtime
	//VALUE IN POINTER
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	flag.Parse()

	//use log.net to create a logger for writing info message
	//3 param: destination, string, additional info to input
	//flag joined using OR operator
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	//instead, use stderr as destination, and Lshortfile to include relevant file
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetview)
	mux.HandleFunc("/snippet/create", snippetcreate)

	// initialize a new http.server struct. we set the addr and handler fields
	// so that the server uses the same network address and routed ass before
	// errorlog used is errorLog we created before
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
