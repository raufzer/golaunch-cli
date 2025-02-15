package cmd

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
		Usage: "Open apps for a specific context",
		Action: func(c *cli.Context) error {
			context := c.Args().First()
			if context == "" {
				return fmt.Errorf("please provide a context (e.g., dev, design)")
			}

			cfg, err := config.LoadConfig("assets/config.json")
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			apps, ok := cfg[context]
			if !ok {
				return fmt.Errorf("context '%s' not found in config", context)
			}

			for _, app := range apps {
				if err := launcher.Launch(app); err != nil {
					log.Printf("failed to launch %s: %v", app, err)
				} else {
					fmt.Printf("Launched %s successfully!\n", app)
				}
			}

			return nil
		},
	}
}
