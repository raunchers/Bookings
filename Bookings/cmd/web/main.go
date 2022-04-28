package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/raunchers/Bookings/internal/config"
	"github.com/raunchers/Bookings/internal/handlers"
	"github.com/raunchers/Bookings/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main app function
func main() {

	// Change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // Created sessions last 24 hours
	session.Cookie.Persist = true     // Cookie persist after browser closing
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // Insist that cookies are encrypted if true

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create temp cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app on port: %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
