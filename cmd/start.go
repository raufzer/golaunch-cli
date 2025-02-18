package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Initialize the CLI by creating assets and config file",
	Run:   startCommand,
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

	fmt.Println("CLI initialized successfully! Use 'golaunch setup' to add commands.")
}
func init() {
	rootCmd.AddCommand(startCmd)
}
