package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2015, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func sumCombination(combination []int) int {
	total := 0
	for _, v := range combination {
		total += v
	}
	return total
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	containerSizes := conv.ToIntSlice(lines)
	totalEggnogLiters := 150

	combinations := mathx.Combinations(containerSizes)

	validCombos := slices.DeleteFunc(combinations, func(combo []int) bool {
		return sumCombination(combo) != totalEggnogLiters
	})

	fmt.Println(len(validCombos))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	containerSizes := conv.ToIntSlice(lines)
	totalEggnogLiters := 150

	combinations := mathx.Combinations(containerSizes)

	validCombos := slices.DeleteFunc(combinations, func(combo []int) bool {
		return sumCombination(combo) != totalEggnogLiters
	})

	minCombo := slices.MinFunc(validCombos, func(a, b []int) int {
		return len(a) - len(b)
	})
	minLength := len(minCombo)

	minLengthCombos := slices.DeleteFunc(validCombos, func(combo []int) bool {
		return len(combo) != minLength
	})

	fmt.Println(len(minLengthCombos))
}
