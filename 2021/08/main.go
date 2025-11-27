package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

type segments [7]bool

var (
	zero  = segments{true, true, true, false, true, true, true}
	one   = segments{false, false, true, false, false, true, false}
	two   = segments{true, false, true, true, true, false, true}
	three = segments{true, false, true, true, false, true, true}
	four  = segments{false, true, true, true, false, true, false}
	five  = segments{true, true, false, true, false, true, true}
	six   = segments{true, true, false, true, true, true, true}
	seven = segments{true, false, true, false, false, true, false}
	eight = segments{true, true, true, true, true, true, true}
	nine  = segments{true, true, true, true, false, true, true}

	numbers = [10]segments{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
)

func main() {
	input, err := download.ReadInput(2021, 8)
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
		pos := strings.IndexByte(line, '|')
		outputs := strings.Fields(line[pos+1:])
		for _, output := range outputs {
			switch len(output) {
			case 2, 3, 4, 7: // digits 1, 7, 4, 8
				count++
			}
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	sum := 0

	chars := []string{"a", "b", "c", "d", "e", "f", "g"}
	allPermutations := mathx.Permutations(chars)

	for _, line := range lines {
		pos := strings.IndexByte(line, '|')
		patterns := strings.Fields(line[:pos])
		outputs := strings.Fields(line[pos+1:])

		for _, perm := range allPermutations {
			if isValidPermutation(patterns, perm) {
				value := 0
				for _, output := range outputs {
					seg := convertToSegments(perm, output)
					digit := findPattern(seg)
					value = value*10 + digit
				}
				sum += value
				break
			}
		}
	}

	fmt.Println("Part 2", sum)
}

func isValidPermutation(patterns []string, perm []string) bool {
	seen := container.NewSet[int]()
	for _, pattern := range patterns {
		seg := convertToSegments(perm, pattern)
		digit := findPattern(seg)
		if digit == -1 || seen.Contains(digit) {
			return false
		}
		seen.Add(digit)
	}
	return seen.Len() == 10
}

func convertToSegments(perm []string, str string) segments {
	var result segments
	for _, c := range str {
		for i, p := range perm {
			if p == string(c) {
				result[i] = true
				break
			}
		}
	}
	return result
}

func findPattern(seg segments) int {
	for i, pattern := range numbers {
		if pattern == seg {
			return i
		}
	}
	return -1
}
