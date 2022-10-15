package buildingprojecttemplate

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func buildingProjectTemplate() {
	util.PrintCmd("Working with forms", `
	Enabling static files
	- use router fileServer to serve a static folder, e.g. "static"
	- see "static", "working-html" folders and "routes.go" for details
	`)
}

func ProjectTemplate() {
	buildingProjectTemplate()
}
