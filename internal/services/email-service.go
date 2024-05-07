package services

import (
	"net/smtp"
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

	from := "john.doe@example.com"

	auth := smtp.PlainAuth("", " ", " ", "")

	err := smtp.SendMail("", auth, from, []string{to}, []byte(body))

	if err != nil {
		return err
	}

	return nil
}
