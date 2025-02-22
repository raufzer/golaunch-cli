package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "2.1.0"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Show the current version of golaunch",
	Run:     versionCommand,
	GroupID: "misc",
}

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("golaunch version %s\n", Version)
}
func init() {
	rootCmd.AddCommand(versionCmd)
}
