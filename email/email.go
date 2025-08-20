// Package email provides functionality to send HTML emails with embedded QR codes.
// It includes methods to set sender credentials, create email content, and send emails using the built-in gomail package.
package email

import (
	"fmt"
	"log"
	"strings"

	"gopkg.in/gomail.v2"
)

var (
	senderEmail    string
	senderPassword string
	senderName     string
)

// SenderEmail sets the sender's email address.
// This should be used with caution and not hard-coded in production code.
// Consider using environment variables or secure vaults for sensitive information.
func SenderEmail(email string) {
	senderEmail = strings.TrimSpace(email)
}

// SenderPassword sets the sender's email password.
// This should be used with caution and not hard-coded in production code.
// Consider using environment variables or secure vaults for sensitive information.
func SenderPassword(pw string) {
	senderPassword = strings.TrimSpace(pw)
}

// SenderName sets the sender's name.
func SenderName(name string) {
	senderName = strings.TrimSpace(name)
}

// MailData holds the data required to send an email.
// It includes the recipient's email, subject, body, QR code file path, and content ID for embedding images.
type MailData struct {
	ToEmail string // Recipient's email address
	Subject string // Subject of the email
	Body    string // HTML body of the email
	Qrcode  string // File path of the QR code image
	Cid     string // Content ID for the embedded image
}

// Mail sends an HTML email with an embedded QR code image.
// It uses the gomail package to create and send the email.
func Mail(Mail MailData) {
	// sanity check

	// SMTP Config
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	// Create email
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", senderName, senderEmail))
	m.SetHeader("To", Mail.ToEmail)
	m.SetHeader("Subject", Mail.Subject)
	m.SetBody("text/html", Mail.Body)

	// Embed QR code images
	m.Embed(Mail.Qrcode, gomail.SetHeader(map[string][]string{
		"Content-ID": {QrEmbed(Mail.Cid)},
	}))

	// Send email
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

	fmt.Println("HTML email with barcode and QR code sent successfully!")
}

// QrEmbed returns the content ID for embedding QR codes in emails.
// This is a placeholder function that can be customized as needed.
// It is used to generate a content ID for the QR code image in the email.
func QrEmbed(content string) string {
	return content
}
