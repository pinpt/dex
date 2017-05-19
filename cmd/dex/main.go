package dex

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "dex",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}
	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())
	return rootCmd
}
