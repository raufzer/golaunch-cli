package cmd

import (
	"github.com/spf13/cobra"

)

var rootCmd = &cobra.Command{
	Use:   "golaunch",
	Short: "Launch your favorite apps with a single command",
}

func Execute() error {
	return rootCmd.Execute()
}
