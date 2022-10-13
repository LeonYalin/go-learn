package managingsessions

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func managingSessions() {
	util.PrintCmd("Managing Sessions", `
		Session is a storage for the user browsing session (open tab with the app)
	  Managing sessions is done using external packages, e.g. "gorilla" or scs "go get github.com/alexedwards/scs/v2"
		see "main.go" for details
	`)
}

func Sessions() {
	managingSessions()
}
