package cmd

import (
	"golaunch-cli/internal/config"
	"golaunch-cli/internal/launcher"
	"log"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open <command>",
	Short: "Open programs for a specific command",
	Long:  `Open programs associated with a specific command by specifying the command name.`,
	Args:  cobra.ExactArgs(1),
	Run:   openCommand,
}

func openCommand(cmd *cobra.Command, args []string) {
	command := args[0]

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		cmd.Printf("Failed to load config: %v\n", err)
		return
	}

	programs, ok := cfg[command]
	if !ok {
		cmd.Printf("Command '%s' not found in config.\n", command)
		return
	}

	for _, program := range programs {
		if err := launcher.Launch(program); err != nil {
			log.Printf("Failed to launch %s: %v\n", program, err)
		} else {
			cmd.Printf("Launched %s successfully!\n", program)
		}
	}
}

func init() {

	rootCmd.AddCommand(openCmd)
}
