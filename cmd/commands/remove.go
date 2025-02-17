package commands

import (
	"fmt"
	"golaunch-cli/internal/config"
	"os"

	"github.com/urfave/cli/v2"
)

func RemoveCommand() *cli.Command {
	return &cli.Command{
		Name:      "remove",
		Usage:     "Remove a program from a specific command",
		ArgsUsage: "<command>",
		Action:    removeProgramFromCommand,
	}
}

func removeProgramFromCommand(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("missing command name. Usage: golaunch remove <command>")
	}

	command := c.Args().First()

	if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {
		return fmt.Errorf("config file not found. Run 'golaunch setup' to create a command.")
	}

	cfg, err := config.LoadConfig("assets/config.json")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	programs, exists := cfg[command]
	if !exists {
		return fmt.Errorf("command '%s' does not exist", command)
	}

	fmt.Printf("Programs for command '%s':\n", command)
	for i, program := range programs {
		fmt.Printf("%d: %s\n", i+1, program)
	}

	var index int
	fmt.Print("Enter the number of the program to remove: ")
	_, err = fmt.Scan(&index)
	if err != nil || index < 1 || index > len(programs) {
		return fmt.Errorf("invalid program number")
	}

	cfg[command] = append(programs[:index-1], programs[index:]...)

	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Program removed from command '%s'.\n", command)
	return nil
}
