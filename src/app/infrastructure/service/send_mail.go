package service

import (
	"board/config"
	"fmt"
	"net/smtp"
	"strings"
)

var (
	username = "user@example.com"
	password = "password"
)

func SendEmail(
	from string,
	recipients []string,
	subject string,
	body string,
) error {
	config := config.GetConfig()
	smtpServer := fmt.Sprintf("%s:%d", config.Mail.Host, config.Mail.Port)
	auth := smtp.CRAMMD5Auth(username, password)
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body))

	if err := smtp.SendMail(smtpServer, auth, from, recipients, msg); err != nil {
		return err
	}
	return nil
}
