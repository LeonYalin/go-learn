package writingtests

import "github.com/LeonYalinAgentVI/go-learn/src/util"

func sendingEmails() {
	util.PrintCmd("Sending Emails", `
		- We'll use "Mailhog" as a test email tool for localhost. (added to docker-compose)
		- We'll use a channel to be able to send a mail request from everywhere in the code (see config and main.go)
		- use a "simple mail" package: "go get github.com/xhit/go-simple-mail/v2"
		- We'll use a "foundation for email" framework for styling emails
		- download a "drip" template, adjust its html and add placeholders to replace with actual content (see basic.html)
		- use a "Template" param to toggle usage of the email template file (see handlers.go)

	Adding authentication
	- We'll add a new login page, route and a handler including methods for users CRUD in the dbrepo.
	- We'll store user passwords as hashed form, using bcrypt (see postgres.go)
	- We'll store a dummy user in the users table, add the generated password using this code:
		password := "password"
		p, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
		fmt.Println(string(p))
	- Then, we'll add a "user_id" variable as a DefaultData which is available to evety page
	- We'll add a login/logout flow, which will create/destroy a user session
	- We'll create a protected page, which will only be abilable to logged-in users (admin-dashboard)
	- Also see the "admin" route in routes.go, and an "Auth" middelware
	`)
}

func adminDashboard() {
	util.PrintCmd("Admin Dashboard", `
		- We'll use this free template for admin: https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template
		- Download, unzip & copy to /static/admin, create a link in the header menu
		- create pages for all reservations, new reservations and reservation calendar.
		- add a js library for handling tables: https://github.com/fiduswriter/Simple-DataTables
	`)
}

func Emails() {
	sendingEmails()
	adminDashboard()
}
