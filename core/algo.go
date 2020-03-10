package core

import (
	"errors"
	"strconv"
)

// Fizzbuzz returns a slice of strings with numbers from 1 to limit included.
// All multiples of int1 are replaced by str1.
// All multiples of int2 are replaced by str2.
// All multiples of int1 and int2 are replaced by str1str2.
func Fizzbuzz(int1, int2, limit uint64, str1, str2 string) ([]string, error) {
	if int1 == 0 || int2 == 0 {
		return nil, errors.New("The int1 and int2 parameters must be greater than 0")
	}

	var result []string
	for i := uint64(1); i <= limit; i++ {
		currentStr := ""
		if i%int1 == 0 {
			currentStr = str1
		}
		if i%int2 == 0 {
			currentStr += str2
		}
		if currentStr == "" {
			currentStr = strconv.FormatUint(i, 10)
		}
		result = append(result, currentStr)
	}

	return result, nil
}
