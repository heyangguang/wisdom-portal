package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)


func main() {
	err := qrcode.WriteFile("http://blog.csdn.net/wangshubo1989", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("sad")
	}
}
