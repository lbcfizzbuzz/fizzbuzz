package tests

import (
	conf "github.com/lbcfizzbuzz/fizzbuzz/config"
	"testing"
)

func TestConfigRead(t *testing.T) {
	config := conf.Configuration{}

	config.Port = 10
	if config.Validate() == nil {
		t.Errorf("failed to validate the not valid port")
	}

	config.Port = 8080
	if config.Validate() != nil {
		t.Errorf("failed to validate the valid port")
	}
}
