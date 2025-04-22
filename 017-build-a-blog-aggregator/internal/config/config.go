package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, configFileName)
	return filePath, nil
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}

func write(config Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	fileData, err := json.Marshal(config)

	if err != nil {
		return err
	}

	os.WriteFile(filePath, fileData, 0666)

	return nil
}
