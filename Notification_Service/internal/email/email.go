package email1

import (
	"log"

	"gopkg.in/gomail.v2"
)

func sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "xihuannibao520@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "xihuannibao520@gmail.com", "yqgk jfzq klhx amqd")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	
	return nil
}

func Sent(email, body string) error {
	to := email
	subject := "New Notification"

	if err := sendEmail(to, subject, body); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
