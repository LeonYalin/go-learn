package main

import (
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const staticPath = "./src/10.sendingEmails/project/static/"

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	// available only to logged in users
	mux.Route("/admin", func(r chi.Router) {
		// mux.With(Auth).Route("/admin", func(r chi.Router) {
		r.Get("/dashboard", handlers.Repo.AdminDashboard)
		r.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		r.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		r.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		r.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)
		r.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
		r.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)

		r.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservation)
		r.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservation)
	})

	fileServer := http.FileServer(http.Dir(staticPath))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
