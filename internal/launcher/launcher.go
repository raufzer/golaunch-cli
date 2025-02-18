package launcher

import (
	"fmt"
	"golaunch-cli/internal/history"
	"golaunch-cli/internal/logger"
	"os/exec"
	"runtime"
	"strings"
)

func Launch(app string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", strings.TrimSpace(app))
	case "darwin":
		cmd = exec.Command("open", "-a", app)
	case "linux":
		cmd = exec.Command(app)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// Log the command
	logger.LogCommand(app)

	// Add the command to history
	if err := history.AddCommand(app); err != nil {
		return fmt.Errorf("failed to add command to history: %v", err)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to launch %s: %v", app, err)
	}
	return nil
}
