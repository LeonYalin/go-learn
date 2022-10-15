package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate/project/pkg/config"
	"github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate/project/pkg/handlers"
	"github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate/project/pkg/routes"
	"github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate/project/render"
	"github.com/LeonYalinAgentVI/go-learn/src/5.buildingProjectTemplate/project/pkg/middleware"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"
var app config.AppConfig

func main() {
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
