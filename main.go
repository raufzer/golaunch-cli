package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"golaunch-cli/cmd"
)

func main() {
	app := &cli.App{
		Name:  "golaunch",
		Usage: "Launch your favorite apps with a single command",
		Commands: []*cli.Command{
			cmd.OpenCommand(),
			cmd.ConfigCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
