package writingtests

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func writingTests() {
	util.PrintCmd("Writing tests", `
		- use "go test" to run tests, "go test -v" for verbose, "go test -cover" for coverage, "go test ./..." to run all tests in the project
		- use "go test -coverprofile=coverage.out && go tool cover -html" to show coverage misses
		- create "filename_test.go" to test a file, then "TestFuncname" to test a function
		- "go test src/7.writingTests/project/cmd/web/*.go"
		- create "setup_test.go" file to run and store variables before every test run
		- tested the main package, "GET" and "POST" handlers and forms

		Add executeble script to the project
		- add "run.sh" at the root level of the project
		- use "go build" and point to the path of the main package folder
		- the "go build" command will not build the test files
		- make the script executable by running "chmod +x ./run.sh", and then run it: "./run.sh"
	`)
}

func Tests() {
	writingTests()
}