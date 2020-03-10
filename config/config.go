package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Configuration represents the configuration of the server
type Configuration struct {
	Datastore string `json:"datastore"`
	Dsn       string `json:"dsn"`
	Port      int    `json:"port"`
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

// Validate check for errors in the configuration
func (config *Configuration) Validate() error {
	// Check the port
	if config.Port < 1024 || config.Port > 65534 {
		return errors.New("your port should be in the range 1024 - 65534 please check your configuration file")
	}
	return nil
}
