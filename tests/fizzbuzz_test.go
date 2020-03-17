package tests

import (
	"github.com/lbcfizzbuzz/fizzbuzz/core"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	// Test error handling
	badParameters := []struct {
		int1 uint64
		int2 uint64
	}{
		{0, 5},
		{5, 0},
		{0, 0},
	}

	for i, bad := range badParameters {
		result := core.Fizzbuzz(uint64(i), bad.int1, bad.int2, "fizz", "buzz")
		if result != "" {
			t.Errorf("Error not handled correctly")
		}
	}

	// Test with no error
	expectedResults := []string{
		"1",
		"2",
		"fizz",
		"4",
		"buzz",
		"fizz",
		"7",
		"8",
		"fizz",
		"buzz",
		"11",
		"fizz",
		"13",
		"14",
		"fizzbuzz",
	}
	for i, expected := range expectedResults {
		result := core.Fizzbuzz(uint64(i+1), 3, 5, "fizz", "buzz")
		if result == "" {
			t.Errorf("An unhandled error occured")
		}
		if result != expected {
			t.Errorf("Expected: " + expected + " received: " + result)
		}
	}
}
