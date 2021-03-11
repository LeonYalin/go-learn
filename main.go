package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	fmt.Println("To run a program, just run 'go run main.go' from the cmd.")
	fmt.Println("To compile a program, run 'go build -o hw main.go' from the cmd. That will create a 'hw' binary file. To run it, type './hw' from the cmd")

	var msg string = "Leon Yalin" // use variables
	fmt.Println(sayHello(msg)) 
}

func sayHello(msg string) string {
	return msg
}