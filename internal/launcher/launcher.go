package launcher

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Launch(app string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", app)
	case "darwin":
		cmd = exec.Command("open", "-a", app)
	case "linux":
		cmd = exec.Command(app)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	return cmd.Run()
}
