package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func isTrap(prev string, pos int) bool {
	left := pos > 0 && prev[pos-1] == '^'
	center := prev[pos] == '^'
	right := pos < len(prev)-1 && prev[pos+1] == '^'

	return (left && center && !right) ||
		(!left && center && right) ||
		(left && !center && !right) ||
		(!left && !center && right)
}

func generateNextRow(currentRow string) string {
	var nextRow strings.Builder
	nextRow.Grow(len(currentRow))

	for i := range len(currentRow) {
		if isTrap(currentRow, i) {
			nextRow.WriteRune('^')
		} else {
			nextRow.WriteRune('.')
		}
	}

	return nextRow.String()
}

func countSafeTiles(firstRow string, totalRows int) int {
	if totalRows <= 0 {
		return 0
	}

	safeTiles := strings.Count(firstRow, ".")
	currentRow := firstRow

	for row := 1; row < totalRows; row++ {
		currentRow = generateNextRow(currentRow)
		safeTiles += strings.Count(currentRow, ".")
	}

	return safeTiles
}

func part1(input string) {
	result := countSafeTiles(input, 40)
	fmt.Println("Part 1", result)
}

func part2(input string) {
	result := countSafeTiles(input, 400000)
	fmt.Println("Part 2", result)
}
