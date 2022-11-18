package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type ResponseWriterMock struct{}

func (rw *ResponseWriterMock) Header() http.Header {
	return http.Header{}
}

func (rw *ResponseWriterMock) WriteHeader(i int) {

}

func (rw *ResponseWriterMock) Write (b []byte) (int, error) {
	length := len(b)
	return length, nil
}