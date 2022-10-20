package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/6.convertingToGoTemplates/project/internal/models"
	"github.com/justinas/nosurf"
)

const tmplPath = "./src/6.convertingToGoTemplates/project/templates/"

var app *config.AppConfig

func NewTemplates(ac *config.AppConfig) {
	app = ac
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get template from cache
	t, ok := tc[tpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// execute check: execute the template to buffer
	buf := new(bytes.Buffer)
	err = t.Execute(buf, addDefaultData(td, r))
	if err != nil {
		log.Printf(err.Error())
	}

	// write the result back to writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err.Error())
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all .page.gohtml files
	pages, err := filepath.Glob(tmplPath + "*.page.*")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(tmplPath + "*.layout.*")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob(tmplPath + "*.layout.*")
		}

		myCache[name] = tmpl
	}

	return myCache, nil
}

func addDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	// PopString is a one-time add
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRFToken = nosurf.Token(r)
	return td
}
