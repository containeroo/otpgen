package cmd

import (
	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion [bash|zsh|powershell]",
		Short:                 "Generate completion scripts",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		RunE:                  runCompletion,
	}
}

func runCompletion(cmd *cobra.Command, args []string) error {
	switch args[0] {
	case "bash":
		return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
	case "zsh":
		return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
	case "powershell":
		return cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
	default:
		return nil
	}
}
