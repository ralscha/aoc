package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	fmt.Println(input)
	if input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	for !isValid(input) {
		input = increment(input)
	}
	next := input
	fmt.Println(next)

	next = increment(next)
	for !isValid(next) {
		next = increment(next)
	}
	fmt.Println(next)
}

func increment(s string) string {
	lastPos := len(s) - 1
	s = s[:lastPos] + string(s[lastPos]+1)
	if s[lastPos] > 'z' {
		s = increment(s[:lastPos]) + "a"
	}
	return s
}

func isValid(s string) bool {
	if containsForbiddenLetters(s) {
		return false
	}
	if !containsIncreasingStraight(s) {
		return false
	}
	if !containsTwoNonOverlappingPairs(s) {
		return false
	}
	return true
}

func containsForbiddenLetters(s string) bool {
	for _, c := range s {
		if c == 'i' || c == 'o' || c == 'l' {
			return true
		}
	}
	return false
}

func containsIncreasingStraight(s string) bool {
	for i := range len(s) - 2 {
		if s[i+1] == s[i]+1 && s[i+2] == s[i]+2 {
			return true
		}
	}
	return false
}

func containsTwoNonOverlappingPairs(s string) bool {
	pairs := 0
	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}
