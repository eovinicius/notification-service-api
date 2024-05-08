package services

import (
	"github/eovinicius/notification/internal/entity"
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

	_ = godotenv.Load()

	FROM := os.Getenv("SMTP_FROM")
	PASSWORD := os.Getenv("SMTP_PASSWORD")
	HOST := os.Getenv("SMTP_HOST")

	auth := smtp.PlainAuth("", FROM, PASSWORD, HOST)

	message := entity.Mail{
		From:    FROM,
		To:      to,
		Subject: subject,
		Date:    time.Now(),
		Body:    body,
	}

	err := smtp.SendMail("", auth, FROM, []string{to}, message.ToByte())

	if err != nil {
		return err
	}

	return nil
}