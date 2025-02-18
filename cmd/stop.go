package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"
	"golaunch-cli/internal/launcher"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop <command>",
	Short: "Close all applications linked to a command",
	Args:  cobra.ExactArgs(1),
	Run:   stopCommand,
}

func stopCommand(cmd *cobra.Command, args []string) {
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

	// Stop each program
	for _, program := range programs {
		if err := launcher.StopProgram(program); err != nil {
			fmt.Printf("Error stopping %s: %v\n", program, err)
		} else {
			fmt.Printf("Stopped %s.\n", program)
		}
	}
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
