package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {

	var qrWeb string
	var qrName string

	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("To generate your QR Code, please paste your website link below:")
		fmt.Println("Correct format : https://www.example.com/")
		scanner.Scan()
		qrWeb = scanner.Text()
		fmt.Println("Choose a name for your QR Code file")
		scanner.Scan()
		qrName = scanner.Text() + ".png"
		err := qrcode.WriteFile(qrWeb, qrcode.Medium, 256, qrName)

		if err != nil {
			fmt.Println("Error generating QR code:", err)
		} else {
			fmt.Println("QR code generated and saved as ''"+qrName, "''")
		}
	}
}
