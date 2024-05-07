package services

import (
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type EmailSender interface {
	SendEmail(to string, subject string, body string) error
}

type emailService struct {
}

func NewEmailService() *emailService {
	return &emailService{}
}

func (es *emailService) SendEmail(to string, subject string, body string) error {

	_ = godotenv.Load(".env")

	FROM := os.Getenv("SMTP_FROM")
	PASSWORD := os.Getenv("SMTP_PASSWORD")
	HOST := os.Getenv("SMTP_HOST")

	auth := smtp.PlainAuth("", FROM, PASSWORD, HOST)

	message := mail{
		From:    FROM,
		To:      to,
		Subject: subject,
		Date:    time.Now(),
		Body:    body,
	}

	err := smtp.SendMail("", auth, FROM, []string{to}, message.toByte())

	if err != nil {
		return err
	}

	return nil
}

type mail struct {
	From    string
	To      string
	Subject string
	Date    time.Time
	Body    string
}

func (m mail) toByte() []byte {
	return []byte("From: " + m.From + "\r\n" +
		"To: " + m.To + "\r\n" +
		"Subject: " + m.Subject + "\r\n" +
		"Date: " + m.Date.String() + "\r\n" +
		"\r\n" +
		m.Body)
}
