package smtputil

import (
	"fmt"

	"github.com/wildanpurnomo/nsq-ayayaclap/smtp-service/libs"
	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "127.0.0.1"
const CONFIG_SMTP_PORT = 1025
const CONFIG_SENDER_NAME = "SMTP Service <smtpservice@example.com>"

func SendRegistrationConfirmationMail(toEmail string) error {
	token, err := libs.GenerateEmailConfirmationToken(toEmail)
	if err != nil {
		return err
	}

	emailBody := fmt.Sprintf(`<a href="http://localhost:8080/api/user/email-confirmation?redir_token=%v" target="_blank">Confirm your registration</a>`, token)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Test Registration Confirmation Email")
	mailer.SetBody("text/html", emailBody)

	dialer := &gomail.Dialer{Host: CONFIG_SMTP_HOST, Port: CONFIG_SMTP_PORT}

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
