package history

import (
	"encoding/json"
	"golaunch-cli/internal/models"
	"os"
	"time"
)

var historyFile = "assets/command_history.json"

func AddCommand(command string) error {
	history, err := loadHistory()
	if err != nil {
		return err
	}

	entry := models.CommandEntry{
		Command:   command,
		Timestamp: time.Now(),
	}
	history = append(history, entry)

	return saveHistory(history)
}

func GetHistory() ([]models.CommandEntry, error) {
	return loadHistory()
}

func loadHistory() ([]models.CommandEntry, error) {
	file, err := os.ReadFile(historyFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.CommandEntry{}, nil
		}
		return nil, err
	}

	var history []models.CommandEntry
	if err := json.Unmarshal(file, &history); err != nil {
		return nil, err
	}

	return history, nil
}

func saveHistory(history []models.CommandEntry) error {
	file, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(historyFile, file, 0644)
}
