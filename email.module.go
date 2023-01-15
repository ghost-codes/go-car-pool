package main

import (
	"os"

	"gopkg.in/gomail.v2"
)

type SendMailData struct {
	To         string
	Subject    string
	Body       string
	Attachment []string
}

func SendMail(data SendMailData) error {
	fromEmail := os.Getenv("fromEmail")
	gmailPassword := os.Getenv("gmailPassword")

	msg := gomail.NewMessage()
	msg.SetHeader("From", fromEmail)
	msg.SetHeader("To", data.To)
	msg.SetHeader("Subject", data.Subject)
	msg.SetBody("text/html", data.Body)
	for _, img := range data.Attachment {
		msg.Attach(img)
	}

	dialer := gomail.NewDialer("smtp.gmail.com", 587, fromEmail, gmailPassword)

	return dialer.DialAndSend(msg)
}
