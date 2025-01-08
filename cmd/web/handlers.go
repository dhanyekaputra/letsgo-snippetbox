package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.net/internal/models"
)

// define handler so its defined as a method
// against struct *application
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	///// USING MODELS LATEST() IN HANDLER
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// call the newTemplateData() helper to get a templateData struct containing
	// the default data and add the snippet slice to it
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// use the new render helpers
	app.render(w, http.StatusOK, "home.tmpl.html", data)
}

// snippetview handleer function
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	//use the snippetmodel objects get method to retrieve the data for specific record
	//based on its id
	//return 404 not found response
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// call the newTemplateData() helper to get a templateData struct containing
	// the default data and add the snippet slice to it
	data := app.newTemplateData(r)
	data.Snippet = snippet

	// use the new render helpers
	app.render(w, http.StatusOK, "view.tmpl.html", data)

}

// snippetcreate handler function
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// create dummy data
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	//pass the data to the SnippetModel.Insert() method
	//receiving the ID of the new record back
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//redirect to relevant page
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
