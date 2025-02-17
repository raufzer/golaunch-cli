package commands

import (
	"bufio"
	"fmt"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:      "add",
		Usage:     "Add a new program to an existing custom command",
		ArgsUsage: "<command>",
		Action:    addProgramToCommand,
	}
}

func addProgramToCommand(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("missing command name. Usage: golaunch add <command>")
	}

	command := c.Args().First()

	if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {
		return fmt.Errorf("config file not found. Run 'golaunch setup' to create a command.")
	}

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	if _, exists := cfg[command]; !exists {
		return fmt.Errorf("command '%s' does not exist", command)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the path to the program: ")
	program, _ := reader.ReadString('\n')
	program = strings.TrimSpace(program)

	cfg[command] = append(cfg[command], program)

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Program added to command '%s'.\n", command)
	return nil
}
