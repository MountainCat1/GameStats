package tools

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type ServerConfig struct {
	Port int `json:"port"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type DatabaseConfig struct {
	Provider string `json:"provider"`

	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

// ConfigNotFoundError is the error returned when the configuration file is not found.
var ConfigNotFoundError = fmt.Errorf("config file not found")

// LoadConfig loads the configuration from a JSON file.
func LoadConfig() (Config, error) {
	configFileName := getConfigFilePath()

	log.Info("Loading config from ", configFileName)
	file, err := os.Open(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, ConfigNotFoundError
		}
		return Config{}, fmt.Errorf("error opening config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding config JSON: %w", err)
	}

	return config, nil
}

// getConfigFilePath returns the path to the configuration file.
func getConfigFilePath() string {
	environmentName := os.Getenv("ENVIRONMENT")
	if environmentName == "" {
		environmentName = "local"
	}
	return filepath.Join("config", fmt.Sprintf("%s.json", environmentName))
}
