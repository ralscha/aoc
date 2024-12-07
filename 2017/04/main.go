package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
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
	seen := container.NewSet[string]()
	for _, word := range words {
		if seen.Contains(word) {
			return false
		}
		seen.Add(word)
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
	seen := container.NewSet[string]()
	for _, word := range words {
		chars := strings.Split(word, "")
		slices.Sort(chars)
		word = strings.Join(chars, "")
		if seen.Contains(word) {
			return false
		}
		seen.Add(word)
	}
	return true
}
