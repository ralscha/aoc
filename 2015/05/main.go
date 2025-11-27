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
	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	count := 0
	for _, line := range lines {
		if isNice2(line) {
			count++
		}
	}
	fmt.Println("Part 2", count)
}

func isNice(line string) bool {
	if stringutil.CountVowels(line) < 3 {
		return false
	}

	if !stringutil.HasRepeatedChar(line) {
		return false
	}

	if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		return false
	}

	return true
}

func isNice2(line string) bool {
	if !stringutil.HasRepeatedPair(line) {
		return false
	}

	return stringutil.HasSandwichedChar(line)
}
