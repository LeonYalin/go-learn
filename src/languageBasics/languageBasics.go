package languageBasics

import (
	"fmt"
	"github.com/LeonYalinAgentVI/go-learn/src/util"
)

func gettingStarted() {
	util.PrintCmd("Getting started", `
	Basic "Hello world" - fmt.Println("Hello, World")
	To run a program, just run 'go run main.go' from the cmd.
	To compile a program, run 'go build -o hw main.go' from the cmd. That will create a 'hw' binary file. To run it, type './hw' from the cmd.
	`)

	var msg string = "Leon Yalin" // use variables
	fmt.Println(sayHello(msg))
}

func sayHello(msg string) string {
	return msg
}

func variablesAndFunctions() {
	util.PrintCmd("Variable & Functions", `
	// TODO:
	`)
}

func LanguageBasics() {
	gettingStarted()
	variablesAndFunctions()
}
