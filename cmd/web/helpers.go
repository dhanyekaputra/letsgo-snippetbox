package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// the serverError helper writes and error msg
// and stack trace to the errorlog
// send generic 500 internal server error response
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// sends a specific status code and corresponding description
// used when there is a problem with the request user sent
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// send 404 to clientError function
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// retrieve the appropriate template set from the cache
	// based on the page name ("like 'home.tmpl")
	// return the error if exist
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("template %s doesn't exist", page)
		app.serverError(w, err)
		return
	}

	//initialize a new buffer
	buf := new(bytes.Buffer)

	//execute the template set to buffer
	// if there is an error call out server error
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	// writeout the provided http status code
	w.WriteHeader(status)

	buf.WriteTo(w)
}

// create a newTemplateData() helper, which returns a pointer to a templateData
// struct
func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}
