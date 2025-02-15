package cmd

import (
	"bufio"
	"fmt"
	"golaunch-cli/internal/config"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func SetupCommand() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Set up custom commands and program paths",
		Action: func(c *cli.Context) error {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter a custom command (e.g., dev, design): ")
			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command)

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

			cfg, err := config.LoadConfig("assets/config.json")
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			cfg[command] = programs

			if err := config.SaveConfig("assets/config.json", cfg); err != nil {
				return fmt.Errorf("failed to save config: %v", err)
			}

			fmt.Printf("Setup complete! Use 'golaunch open %s' to launch your programs.\n", command)
			return nil
		},
	}
}
