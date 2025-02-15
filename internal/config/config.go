package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string) (map[string][]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg map[string][]string
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func SaveConfig(path string, cfg map[string][]string) error {
	file, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0644)
}
