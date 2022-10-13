package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/LeonYalinAgentVI/go-learn/src/4.managingSessions/project/pkg/config"
	"github.com/LeonYalinAgentVI/go-learn/src/4.managingSessions/project/pkg/models"
)

const tmplPath = "./src/4.managingSessions/project/templates/"

var app *config.AppConfig

func NewTemplates(ac *config.AppConfig) {
	app = ac
}

func RenderTemplate(w http.ResponseWriter, tpl string, td *models.TemplateData) {

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
	err = t.Execute(buf, addDefaultData(td))
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

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
