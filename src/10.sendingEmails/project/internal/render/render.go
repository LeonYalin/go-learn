package render

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/config"
	"github.com/LeonYalinAgentVI/go-learn/src/10.sendingEmails/project/internal/models"
	"github.com/justinas/nosurf"
)

var tmplPath = "./src/10.sendingEmails/project/templates/"

var app *config.AppConfig

func NewRender(ac *config.AppConfig) {
	app = ac
}

// functions that will be passed templates
var functions = template.FuncMap{
	"humanDate": func(t time.Time) string {
		return t.Format("2006-01-02")
	},
	"formatDate": func(t time.Time, f string) string {
		return t.Format(f)
	},
	"iterate": func(count int) []int {
		var i int
		var items []int
		for i = 0; i < count; i++ {
			items = append(items, i)
		}
		return items
	},
	"add": func(a, b int) int {
		return a + b
	},
}

func Template(w http.ResponseWriter, r *http.Request, tpl string, td *models.TemplateData) error {

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
		msg := "Could not get template from template cache"
		return errors.New(msg)
	}

	// execute check: execute the template to buffer
	buf := new(bytes.Buffer)
	err = t.Execute(buf, addDefaultData(td, r))
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	// write the result back to writer
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
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
		tmpl, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob(tmplPath + "*.layout.*")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob(tmplPath + "*.layout.*")
			if err != nil {
				return myCache, err
			}
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
	if app.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}
	return td
}
