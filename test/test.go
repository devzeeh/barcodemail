package main

import (
	"fmt"
	"os"

	"github.com/devzeeh/barcodemail"
	"github.com/devzeeh/barcodemail/email"
)

func main() {
	// load an env

	email.SenderEmail(os.Getenv("AUTH_EMAIL"))
	email.SenderPassword(os.Getenv("AUTH_PASSWORD"))
	email.SenderName("Barcode")

	data, file := barcodemail.Qrcode()
	contentID := email.QrEmbed("CODE")

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
		ToEmail: "user@test.com",
		Subject: "Barcodemail",
		Body:    body,
		Qrcode:  file,
		Cid:     contentID,
	}

	email.Mail(sendData)
}
