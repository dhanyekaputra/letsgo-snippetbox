package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.net/internal/models"
)

// define a templateData to act as holding structure
// for any dynamic data that we want to pass to our
// HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
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

		// file := []string{
		// 	"./ui/html/base.tmpl.html",
		// 	"./ui/html/partials/nav.tmpl.html",
		// 	page,
		// }

		// parse the base template file into a template set
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
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
