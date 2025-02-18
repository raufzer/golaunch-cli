package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd is the base command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "golaunch",
	Short: "Launch your favorite apps with a single command",
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
