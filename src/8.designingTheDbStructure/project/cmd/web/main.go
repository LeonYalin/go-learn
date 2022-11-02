package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure/project/internal/handlers"
	"github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure/project/internal/helpers"
	"github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure/project/internal/models"
	"github.com/LeonYalinAgentVI/go-learn/src/8.designingTheDbStructure/project/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger 
var errorLog *log.Logger 

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application on port %s", port)
	srv := &http.Server{
		Addr:    port,
		Handler: Routes(),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.UseCache = false
	app.TemplateCache = tc

	render.NewTemplates(&app)
	handlers.InitRepo(&app)
	helpers.NewHelpers(&app)

	return nil
}
