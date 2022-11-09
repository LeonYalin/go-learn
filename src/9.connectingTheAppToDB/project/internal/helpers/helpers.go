package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/LeonYalinAgentVI/go-learn/src/9.connectingTheAppToDB/project/internal/config"
)

var app *config.AppConfig

func NewHelpers(ac *config.AppConfig) {
	app = ac
}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("client error with a status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
