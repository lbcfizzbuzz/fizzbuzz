package tests

import (
	cfg "github.com/lbcfizzbuzz/fizzbuzz/config"
	ds "github.com/lbcfizzbuzz/fizzbuzz/datastore"
)

// GetInitializedDatastore creates a datastore according to
// the configuration, and initialize its connection
func GetInitializedDatastore() (ds.Datastore, error) {
	// Read the configuration file
	config := cfg.Configuration{}
	if err := config.Read("../config/config-test.json"); err != nil {
		return nil, err
	}

	// Get the datastore
	datastore, err := ds.GetDatastore(config)
	if err != nil {
		return nil, err
	}

	// Initialize the connection of the datastore
	err = datastore.Init()
	if err != nil {
		return nil, err
	}
	return datastore, nil
}
