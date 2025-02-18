package cmd

import (
	"bufio"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <command>",
	Short: "Add a new program to an existing custom command",
	Long:  `Add a new program to an existing command by specifying the command name and the program path.`,
	Args:  cobra.ExactArgs(1),
	Run:   addProgramToCommand,
}

func addProgramToCommand(cmd *cobra.Command, args []string) {
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

	cmd.Print("Enter the path to the program: ")
	program, _ := reader.ReadString('\n')
	program = strings.TrimSpace(program)

	cfg[command] = append(cfg[command], program)

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		cmd.Printf("Failed to save config: %v\n", err)
		return
	}

	cmd.Printf("Program added to command '%s'.\n", command)
}

func init() {

	rootCmd.AddCommand(addCmd)
}
