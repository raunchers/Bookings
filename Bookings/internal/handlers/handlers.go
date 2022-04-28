package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/raunchers/Bookings/internal/Models"
	"github.com/raunchers/Bookings/internal/config"
	"github.com/raunchers/Bookings/internal/render"
	"log"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &Models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, about page."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the temp
	render.RenderTemplate(w, r, "about.page.tmpl", &Models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &Models.TemplateData{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservations.page.tmpl", &Models.TemplateData{})
}

// Generals renders the generals room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &Models.TemplateData{})
}

// Majors renders the majors room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &Models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &Models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start Date: %s, End Date: %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles room availability request and sends a JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		OK:      false,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println("Error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
