package workingwithforms

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func workingWithForms() {
	util.PrintCmd("Working with forms", `
	Enabling static files
	- use router fileServer to serve a static folder, e.g. "static"
	- see "static", "working-html" folders and "routes.go" for details
	`)
}

func Forms() {
	workingWithForms()
}