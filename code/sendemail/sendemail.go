package sendemail

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

const (
	emailFrom = "golangemail@gmail.com"
	password  = "M'7#uX~4k2JJQw#9"
	smtpHost  = "smtp.gmail.com"
	smtpPort  = 587
	emailTo   = "dima220v@gmail.com"
)

var (
	m = gomail.NewMessage()
	d = gomail.NewDialer("smtp.gmail.com", 587, emailFrom, password)
)

func gomailInit() {
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Gomail test subject")
	m.SetBody("text/plain", "Gomail test subject")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func SendEmail(subject, body string) {
	gomailInit()
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// func main() {
// 	gomailInit()
// 	// To send Email call SendEmail
// 	// SendEmail("Gomail test subject", "Gomail test subject")
// }
