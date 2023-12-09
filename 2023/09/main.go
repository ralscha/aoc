package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)

	sumLast := 0
	sumFirst := 0
	for _, line := range lines {
		numbers := conv.ToIntSlice(strings.Fields(line))

		var sequence [][]int
		for !allZeroes(numbers) {
			sequence = append(sequence, numbers)
			differences := calculateDifferences(numbers)
			numbers = differences
		}

		extrapolatedLast := 0
		extrapolatedFirst := 0
		for i := len(sequence) - 1; i >= 0; i-- {
			extrapolatedLast += sequence[i][len(sequence[i])-1]
			extrapolatedFirst = sequence[i][0] - extrapolatedFirst
		}
		sumLast += extrapolatedLast
		sumFirst += extrapolatedFirst
	}

	fmt.Println(sumLast)
	fmt.Println(sumFirst)
}

func calculateDifferences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		differences[i] = numbers[i+1] - numbers[i]
	}
	return differences
}

func allZeroes(numbers []int) bool {
	for _, num := range numbers {
		if num != 0 {
			return false
		}
	}
	return true
}
