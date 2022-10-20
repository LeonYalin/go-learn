package routes

import (
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/handlers"
	mdlware "github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const staticPath = "./src/6.convertingToGoTemplates/project/static/"

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(mdlware.NoSurf)
	router.Use(mdlware.SessionLoad)

	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	router.Get("/contact", handlers.Repo.Contact)
	router.Get("/generals-quarters", handlers.Repo.Generals)
	router.Get("/majors-suite", handlers.Repo.Majors)

	router.Get("/search-availability", handlers.Repo.Availability)
	router.Post("/search-availability", handlers.Repo.PostAvailability)
	router.Post("/search-availability-json", handlers.Repo.SearchAvailabilityJson)

	router.Get("/make-reservation", handlers.Repo.Reservation)
	router.Post("/make-reservation", handlers.Repo.PostReservation)
	router.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir(staticPath))
	router.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return router
}
