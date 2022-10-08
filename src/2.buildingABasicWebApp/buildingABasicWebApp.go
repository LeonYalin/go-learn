package buildingabasicwebapp

import (
	"github.com/LeonYalinAgentVI/go-learn/src/util"
)

func helloWorldWeb() {
	util.PrintCmd("Basic web server code - hello world", `
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			n, err := fmt.Fprintf(w, "Hello, World")
			if err != nil {
				fmt.Printf("Number of bytes received is %d", n)
			}
		})
		http.ListenAndServe(":8080", nil)
	`)
}

func addingRoutes() {
	util.PrintCmd("Adding routes", `
		const port = ":8080"
		Home := func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "This is the Home page")
		}
		About := func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "This is the about page")
		}

		http.HandleFunc("/", Home)
		http.HandleFunc("/about", About)
		http.ListenAndServe(port, nil)
	`)
}

func servingHtmlTemplates() {
	util.PrintCmd("Serving HTML Templates", `
		See the 'project' folder for example

		const port = ":8080"
		renderTemplate := func(w http.ResponseWriter, tpl string) {
			parsedTemplate, parsingErr := template.ParseFiles("./src/2.buildingABasicWebApp/project/" + tpl)
			if parsingErr != nil {
				fmt.Fprint(w, "error parsing template", parsingErr)
				return
			}
			servingErr := parsedTemplate.Execute(w, nil)
			if servingErr != nil {
				fmt.Fprint(w, "error serving template", servingErr)
				return
			}
		}
		Home := func(w http.ResponseWriter, r *http.Request) {
			renderTemplate(w, "home.page.gohtml")
		}
		About := func(w http.ResponseWriter, r *http.Request) {
			renderTemplate(w, "about.page.gohtml")
		}
	
		http.HandleFunc("/", Home)
		http.HandleFunc("/about", About)
		http.ListenAndServe(port, nil)
	`)
}

func addingBaseLayout() {
	util.PrintCmd("Using a base layout", `
	  See templates/base.layout.gohtml for details
	`)
}

func addingTemplateCache() {
	util.PrintCmd("Adding template cache", `
	  See render.go for details
	`)
}

func addingAppConfig() {
	util.PrintCmd("Adding app config", `
	  See config.go for details
	`)
}

func passingTemplateData() {
	util.PrintCmd("Passing template data", `
	  See about.page.go and handlers.go for details
	`)
}

func BasicWebApp() {
	helloWorldWeb()
	addingRoutes()
	servingHtmlTemplates()
	addingBaseLayout()
	addingTemplateCache()
	addingAppConfig()
	passingTemplateData()
}
