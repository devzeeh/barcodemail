module github.com/devzeeh/barcodemail

go 1.25.0

retract (
	v1.1.0-beta.1 // published by mistake
	v1.0.0 // published by mistake
)

require (
	github.com/boombuler/barcode v1.1.0
	github.com/joho/godotenv v1.5.1
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
