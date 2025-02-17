package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const Version = "1.1.0"

func VersionCommand() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show the current version of golaunch",
		Action:  showVersion,
	}
}

func showVersion(c *cli.Context) error {
	fmt.Printf("golaunch version %s\n", Version)
	return nil
}
