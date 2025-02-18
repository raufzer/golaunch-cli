package cmd

import (
	"golaunch-cli/internal/config"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all existing commands",
	Long:    `List all commands that have been set up in the configuration file.`,
	Run:     listCommand,
	GroupID: "core",
}

func listCommand(cmd *cobra.Command, args []string) {

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		cmd.Printf("Failed to load config: %v\n", err)
		return
	}

	if len(cfg) == 0 {
		cmd.Println("No commands have been set up yet.")
		return
	}

	cmd.Println("Setup commands:")
	for command := range cfg {
		cmd.Printf(" - %s\n", command)
	}
}

func init() {

	rootCmd.AddCommand(listCmd)
}
