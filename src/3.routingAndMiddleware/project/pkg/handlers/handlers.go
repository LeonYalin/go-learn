package handlers

import (
	"net/http"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/config"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/pkg/models"
	"github.com/LeonYalinAgentVI/go-learn/src/3.routingAndMiddleware/project/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: map[string]string{
			"first": "Hello, World",
		},
		CSRFToken: "Super Secret token",
	})
}

func NewRepo(ac *config.AppConfig) *Repository {
	return &Repository{
		App: ac,
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

func InitRepo(ac *config.AppConfig) {
	NewHandlers(NewRepo(ac))
}
