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

func init() {
	// Define groups for commands
	rootCmd.AddGroup(&cobra.Group{
		ID:    "setup",
		Title: "Setup & Initialization",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "core",
		Title: "Core Functionality",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "state",
		Title: "Program State",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "history",
		Title: "History & Logging",
	})
	rootCmd.AddGroup(&cobra.Group{
		ID:    "misc",
		Title: "Miscellaneous",
	})
}
