package main

import (
	"fmt"
	"io/ioutil"
)

func GenerateQRCode(code string) []byte {
	return nil
}

func main() {
	fmt.Println("Program for generating QR codes")

	qrcode := GenerateQRCode("vlad")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}