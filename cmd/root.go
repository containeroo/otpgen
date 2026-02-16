package cmd

import (
	"errors"
	"fmt"
	"github.com/pquerna/otp/totp"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

var (
	timeNow      = time.Now
	generateCode = totp.GenerateCode
	rootCmd      = newRootCmd()
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "otpgen [secret]",
		Short:         "Generate one-time passwords",
		Long:          `otpgen is a command line tool to generate one-time passwords.`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runRoot,
	}

	cmd.AddCommand(newCompletionCmd(), newVersionCmd())

	return cmd
}

func runRoot(cmd *cobra.Command, args []string) error {
	secret := strings.TrimSpace(args[0])
	if secret == "" {
		return errors.New("secret must not be empty")
	}

	token, err := generateCode(secret, timeNow())
	if err != nil {
		return fmt.Errorf("generate TOTP: %w", err)
	}

	_, err = fmt.Fprintln(cmd.OutOrStdout(), token)
	return err
}

func Execute() error {
	return rootCmd.Execute()
}
