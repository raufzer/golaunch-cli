package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"
	"os"

	"github.com/spf13/cobra"

)

var removeCmd = &cobra.Command{
	Use:   "remove <command>",
	Short: "Remove a program from a specific command",
	Long:  `Remove a program from a specific command by specifying the command name and selecting the program to remove.`,
	Args:  cobra.ExactArgs(1),
	Run:   removeProgramFromCommand,
}

func removeProgramFromCommand(cmd *cobra.Command, args []string) {
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

	programs, exists := cfg[command]
	if !exists {
		cmd.Printf("Command '%s' does not exist.\n", command)
		return
	}

	cmd.Printf("Programs for command '%s':\n", command)
	for i, program := range programs {
		cmd.Printf("%d: %s\n", i+1, program)
	}

	var index int
	cmd.Print("Enter the number of the program to remove: ")
	_, err = fmt.Scan(&index)
	if err != nil || index < 1 || index > len(programs) {
		cmd.Println("Invalid program number.")
		return
	}

	cfg[command] = append(programs[:index-1], programs[index:]...)

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		cmd.Printf("Failed to save config: %v\n", err)
		return
	}

	cmd.Printf("Program removed from command '%s'.\n", command)
}

func init() {

	rootCmd.AddCommand(removeCmd)
}
