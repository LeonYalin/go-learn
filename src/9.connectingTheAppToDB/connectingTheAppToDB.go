package writingtests

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func connectingTheAppToDB() {
	util.PrintCmd("Connecting the app to the Database", `
		- We'll create a "driver" package to working with DB in a generic way
		- the built-in "database/sql" package provides a generic interface around SQL
		- for psql, we'll use a driver from the "github.com/jackc/pgx/v4" package
		- look on "repository" and "handlers" folders for details
	`)
}

func AppToDb() {
	connectingTheAppToDB()
}