package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	count := 0
	for _, line := range lines {
		if isNice(line) {
			count++
		}
	}
	fmt.Println(count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	count := 0
	for _, line := range lines {
		if isNice2(line) {
			count++
		}
	}
	fmt.Println(count)
}

func isNice(line string) bool {
	// It contains at least three vowels
	if stringutil.CountVowels(line) < 3 {
		return false
	}

	// It contains at least one letter that appears twice in a row
	if !stringutil.HasRepeatedChar(line) {
		return false
	}

	// It does not contain the strings: ab, cd, pq, or xy
	if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		return false
	}

	return true
}

func isNice2(line string) bool {
	// It contains a pair of any two letters that appears at least twice in the string without overlapping
	if !stringutil.HasRepeatedPair(line) {
		return false
	}

	// It contains at least one letter which repeats with exactly one letter between them
	return stringutil.HasSandwichedChar(line)
}
