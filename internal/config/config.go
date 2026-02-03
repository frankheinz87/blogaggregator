package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	filePath := homeDir + "/.gatorconfig.json"

	config, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(config, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	filePath := homeDir + "/.gatorconfig.json"
	fmt.Println("Writing to:", filePath)

	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
