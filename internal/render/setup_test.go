package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/rajath002/bookings/internal/config"
	"github.com/rajath002/bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func TestMain(m *testing.M) {
	// what I'm going to put in session
	gob.Register(models.Reservation{})

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	// change this to true when in Production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (tw *myWriter) WriteHeader(i int) {}
