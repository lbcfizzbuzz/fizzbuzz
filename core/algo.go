package core

import (
	"strconv"
)

// Fizzbuzz returns a string containing:
// - str1 if nb is a multiple of firstDivisor
// - str2 if nb is a multiple of secondDivisor
// - a concatenation of str1 and str2 if nb is a multiple of both firstDivisor and secondDivisor
// - an empty string if the division is impossible
func Fizzbuzz(nb, firstDivisor, secondDivisor uint64, str1, str2 string) string {
	result := ""
	if firstDivisor == 0 || secondDivisor == 0 {
		return result
	}

	if nb%firstDivisor == 0 {
		result = str1
	}
	if nb%secondDivisor == 0 {
		result += str2
	}
	if result == "" {
		result = strconv.FormatUint(nb, 10)
	}
	return result
}
