package routes

import (
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/handlers"
	mdlware "github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(mdlware.NoSurf)
	router.Get("/", handlers.Repo.Home)
	router.Get("/", handlers.Repo.About)
	return router
}
