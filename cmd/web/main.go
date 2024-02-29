package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kylelaverty/bookings/internal/config"
	"github.com/kylelaverty/bookings/internal/driver"
	"github.com/kylelaverty/bookings/internal/handlers"
	"github.com/kylelaverty/bookings/internal/helpers"
	"github.com/kylelaverty/bookings/internal/models"
	"github.com/kylelaverty/bookings/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

const portNumber = ":8080"

// main is the entry point for the application.
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// What will be placed in the session.
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	// Change this to true when in production
	app.InProduction = false

	// Create a logger
	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.LstdFlags|log.Lshortfile)
	app.ErrorLog = errorLog

	// Setup a session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // Set to true in production

	app.Session = session

	// Connect to the database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=XXXXXX sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
		return nil, err
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
