package logger

import (
	"fmt"
	"os"
	"time"

)

var (
	logEnabled = false
	logFile    = "assets/command_logs.txt"
)

func EnableLogging() {
	logEnabled = true
}

func DisableLogging() {
	logEnabled = false
}

func LogCommand(command string) {
	if !logEnabled {
		return
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, command)

	if _, err := file.WriteString(logEntry); err != nil {
		fmt.Printf("Failed to write to log file: %v\n", err)
	}
}
