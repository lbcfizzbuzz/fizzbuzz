package main

import (
	"errors"
	"flag"
)

const defaultConfigPath = "../../config/config-dev.json"

// Args represents the command line arguments values
type Args struct {
	ConfigPath string
}

// Init parses the command line arguments
func (a *Args) Init() {
	flag.StringVar(&(a.ConfigPath), "config", defaultConfigPath, "The path to the config file to use")
	flag.Parse()
}

// Check validates the command line arguments
func (a *Args) Check() error {
	if a.ConfigPath == "" {
		return errors.New("a path to a config file should be specified")
	}
	return nil
}
