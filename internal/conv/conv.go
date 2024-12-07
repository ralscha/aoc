package conv

import (
	"log"
	"strconv"
	"strings"
)

// MustAtoi converts a string to an integer.
// If the conversion fails, it logs a fatal error and terminates the program.
func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("converting to int failed: %v", err)
	}
	return i
}

// MustAtoi64 converts a string to a 64-bit integer.
// If the conversion fails, it logs a fatal error and terminates the program.
func MustAtoi64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("converting to int64 failed: %v", err)
	}
	return i
}

// SplitNewline splits a string by newline characters.
// It removes any empty line at the end of the resulting slice.
func SplitNewline(s string) []string {
	splitted := strings.Split(s, "\n")
	if len(splitted) > 0 && splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return splitted
}

// ToIntSlice converts a slice of strings to a slice of integers.
// Uses MustAtoi for conversion, so it will terminate the program if any conversion fails.
func ToIntSlice(s []string) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = MustAtoi(v)
	}
	return result
}

// ToIntSliceComma splits a comma-separated string into a slice of integers.
// Uses MustAtoi for conversion, so it will terminate the program if any conversion fails.
func ToIntSliceComma(s string) []int {
	splitted := strings.Split(s, ",")
	result := make([]int, len(splitted))
	for i, v := range splitted {
		result[i] = MustAtoi(v)
	}
	return result
}

// ToInt64SliceComma splits a comma-separated string into a slice of 64-bit integers.
// Uses MustAtoi64 for conversion, so it will terminate the program if any conversion fails.
func ToInt64SliceComma(s string) []int64 {
	splitted := strings.Split(s, ",")
	result := make([]int64, len(splitted))
	for i, v := range splitted {
		result[i] = MustAtoi64(v)
	}
	return result
}
