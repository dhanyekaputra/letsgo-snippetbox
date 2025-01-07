package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	//import the model package we just create
	"snippetbox.net/internal/models"
)

// define an application struct to hold the application-wide
// dependencies for the web application

// add snippets field
// this will allow us to make the snippetmodel object available
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
	// add a templateCache field
	templateCache map[string]*template.Template
}

const (
	username = "admin"
	password = "admin"
	hostname = "127.0.0.1:3306"
	dbname   = "snippetbox"
)

// function open the DB connection
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	//define command line with name addr, default 4000
	// and some explaining what the flag control
	// flag will be stored in the addr cariable at runtime
	// VALUE IN POINTER
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbname)

	flag.Parse()

	//use log.net to create a logger for writing info message
	//3 param: destination, string, additional info to input
	//flag joined using OR operator
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	//instead, use stderr as destination, and Lshortfile to include relevant file
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//pass OpenDB value to db and check error
	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// we also defer a call to db.Close, so connection closed
	// before the main function exit
	defer db.Close()

	// initizalize new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	//initialize new instance of application struct
	// initialize a model.SnippetModel instance and add it
	//to application dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
		// add to app dependencies
		templateCache: templateCache,
	}

	// initialize a new http.server struct. we set the addr and handler fields
	// so that the server uses the same network address and routed ass before
	// errorlog used is errorLog we created before
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
