package config

import (
	"encoding/json"
	"golaunch-cli/internal/models"
	"os"
)

func LoadConfig(path string) (models.Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg models.Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func SaveConfig(path string, cfg models.Config) error {
	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0644)
}
