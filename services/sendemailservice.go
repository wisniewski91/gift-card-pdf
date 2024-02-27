package services

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/go-mail/mail"

	"github.com/wisniewski91/gift-card-pdf/types"
)

func SendEmail(data types.GiftCard, filename string) error {

	host := "smtp.gmail.com"
	email := ""
	password := ""
	port := 587

	template, _ := template.ParseFiles("./views/email/email.html")

	var body bytes.Buffer

	template.Execute(&body, struct {
		Name    string
		Service string
		Price   string
	}{
		Name:    "Puneet Singh",
		Service: "This is a test message in a HTML template",
		Price:   "200 zł",
	})

	m := mail.NewMessage()

	m.SetHeader("From", email)

	m.SetHeader("To", data.Email)

	m.SetHeader("Subject", "Twój bon upominkowy")

	m.SetBody("text/html", body.String())

	m.Attach(fmt.Sprintf("./static/giftcards/%s", filename))

	d := mail.NewDialer(host, port, email, password)

	// Send the email to Kate, Noah and Oliver.

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	fmt.Println("Email Sent")
	return nil
}
