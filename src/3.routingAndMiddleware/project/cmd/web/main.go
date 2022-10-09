package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/config"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/handlers"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/render"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/routes"
)

const port = ":8080"

func main() {
	var ac config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	ac.UseCache = false
	ac.TemplateCache = tc
	render.NewTemplates(&ac)
	handlers.InitRepo(&ac)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr: port,
		Handler: routes.Routes(),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
