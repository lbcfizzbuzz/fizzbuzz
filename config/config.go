package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration represents the configuration of the server
type Configuration struct {
	DbType string `json:"db_type"`
	Dsn    string `json:"dsn"`
}

// Read reads the config file passed in param
func (config *Configuration) Read(path string) error {
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened " + path)
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(config)
	if err != nil {
		return err
	}
	return nil
}
