package InformationSystem

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	from := "iemre2003@gmail.com"
	password := "igqk nfbt zfbv gqom"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email with error: %v", err)
	}

	return nil
}
