package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "3.0.0"

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print the version number of otpgen",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
