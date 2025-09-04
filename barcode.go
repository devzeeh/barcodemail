// Package barcodemail provides functionality to generate barcodes and QR codes.
// It includes methods to create barcodes, QR codes, and send emails with embedded images.
package barcodemail

import (
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// Random number generator for generating unique barcodes and QR codes.
// It uses a seeded random number generator to ensure different results on each run.
// The random numbers are always 4 digits long (1000-9999).
// This is useful for generating unique identifiers for barcodes and QR codes.
var (
	vrand         = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomQR      = vrand.Intn(9000) + 1000// Always generate a 4-digit random number (1000-9999)
)

// Filepath creates the file and returns the file pointer
func filepath(filename string) *os.File {
	fileBarcode, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return fileBarcode
}

// Barcode generates a barcode image and returns the data and file path.
// Qrcode generates a QR code with a random 4-digit number.
// Example: "QRcode1234"
func Qrcode() (string, string) {
	qrData := "QRcode" + strconv.Itoa(randomQR)

	qrcode, err := qr.Encode(qrData, qr.M, qr.Auto)
	if err != nil {
		panic(err)
	}

	qrcode, err = barcode.Scale(qrcode, 256, 256)
	if err != nil {
		panic(err)
	}

	filename := "qrcode.png"
	file := filepath(filename)
	defer file.Close()

	if err = png.Encode(file, qrcode); err != nil {
		panic(err)
	}

	return qrData, filename
}
