package routes

import (
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/5.workingWithForms/project/pkg/handlers"
	mdlware "github.com/LeonYalinAgentVI/go-learn/src/5.workingWithForms/project/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const staticPath = "./src/5.workingWithForms/project/static/"

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(mdlware.NoSurf)
	router.Use(mdlware.SessionLoad)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return router
}
