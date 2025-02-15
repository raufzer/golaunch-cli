package cmd

import (
	"fmt"
	"golaunch-cli/internal/config"

	"github.com/urfave/cli/v2"
)

func ConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "Configure apps for a specific context",
		Action: func(c *cli.Context) error {
			context := c.Args().First()
			apps := c.Args().Tail()

			if context == "" || len(apps) == 0 {
				return fmt.Errorf("usage: golaunch config <context> <app1> <app2> ...")
			}

			cfg, err := config.LoadConfig("assets/config.json")
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			cfg[context] = apps

			if err := config.SaveConfig("assets/config.json", cfg); err != nil {
				return fmt.Errorf("failed to save config: %v", err)
			}

			fmt.Printf("Updated config for context '%s': %v\n", context, apps)
			return nil
		},
	}
}
