package middleware

import (
	"fmt"
	"net/http"

	"github.com/LeonYalinAgentVI/go-learn/src/5.workingWithForms/project/pkg/config"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewMiddleware(ac *config.AppConfig){
	app = ac
}

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("WriteToConsole middleware")
		next.ServeHTTP(w, r)
	})
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
