package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	input = strings.TrimSpace(input)
	fmt.Println("Part 1", len(react(input)))
}

func part2(input string) {
	input = strings.TrimSpace(input)
	minLen := len(input)

	// Use Set to track unique lowercase letters in input
	uniqueChars := container.NewSet[rune]()
	for _, c := range input {
		uniqueChars.Add(toLower(c))
	}

	// Try removing each unique character
	for _, c := range uniqueChars.Values() {
		// Remove both cases of the current character
		filtered := strings.Map(func(r rune) rune {
			if toLower(r) == c {
				return -1 // Will be removed by strings.Map
			}
			return r
		}, input)

		// React the filtered polymer
		result := react(filtered)
		if len(result) < minLen {
			minLen = len(result)
		}
	}

	fmt.Println("Part 2", minLen)
}

// react performs the polymer reaction
func react(polymer string) string {
	stack := make([]rune, 0, len(polymer))

	for _, c := range polymer {
		if len(stack) > 0 && areReactive(stack[len(stack)-1], c) {
			stack = stack[:len(stack)-1] // Pop
		} else {
			stack = append(stack, c) // Push
		}
	}

	return string(stack)
}

// areReactive checks if two units react
func areReactive(a, b rune) bool {
	return a != b && toLower(a) == toLower(b)
}

// toLower converts a rune to lowercase
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}
