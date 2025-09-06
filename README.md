# barcodemail
[![GoDoc](https://godoc.org/github.com/devzeeh/barcodemail?status.svg)](https://pkg.go.dev/github.com/devzeeh/barcodemail)
[![MITlicensed](https://img.shields.io/badge/License-MIT-blue.svg)](https://raw.githubusercontent.com/devzeeh/barcodemail/main/LICENSE)
[![GitHub Repo stars](https://img.shields.io/github/stars/devzeeh/barcodemail?style=social&label=Stars)](https://github.com/devzeeh/barcode)

Generate **QR codes** in Go, and send them via email.

> **Note:**  
> - This project is primarily for my own learning and personal projects.  
> - Intended for **small student projects, hobby apps, or testing purposes only**.  
> - Currently supports Only **Gmail SMTP** (other providers may require adjustments).  
> - Feel free to explore, use, or modify, but expect limited support.

---

## Features  
- Generate **QR codes**  
- Send QR codes via **email**  
- Built with **Go** for speed and simplicity

## Contributing & Feedback
Feel free to use this project, and don’t hesitate to [open an issue](https://github.com/devzeeh/barcodemail/issues) if you find a bug or want to suggest an improvement!

## Example
```go
package main

import (
	"fmt"

	"github.com/devzeeh/barcodemail"
	"github.com/devzeeh/barcodemail/email"
)

func main() {
	// Set sender details
	email.SenderEmail("sender@test.com") // Your email address
	email.SenderPassword("ABCD EF12 3456") // Use app password for Gmail
	email.SenderName("Devzeeh")

	data, file := barcodemail.Qrcode()
	contentID := email.QrEmbed("QRCODE") // Content-ID for the embedded image

	// HTML email body with embedded QR code
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

        // Prepare email data
	sendData := email.MailData{
		From: 	 "user@test.com",
		Subject: "Barcodemail",
		Body:    body,
		Qrcode:  file,
		Cid:     contentID,
	}

        // Send the email
	email.Mail(sendData)
}
```

## Requirements
- Make sure you have **Go version 1.22 or later** installed
- An email account with SMTP enabled (e.g., Gmail, Outlook, etc.). Requires App Password. 

## Usage Notes
- By default, this library saves generated QR codes as PNG files.
- Ensure your SMTP provider allows third-party app access (some may require app-specific passwords).
- Works best for **testing, student projects, and hobby apps**.

## License
This project is licensed under the [MIT License](https://github.com/devzeeh/barcodemail/blob/main/LICENSE).

## Acknowledgments
Inspired by packages such as
- [boombuler/barcode](https://github.com/boombuler/barcode)
- [go-gomail/gomail](https://github.com/go-gomail/gomail)
  
Thanks to the **Go** community for useful libraries, packages and tools
