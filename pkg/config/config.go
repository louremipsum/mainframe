package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	AIModel      string `json:"ai_model"`
	APIKey       string `json:"api_key"`
	Debug        bool   `json:"debug"`
	Logs         bool   `json:"logs"`
	Experimental bool   `json:"experimental"`
}

var (
	defaultConfig = Config{
		AIModel:      "local",
		APIKey:       "",
		Debug:        false,
		Logs:         false,
		Experimental: false,
	}
	configDir  string
	configPath string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	configDir = filepath.Join(homeDir, ".mainframe")
	configPath = filepath.Join(configDir, "config.json")
}

func Load() (*Config, error) {
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return &defaultConfig, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &defaultConfig, Save(&defaultConfig)
		}
		return &defaultConfig, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &defaultConfig, err
	}

	return &config, nil
}

func Save(config *Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
