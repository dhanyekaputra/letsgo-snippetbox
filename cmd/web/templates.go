package main

import "snippetbox.net/internal/models"

// define a templateData to act as holding structure
// for any dynamic data that we want to pass to our
// HTML templates
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
