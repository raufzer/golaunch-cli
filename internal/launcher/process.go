package launcher

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

)

func IsRunning(program string) (bool, error) {
	switch runtime.GOOS {
	case "windows":
		out, err := exec.Command("tasklist", "/FI", "IMAGENAME eq "+program).Output()
		if err != nil {
			return false, fmt.Errorf("failed to check if %s is running: %v", program, err)
		}
		return strings.Contains(string(out), program), nil
	case "darwin", "linux":
		out, err := exec.Command("pgrep", program).Output()
		if err != nil {

			if strings.Contains(err.Error(), "exit status 1") {
				return false, nil
			}
			return false, fmt.Errorf("failed to check if %s is running: %v", program, err)
		}
		return len(out) > 0, nil
	default:
		return false, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}

func StopProgram(program string) error {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("taskkill", "/IM", program, "/F")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to stop %s: %v", program, err)
		}
	case "darwin", "linux":
		cmd := exec.Command("pkill", program)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to stop %s: %v", program, err)
		}
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	return nil
}
