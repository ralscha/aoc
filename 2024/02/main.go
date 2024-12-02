package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	totalSafe := 0
	for _, line := range lines {
		levels := conv.ToIntSlice(strings.Fields(line))
		if isSafe(levels) {
			totalSafe++
		}
	}
	fmt.Println("Part 1", totalSafe)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	totalSafe := 0

	for _, line := range lines {
		numbers := conv.ToIntSlice(strings.Fields(line))

		for i := range len(numbers) {
			reduced := make([]int, 0, len(numbers)-1)
			reduced = append(reduced, numbers[:i]...)
			reduced = append(reduced, numbers[i+1:]...)

			if isSafe(reduced) {
				totalSafe++
				break
			}
		}
	}
	fmt.Println("Part 2", totalSafe)
}

func isSafe(levels []int) bool {
	ascending := true
	descending := true

	for i := range len(levels) - 1 {
		diff := mathx.Abs(levels[i] - levels[i+1])
		if diff < 1 || diff > 3 {
			return false
		}
		if levels[i] > levels[i+1] {
			ascending = false
		}
		if levels[i] < levels[i+1] {
			descending = false
		}
		if !ascending && !descending {
			return false
		}
	}

	return ascending || descending
}
