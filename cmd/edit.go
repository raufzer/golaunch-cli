package cmd

import (
	"bufio"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/spf13/cobra"

)

var editCmd = &cobra.Command{
	Use:   "edit <command>",
	Short: "Edit the command configuration, e.g., change program paths",
	Long:  `Edit the program paths associated with a specific command.`,
	Args:  cobra.ExactArgs(1),
	Run:   editCommand,
}

func editCommand(cmd *cobra.Command, args []string) {
	command := args[0]

	if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {
		cmd.Println("Config file not found. Run 'golaunch setup' to create a command.")
		return
	}

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		cmd.Printf("Failed to load config: %v\n", err)
		return
	}

	if _, exists := cfg[command]; !exists {
		cmd.Printf("Command '%s' does not exist.\n", command)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var programs []string

	for {
		cmd.Print("Enter the path to a program (or press Enter to finish): ")
		program, _ := reader.ReadString('\n')
		program = strings.TrimSpace(program)

		if program == "" {
			break
		}

		programs = append(programs, program)
	}

	cfg[command] = programs

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		cmd.Printf("Failed to save config: %v\n", err)
		return
	}

	cmd.Printf("Command '%s' updated successfully.\n", command)
}

func init() {

	rootCmd.AddCommand(editCmd)
}
