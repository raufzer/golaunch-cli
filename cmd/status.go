package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"
	"golaunch-cli/internal/launcher"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <command>",
	Short: "Check if the applications linked to a command are running",
	Args:  cobra.ExactArgs(1),
	Run:   statusCommand,
}

func statusCommand(cmd *cobra.Command, args []string) {
	command := args[0]

	// Load config to get programs linked to the command
	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		return
	}

	programs, exists := cfg[command]
	if !exists {
		fmt.Printf("Command '%s' not found in config.\n", command)
		return
	}

	// Check if each program is running
	for _, program := range programs {
		running, err := launcher.IsRunning(program)
		if err != nil {
			fmt.Printf("Error checking status of %s: %v\n", program, err)
			continue
		}
		if running {
			fmt.Printf("%s is running.\n", program)
		} else {
			fmt.Printf("%s is not running.\n", program)
		}
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
