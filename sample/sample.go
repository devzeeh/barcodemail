// Sample for version v1.0.1 of barcodemail
package main

import (
	"fmt"

	"github.com/devzeeh/barcodemail"
	"github.com/devzeeh/barcodemail/email"
)

func main() {
	//
	email.SenderEmail("sender@test.com")
	email.SenderPassword("ABCD EF12 3456")
	email.SenderName("Devzeeh")

	data, file := barcodemail.Qrcode()
	contentID := email.QrEmbed("QRCODE") // Content-ID for the embedded image

	//
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif;">
			<h2>Hello</h2>
			<p>Here are your codes: %s</p>
			<img src="cid:%s" alt="QR Code" style="border:1px solid #ddd; padding:5px;">
			<p>Thank you,<br></p>
		</body>
		</html>
	`, data, contentID)

	sendData := email.MailData{
		From: "user@test.com",
		Subject: "Barcodemail",
		Body:    body,
		Qrcode:  file,
		Cid:     contentID,
	}

	email.Mail(sendData)
}
