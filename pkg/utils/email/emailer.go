package email

import (
	"fmt"
	"net/smtp"
	"orderly/internal/infrastructure/config"
)

var (
	from string
	auth smtp.Auth
	addr string
)

func init() {
	from = config.Configs.Emailing.FromEmail
	auth = smtp.PlainAuth("", from, config.Configs.Emailing.AppPassword, config.Configs.Emailing.SmtpServerAddress)
	addr = fmt.Sprintf("%s:%s", config.Configs.Emailing.SmtpServerAddress, config.Configs.Emailing.SmtpsPort)
}

func SendCredentials(to, username, password string) error {
	if config.Configs.Dev_LogCredentials {
		fmt.Println("Email: ", to)
		fmt.Println("Username: ", username)
		fmt.Println("Password: ", password)
		return nil
	}

	subject := "Your credentials for Orderly"
	body := fmt.Sprintf(
		`Hello,

Here are your credentials for Orderly:

Username: %s
Password: %s

Please keep them safe.

Regards,
Orderly Team`,
		username, password)
	return SendEmail(to, subject, body)
}

func SendEmail(to, subject, body string) error {
	if !config.Configs.Dev_AllowSendingEmails {
		return nil
	}

	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body))

	err := smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
