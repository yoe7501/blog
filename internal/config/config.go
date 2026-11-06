package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := Write(*c)
	return err
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path = path + configFileName
	return path, nil
}

func Write(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0o644)
	return err
}

func Read() (Config, error) {
	var data Config
	path, err := getConfigFilePath()
	if err != nil {
		return data, err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
