package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/xlzd/gotp"
)

const version = "2.0.1"

func main() {
	secretKey := flag.StringP("secretkey", "s", "", "supply a valid TOTP secret key to generate a token from")
	printVersion := flag.BoolP("version", "v", false, "Print the current version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *secretKey == "" {
		log.Fatal("secretkey cannot be empty")
	}

	if len(*secretKey) != 16 && len(*secretKey) != 32 {
		log.Fatal("secretkey must be 16 or 32 characters long")
	}

	fmt.Print(gotp.NewDefaultTOTP(*secretKey).Now())
}
