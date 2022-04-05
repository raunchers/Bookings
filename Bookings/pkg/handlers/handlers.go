package handlers

import (
	"github.com/raunchers/Bookings/pkg/Models"
	"github.com/raunchers/Bookings/pkg/config"
	"github.com/raunchers/Bookings/pkg/render"
	"net/http"
)

// Repo the repo used by the handlers
var Repo *Repository

// Repository is the repo type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repo for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr

	// Site wide config (m)
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &Models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, about page."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the temp
	render.RenderTemplate(w, "about.page.tmpl", &Models.TemplateData{
		StringMap: stringMap,
	})
}
