package conv

import (
	"log"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("converting to int failed: %v", err)
	}
	return i
}

func SplitNewline(s string) []string {
	splitted := strings.Split(s, "\n")
	if len(splitted) > 0 && splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return splitted
}

func ToIntSlice(s []string) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = MustAtoi(v)
	}
	return result
}

func ToIntSliceComma(s string) []int {
	splitted := strings.Split(s, ",")
	result := make([]int, len(splitted))
	for i, v := range splitted {
		result[i] = MustAtoi(v)
	}
	return result
}
