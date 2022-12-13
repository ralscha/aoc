package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2015/17/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	containerSizes := conv.ToIntSlice(lines)
	totalEggnogLiters := 150

	combinations := mathx.Combinations(containerSizes)
	totalCombinations := 0
	for _, combination := range combinations {
		total := 0
		for _, v := range combination {
			total += v
		}
		if total == totalEggnogLiters {
			totalCombinations++
		}
	}
	fmt.Println(totalCombinations)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	containerSizes := conv.ToIntSlice(lines)
	totalEggnogLiters := 150

	combinations := mathx.Combinations(containerSizes)
	minContainers := 0
	totalCombinations := 0
	for _, combination := range combinations {
		total := 0
		for _, v := range combination {
			total += v
		}
		if total == totalEggnogLiters {
			combinationLength := len(combination)
			if minContainers == 0 || combinationLength < minContainers {
				minContainers = combinationLength
				totalCombinations = 1
			} else if combinationLength == minContainers {
				totalCombinations++
			}
		}
	}
	fmt.Println(totalCombinations)
}
