package cmd

import (
	"fmt"
	"golaunch-cli/internal/history"

	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the last used commands and timestamps",
	Run:   historyCommand,
}

func historyCommand(cmd *cobra.Command, args []string) {
	entries, err := history.GetHistory()
	if err != nil {
		fmt.Printf("Failed to retrieve history: %v\n", err)
		return
	}

	for _, entry := range entries {
		fmt.Printf("[%s] %s\n", entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Command)
	}
}

func init() {
	rootCmd.AddCommand(historyCmd)
}
