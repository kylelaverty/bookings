package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kylelaverty/bookings/internal/config"
	"github.com/kylelaverty/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// What will be placed in the session.
	gob.Register(models.Reservation{})

	// Change this to true when in production
	testApp.InProduction = false

	// Create a logger
	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.LstdFlags|log.Lshortfile)
	testApp.ErrorLog = errorLog

	// Setup a session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction // Set to true in production

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWriter) WriteHeader(i int) {}

func (mw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
