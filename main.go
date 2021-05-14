package main

import (
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/xlzd/gotp"
)

func main() {
	secretKey := flag.StringP("secretkey", "s", "", "supply a valid TOTP secret key to generate a token from")
	flag.Parse()
	if *secretKey == "" {
		log.Fatal("secretkey cannot be empty")
	}
	fmt.Print(gotp.NewDefaultTOTP(*secretKey).Now())
}
