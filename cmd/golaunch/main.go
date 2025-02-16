package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golaunch-cli/cmd/commands"
)

func main() {
	app := &cli.App{
		Name:  "golaunch",
		Usage: "Launch your favorite apps with a single command",
		Commands: []*cli.Command{
			commands.StartCommand(),
			commands.SetupCommand(),
			commands.OpenCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
