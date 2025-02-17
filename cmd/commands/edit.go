package commands

import (
	"bufio"
	"fmt"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func EditCommand() *cli.Command {
	return &cli.Command{
		Name:      "edit",
		Usage:     "Edit the command configuration, e.g., change program paths",
		ArgsUsage: "<command>",
		Action:    editCommand,
	}
}

func editCommand(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("missing command name. Usage: golaunch edit <command>")
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

	
	var programs []string
	for {
		fmt.Print("Enter the path to a program (or press Enter to finish): ")
		program, _ := reader.ReadString('\n')
		program = strings.TrimSpace(program)

		if program == "" {
			break
		}

		programs = append(programs, program)
	}

	
	cfg[command] = programs

	
	if err := config.SaveConfig("assets/config.json", cfg); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	fmt.Printf("Command '%s' updated successfully.\n", command)
	return nil
}