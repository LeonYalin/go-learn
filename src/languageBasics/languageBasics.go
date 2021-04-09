package languageBasics

import (
	// "fmt"
	// "log"
	// "time"
	// "sort"

	"fmt"
	"log"

	"github.com/LeonYalinAgentVI/go-learn/src/util"
)

func gettingStarted() {
	util.PrintCmd("Getting started", `
	Basic "Hello world" - fmt.Println("Hello, World")
	To run a program, just run 'go run main.go' from the cmd.
	To compile a program, run 'go build -o hw main.go' from the cmd. That will create a 'hw' binary file. To run it, type './hw' from the cmd.
	`)
}

// var packageVar = 2

func variablesAndFunctions() {
	util.PrintCmd("Variable & Functions", `
	- Like in other languages, we have a couple of primitive types:

	var name string = "Leon Yalin"
	var num int = 0

	func sayHello(mas string) {
		fmt.Println(msg)
	}

	var msg string = "Leon Yalin" // use variables
	fmt.Println(sayHello(msg))

	func returnsTwoValues(msg string) (string, string) {
		return msg, "lala"
	}

	// returning multiple values, _ for omitting variables
	var firstVal, _ = returnTwoValues("lulu")
	fmt.Println(firstVal)

	We can have a variable on a package scope, like in Javascript.

	Alternative syntax to create a variables: a := "lala"
	The compiler inferres the variable type and declares it.
	`)
}

func pointers() {
	util.PrintCmd("Pointers", `
	In go, we have pointers and references. References can be used to to get variable memory addresses.
	Pointers are used to get the value stored in particular memory address (reference).

	Example: change string value using pointers (by reference)
	
	var text string = "Hello, Bro"
	log.Println("Printing a reference to a variable:", &text)
	var newText *string = &text
	*newText = "Bye, Broh"
	log.Println(*newText, text)
	`)
}

func structs() {
	util.PrintCmd("Structs", `
	Structs are sets of valuues used together as a structure (same as "type" keyword in typescript).

	type Person struct {
		FirstName string
		LastName  string
		Birthday  time.Time
	}

	person1 := Person{
		FirstName: "Leon",
		LastName:  "Yalin",
	}
	log.Println(person1)
	`)
}

func receivers() {
	util.PrintCmd("Receivers", `
	Receivers are "structs with functions": functions that can be bound to structs
	
	type Person struct {
		FirstName string
		LastName string
	}
	
	func (p *Person) fullName() string {
		return fmt.Sprintf("Person: %s %s", p.FirstName, p.LastName)
	}

	person1 := Person {
		FirstName: "Leon",
		LastName: "Yalin",
	}
	log.Println(person1.fullName())
	`)
}

func maps() {
	util.PrintCmd("Maps", `
	We can create Maps in Go. Maps are similar to maps in other languages.

	myMap := make(map[string]string)
	myMap["first"] = "First"
	myMap["second"] = "Second"
	log.Println(myMap)

	Maps work with structs. For example:

	type Person struct {
		FirstName string
		LastName string
	}
	myMap2 := make(map[string]Person)
	myMap2["leon_yalin"] = Person{
		FirstName: "Leon",
		LastName: "Yalin",
	}
	log.Println(myMap2)
	`)
}

func slices() {
	util.PrintCmd("Slices", `
	Slices are similar to slices in Python. Slices are arrays with boundaries.

	var mySlice []string
	mySlice = append(mySlice, "First", "Second", "Third", "Fourth")
	log.Println(mySlice)
	
	We can sort slices

	sort.Strings(mySlice)
	log.Println(mySlice)
	
	Shorthand

	mySlice2 := []int {1, 2, 3, 4, 5}
	log.Println(mySlice2)

	Use slice of slice :/

	log.Println(mySlice2[2:5])
	`)
}

func decisionStructures() {
	util.PrintCmd("Decision structures", `
	In Go, we have if-else statements & switch cases, just like in every language.

	If-else statements:

	a := 2
	if a >1 {
		log.Println("more than 1")
	} else if a < 0 {
		log.Println("less than 0")
	} else {
		log.Println("other")
	}

	Switch case statements. There is no need to "return" in switch, it exits automatically

		b := "lala"
		switch b {
		case "lala":
			log.Println("b is lala")
		case "lulu":
			log.Println("b is lulu")
		default:
			log.Println("b is unknown")
		}
	`)
}

func loops() {
	util.PrintCmd("Loops", `
	In Go, we have a for loops, and range loops (a.k.a. foreach). No While loop in Go.

	Regular for loops

	letters := []string {"a", "b", "c", "d"}
	for i := 0; i < len(letters); i++ {
		log.Println(letters[i])
	}

	For range loops a.k.a. foreach

	for i, item := range letters {
		log.Println(i, item)
	}

	For loops with maps

	myMap := make(map[string]string)
	myMap["first"] = "First"
	myMap["second"] = "Second"
	for key, value := range myMap {
		log.Println(key, value)
	}
	`)
}

func interfaces() {
	util.PrintCmd("Interfaces", `
	In Go, Interfaces are similar to Typescript interfaces.
	To implement an interface, tou just need to implement the required functions via receivers.

	type Person struct {
		FirstName string
		LastName string
	}
	
	func (p Person) sayHello() string {
		return fmt.Sprintf("Hello, I am %s %s", p.FirstName, p.LastName)
	}
	
	type Greet interface {
		sayHello() string
	}
	
	func greetSomeone(g Greet) {
		g.sayHello()
	}

	person1 := Person {
		FirstName: "Leon",
		LastName: "Yalin",
	}
	greetSomeone(person1)
	`)

}

func LanguageBasics() {
	gettingStarted()
	variablesAndFunctions()
	pointers()
	structs()
	receivers()
	maps()
	slices()
	decisionStructures()
	loops()
	interfaces()
}
