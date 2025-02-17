package main

import (
	"log"
	"os"

	"golaunch-cli/cmd/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "golaunch",
		Usage: "Launch your favorite apps with a single command",
		Commands: []*cli.Command{
			commands.StartCommand(),
			commands.SetupCommand(),
			commands.OpenCommand(),
			commands.ListCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
