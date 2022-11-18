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
	`)
}

func Emails() {
	sendingEmails()
}