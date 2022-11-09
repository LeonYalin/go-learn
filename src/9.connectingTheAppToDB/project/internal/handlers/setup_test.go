package handlers

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/models"
	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager

var staticPath = "./../../static/"
var tmplPath = "./../../templates/"

func getRoutes() http.Handler {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	app.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = true
	app.TemplateCache = tc

	render.NewTemplates(&app)
	InitRepo(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.SearchAvailabilityJson)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir(staticPath))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

// install the github.com/justinas/nosurf package
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	abs_fname, err := filepath.Abs(tmplPath + "*.page.*")

    if err != nil {
        log.Fatal(err, abs_fname)
    }

	// get all .page.gohtml files
	pages, err := filepath.Glob(tmplPath + "*.page.*")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(tmplPath + "*.layout.*")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob(tmplPath + "*.layout.*")
		}

		myCache[name] = tmpl
	}

	return myCache, nil
}
