package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/handlers"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/middleware"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/models"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/routes"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"
var app config.AppConfig

func main() {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = false
	app.TemplateCache = tc

	render.NewTemplates(&app)
	handlers.InitRepo(&app)
	middleware.NewMiddleware(&app)
	
	fmt.Printf("Starting application on port %s", port)
	srv := &http.Server{
		Addr: port,
		Handler: routes.Routes(),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
