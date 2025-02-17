package commands

import (
	"fmt"
	"golaunch-cli/internal/config"

	"github.com/urfave/cli/v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "List all existing commands",
		Action: func(c *cli.Context) error {

			cfg, err := config.LoadConfig("assets/config.json")
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			if len(cfg) == 0 {
				fmt.Println("No commands have been set up yet.")
				return nil
			}

			fmt.Println("Setup commands:")
			for command := range cfg {
				fmt.Printf(" - %s\n", command)
			}

			return nil
		},
	}
}
