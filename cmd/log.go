package cmd

import (
	"fmt"
	"golaunch-cli/internal/logger"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Toggle logging of executed commands",
}

var enableLogCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable command logging",
	Run:   enableLogCommand,
}

var disableLogCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable command logging",
	Run:   disableLogCommand,
}

func enableLogCommand(cmd *cobra.Command, args []string) {
	logger.EnableLogging()
	fmt.Println("Command logging enabled.")
}

func disableLogCommand(cmd *cobra.Command, args []string) {
	logger.DisableLogging()
	fmt.Println("Command logging disabled.")
}

func init() {
	logCmd.AddCommand(enableLogCmd)
	logCmd.AddCommand(disableLogCmd)
	rootCmd.AddCommand(logCmd)
}
