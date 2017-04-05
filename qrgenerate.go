package main

import qrcode "github.com/skip2/go-qrcode"

func main () {
	qrcode.WriteFile("kalym", qrcode.Medium, 256, "qr.png")
}