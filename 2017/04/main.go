package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	lines := conv.SplitNewline(input)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	valid := 0
	for _, line := range lines {
		if isValid(line) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValid(line string) bool {
	words := strings.Split(line, " ")
	seen := make(map[string]bool)
	for _, word := range words {
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}

func part2(lines []string) {
	valid := 0
	for _, line := range lines {
		if isValid2(line) {
			valid++
		}
	}
	fmt.Println(valid)
}

func isValid2(line string) bool {
	words := strings.Split(line, " ")
	seen := make(map[string]bool)
	for _, word := range words {
		chars := strings.Split(word, "")
		slices.Sort(chars)
		word = strings.Join(chars, "")
		if seen[word] {
			return false
		}
		seen[word] = true
	}
	return true
}
