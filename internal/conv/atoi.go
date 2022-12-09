package conv

import (
	"log"
	"strconv"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("converting to int failed: %v", err)
	}
	return i
}

func MustAtoiArray(s []string) []int {
	var result []int
	for _, v := range s {
		result = append(result, MustAtoi(v))
	}
	return result
}
