package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/driver"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/handlers"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/helpers"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/models"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	defer close(app.MailChan)

	log.Println("Starting mail listener")
	listenForMail()

	fmt.Printf("Starting application on port %s", port)
	srv := &http.Server{
		Addr:    port,
		Handler: Routes(),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

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

	log.Println("Connecting to the database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=root password=root")
	if err != nil {
		log.Fatal("Cannot connect to the database! Dying..")
	}
	log.Println("Connedted to the database")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
		return nil, err
	}

	app.UseCache = false
	app.TemplateCache = tc

	render.NewRender(&app)
	handlers.InitRepo(&app, db)
	helpers.NewHelpers(&app)

	return db, nil
}
