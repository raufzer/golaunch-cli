package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Initialize the CLI by creating assets and config file",
	Run:     startCommand,
	GroupID: "setup",
}

func startCommand(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("assets"); os.IsNotExist(err) {
		if err := os.Mkdir("assets", 0755); err != nil {
			fmt.Printf("Failed to create assets folder: %v\n", err)
			return
		}
		fmt.Println("Created assets folder.")
	}

	if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {
		if err := config.SaveConfig("assets/config.json", map[string][]string{}); err != nil {
			fmt.Printf("Failed to create config file: %v\n", err)
			return
		}
		fmt.Println("Created config file.")
	}

	if _, err := os.Stat("assets/command_logs.txt"); os.IsNotExist(err) {
		if _, err := os.Create("assets/command_logs.txt"); err != nil {
			fmt.Printf("Failed to create log file: %v\n", err)
			return
		}
		fmt.Println("Created log file.")
	}

	if _, err := os.Stat("assets/command_history.json"); os.IsNotExist(err) {
		if err := os.WriteFile("assets/command_history.json", []byte("[]"), 0644); err != nil {
			fmt.Printf("Failed to create history file: %v\n", err)
			return
		}
		fmt.Println("Created history file.")
	}

	fmt.Println("CLI initialized successfully! Use 'golaunch setup' to add commands.")
}

func init() {
	rootCmd.AddCommand(startCmd)
}
