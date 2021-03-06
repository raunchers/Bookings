package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"github.com/raunchers/Bookings/internal/Models"
	"github.com/raunchers/Bookings/internal/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData sets the default data
func AddDefaultData(td *Models.TemplateData, r *http.Request) *Models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate renders templates using html/tmpl
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *Models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// Get the temp cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Bytes buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	// Render the page
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing templ to browser", err)
	}

}

// CreateTemplateCache Creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
