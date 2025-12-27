package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	from := "test@example.com"
	to := []string{"to@example.com"}
	subject := "Test Mail from Go"
	body := "This is a test email sent from Go to Maildev."

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: to@example.com\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

	err := smtp.SendMail(
		"localhost:1025", // Maildev SMTP
		nil,              // No auth
		from,
		to,
		msg,
	)
	if err != nil {
		log.Fatalf("failed to send mail: %v", err)
	}
	fmt.Println("Mail sent successfully!")
}
