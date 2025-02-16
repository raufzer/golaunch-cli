package launcher

import (
	"fmt"
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

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to launch %s: %v", app, err)
	}
	return nil
}
