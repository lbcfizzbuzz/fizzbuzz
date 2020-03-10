package tests

import (
	"github.com/lbcfizzbuzz/fizzbuzz/core"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	// Test error handling
	tables := []struct {
		int1 uint64
		int2 uint64
	}{
		{0, 5},
		{5, 0},
		{0, 0},
	}

	for _, table := range tables {
		_, err := core.Fizzbuzz(table.int1, table.int2, 10, "fizz", "buzz")
		if err == nil {
			t.Errorf("Error not handled correctly")
		}
	}

	// Test with no error
	_, err := core.Fizzbuzz(3, 5, 10, "fizz", "buzz")
	if err != nil {
		t.Errorf("An unhandled error occured")
	}
}
