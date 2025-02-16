package commands

import (
	"fmt"
	"golaunch-cli/internal/config"
	"golaunch-cli/internal/launcher"
	"log"

	"github.com/urfave/cli/v2"
)

func OpenCommand() *cli.Command {
	return &cli.Command{
		Name:  "open",
		Usage: "Open programs for a specific command",
		Action: func(c *cli.Context) error {
			command := c.Args().First()
			if command == "" {
				return fmt.Errorf("please provide a command (e.g., dev, design)")
			}

			cfg, err := config.LoadConfig("assets/config.json")
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			programs, ok := cfg[command]
			if !ok {
				return fmt.Errorf("command '%s' not found in config", command)
			}

			for _, program := range programs {
				if err := launcher.Launch(program); err != nil {
					log.Printf("failed to launch %s: %v", program, err)
				} else {
					fmt.Printf("Launched %s successfully!\n", program)
				}
			}

			return nil
		},
	}
}
