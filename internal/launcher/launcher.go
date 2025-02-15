package launcher

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Launch opens an app based on the OS
func Launch(app string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", app) // Windows: start the app
	case "darwin": // macOS
		cmd = exec.Command("open", "-a", app) // macOS: open the app
	case "linux":
		cmd = exec.Command(app) // Linux: run the app directly
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to launch %s: %v", app, err)
	}

	fmt.Printf("Launched %s successfully!\n", app)
	return nil
}
