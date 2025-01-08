package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.net/internal/models"
)

// define a templateData to act as holding structure
// for any dynamic data that we want to pass to our
// HTML templates
type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// create a humanDate function which returns a nicely formatted string
// representation of time.Time object
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// use the filepath.Glob() to get a slice of all filepath
	// that match the patter "./ui/html/pages/*.tmpl"
	// [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//extract the file name (like "home.tmpl") from the full filepath
		//assign it to the name variable
		name := filepath.Base(page)

		//template.FuncMap Must be registered with
		//template set before parsefiles method
		// use template.New() to create an empty template
		// set
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}
		// file := []string{
		// 	"./ui/html/base.tmpl.html",
		// 	"./ui/html/partials/nav.tmpl.html",
		// 	page,
		// }

		// parse the base template file into a template set
		ts, err = ts.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// call parsegblob *on this template set* to add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// call parsefiles *on this template sed* to add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
