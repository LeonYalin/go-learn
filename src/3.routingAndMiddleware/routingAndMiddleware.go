package routingandmiddleware

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func routing() {
	util.PrintCmd("Routing", `
		- use "go get <package-name>" to install a package
		- this will add a new entry in a "go.mod" file
		- we can use the "pat" package for simple routing, or "chi" for more complex one
		- use "go mod tidy" to clean unused packages 
	`)
}

func middleware() {
	util.PrintCmd("Middleware", `
		look on "pkg/middleware/middleware.go" for details
	`)
}

func RoutingMiddleware() {
	routing()
	middleware()
}