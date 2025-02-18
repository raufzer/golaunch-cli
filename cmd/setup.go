package cmd

import (
	"bufio"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/spf13/cobra"

)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Set up custom commands and program paths",
	Run:   setupCommand,
	GroupID: "setup",
}

func setupCommand(cmd *cobra.Command, args []string) {
	if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {
		cmd.Println("Config file not found. Run 'golaunch start' to initialize the CLI.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	cmd.Print("Enter a custom command (e.g., dev, design): ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

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

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		cmd.Printf("Failed to load config: %v\n", err)
		return
	}

	cfg[command] = programs

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		cmd.Printf("Failed to save config: %v\n", err)
		return
	}

	cmd.Printf("Setup complete! Use 'golaunch open %s' to launch your programs.\n", command)
}

func init() {
	// Add the setup command to the root command
	rootCmd.AddCommand(setupCmd)
}
