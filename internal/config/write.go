package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return  err
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil

}