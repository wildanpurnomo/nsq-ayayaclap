package smtputil

import (
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "127.0.0.1"
const CONFIG_SMTP_PORT = 1025
const CONFIG_SENDER_NAME = "SMTP Service <smtpservice@example.com>"

func SendRegistrationConfirmationMail(toEmail string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Test Registration Confirmation Email")
	mailer.SetBody("text/plain", "Hello")

	dialer := &gomail.Dialer{Host: CONFIG_SMTP_HOST, Port: CONFIG_SMTP_PORT}

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
