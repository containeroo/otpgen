package main

import (
	"flag"
	"fmt"
	"github.com/xlzd/gotp"
	"log"
)

func main() {
	secretKey := flag.String("secretKey", "", "supply a valid TOTP secret key to generate a token from")
	flag.Parse()
	if *secretKey == "" {
		log.Fatal("secretKey cannot be empty")
	}
	fmt.Print(gotp.NewDefaultTOTP(*secretKey).Now())
}
