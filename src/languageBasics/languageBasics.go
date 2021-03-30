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
	fmt.Println(returnTwoValues("lulu"))

	// returning multiple values, _ for omitting variables
	var firstVal, _ = returnTwoValues("lulu")
	fmt.Println(firstVal)
}

func sayHello(msg string) string {
	return msg
}

func returnTwoValues(msg string) (string, string) {
	return msg, "lala"
}

func variablesAndFunctions() {
	util.PrintCmd("Variable & Functions", `
	- Like in other languages, we have a couple of primitive types:
	var name string = "Leon Yalin"
	var num int = 0
	func sayHello(mas string) {
		fmt.Println(msg)
	}
	func returnsTwoValues(msg string) (string, string) {
		return msg, "lala"
	}
	`)
}

func pointers() {
	util.PrintCmd("Pointers", `
	lala
	`)
}

func LanguageBasics() {
	gettingStarted()
	variablesAndFunctions()
	pointers()
}
