package commands

import (
	"fmt"
	"golaunch-cli/internal/config"
	"os"

	"github.com/urfave/cli/v2"
)

func StartCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Initialize the CLI by creating the assets folder and config file",
		Action: func(c *cli.Context) error {

			if _, err := os.Stat("assets"); os.IsNotExist(err) {

				if err := os.Mkdir("assets", 0755); err != nil {
					return fmt.Errorf("failed to create assets folder: %v", err)
				}
				fmt.Println("Created assets folder.")
			}

			if _, err := os.Stat("assets/config.json"); os.IsNotExist(err) {

				if err := config.SaveConfig("assets/config.json", map[string][]string{}); err != nil {
					return fmt.Errorf("failed to create config file: %v", err)
				}
				fmt.Println("Created config file.")
			}

			fmt.Println("CLI initialized successfully! Use 'golaunch setup' to add commands and program paths.")
			return nil
		},
	}
}
