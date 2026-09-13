// Command otpgen generates time-based one-time passwords.
package main

import (
	"fmt"
	"os"

	"github.com/containeroo/otpgen/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
