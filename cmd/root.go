package cmd

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "otpgen [secret]",
	Short: "Generate one-time passwords",
	Long:  `otpgen is a command line tool to generate one-time passwords.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token, err := totp.GenerateCode(args[0], time.Now())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(token)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd, versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
