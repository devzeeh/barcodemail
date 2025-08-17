# barcodemail

Generate **QR codes** in Go, and send them via email.

> **Note:** This is primarily for my own project. Feel free to explore, use, or modify, but expect limited support.

---

## Features  
- Generate **QR codes**  
- Send QR codes via **email**  
- Built with **Go** for speed and simplicity

## Contributing & Feedback
Feel free to use this project, and don’t hesitate to [open an issue](https://github.com/devzeeh/barcodemail/issues) if you find a bug or want to suggest an improvement!

## Example
```
package main

import (
	"fmt"

	"github.com/devzeeh/barcodemail"
	"github.com/devzeeh/barcodemail/email"
)

func main() {
	// sender credential
	email.SenderEmail("Sender Email")
	email.SenderPassword("Sender Password")
	email.SenderName("Sender Name")

	qrData, filePath := barcodemail.Qrcode() // Generate QR code and get data and file path
	cid := email.QrEmbed("QRCODE")           // Content-ID for the embedded image

	// Email content
	htmlBody := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif;">
			<h2>Hello</h2>
			<p>Here are your codes: %s</p>
			<img src="cid:%s" alt="QR Code" style="border:1px solid #ddd; padding:5px;">
			<p>Thank you,<br></p>
		</body>
		</html>
	`, qrData, cid)

	sendData := email.MailData{
		ToEmail: "Receiver Email",
		Subject: "Subject of the email",
		Body:    htmlBody,
		Qrcode:  filePath,
		Cid:     cid,
	}

	email.Mail(sendData)
}

```

## Requirements
- Make sure you have **Go version 1.22 or later** installed
- An email account with SMTP enabled (e.g., Gmail, Outlook, etc.)

## Usage Notes
- By default, this library saves generated QR codes as PNG files.
- Ensure your SMTP provider allows third-party app access (some may require app-specific passwords).
- Best suited for testing, small projects

## License
This project is licensed under the [MIT License](https://github.com/devzeeh/barcodemail/blob/main/LICENSE).

## Acknowledgments
Inspired by packages such as
- [boombuler/barcode](https://github.com/boombuler/barcode)
- [go-gomail/gomail](https://github.com/go-gomail/gomail)
  
Thanks to the **Go** community for useful libraries, packages and tools
