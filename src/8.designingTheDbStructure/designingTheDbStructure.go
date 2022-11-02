package writingtests

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func designingTheDbStructure() {
	util.PrintCmd("Designing the Database structure", `
		- Before starting with creating tables, we'll build the schema first, using the pgadmin "ERD tool"
		- Create all tables, draw relations between the tables and exporting the SQL at the end

		Creating a database using migrations:
		- use the "soda" package.
		- Installation instructions:
				go get github.com/gobuffalo/pop/...
				go install github.com/gobuffalo/pop/soda
				nano ~./bash_profile -> add 'export PATH="$PATH:/Users/lyalin/go/bin"'
		- after soda is installed, we'll create a database configuration file called "database.yml"
		- the "fizz" language is used to create migrations
		- use "soda generate fizz CreateUserTable" to create up and down migration files
		- write the "up" migration using fizz. Check the "migrations" folder content for details
		- use "soda migrate" to run the migration, "soda migrate down" to revert a migration
		- use "soda reset" to run add down and then up migrations
	`)
}

func DbStructure() {
	designingTheDbStructure()
}