package convertingtogotemplates

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func convertingToGoTemplates() {
	util.PrintCmd("Converting to Go templates", `
		- open "/templates" folder to see changes
	`)
}

func workingWithForms() {
	util.PrintCmd("Working with forms", `
	- rendering forms
	- using server-side form validation
	- using js with ajax
	- showing errors and warnings to the client
	`)
}

func ConvertToGoTemplates() {
	convertingToGoTemplates()
	workingWithForms()
}