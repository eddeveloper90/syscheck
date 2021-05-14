package mailer

import (
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
	"time"
)

// Never fails, tries to format the address if possible
func formatEmailAddress(addr string) string {
	e, err := mail.ParseAddress(addr)
	if err != nil {
		return addr
	}
	return e.String()
}

func encodeRFC2047(str string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{Address: str}
	return strings.Trim(addr.String(), " <>")
}

func composeMimeMail(to string, from string, subject string, body string) []byte {
	t := time.Now()
	header := make(map[string]string)
	header["From"] = formatEmailAddress(from)
	header["To"] = formatEmailAddress(to)
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"
	//header["Date"] = "Wed, 3 Feb 2021 15:42:38 +0330"
	header["Date"] = t.Format("Mon, 2 Jan 2006 15:04:05 ") + "+0330"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	return []byte(message)
}

func main() {
	// Sender data.
	from := "report@safarmarket.com"
	password := "rprprp1361368"

	// Receiver email address.
	to := []string{
		"milad.golfam@safarmarket.com",
	}

	// smtp server configuration.
	smtpHost := "mail.safarmarket.com"
	smtpPort := "587"

	subject := "sbj: golang Mailer"

	// Message.
	//message := []byte("golang mailer.")

	toSingleEmail := "milad.golfam@gmail.com"

	msg := composeMimeMail(toSingleEmail, from, subject, "golang mailer.")

	// Authentication.
	auth := smtp.PlainAuth(from, from, password, smtpHost)

	// Sending email.
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
