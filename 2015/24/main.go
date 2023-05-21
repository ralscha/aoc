package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	inputFile := "./2015/24/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 3)
	part1and2(input, 4)
}

func part1and2(input string, groups int) {
	lines := conv.SplitNewline(input)
	packages := make([]int, len(lines))
	for i, line := range lines {
		packages[i] = conv.MustAtoi(line)
	}

	sum := 0
	for _, num := range packages {
		sum += num
	}
	targetWeight := sum / groups

	combs := combinations(packages, targetWeight)

	leastAmountOfPackages := math.MaxInt
	smallestQuantumEntanglement := math.MaxInt
	for _, group := range combs {
		if len(group) < leastAmountOfPackages {
			leastAmountOfPackages = len(group)
			smallestQuantumEntanglement = quantumEntanglement(group)
		} else if len(group) == leastAmountOfPackages {
			qe := quantumEntanglement(group)
			if qe < smallestQuantumEntanglement {
				smallestQuantumEntanglement = qe
			}
		}
	}

	fmt.Println(smallestQuantumEntanglement)

}

func quantumEntanglement(packages []int) int {
	product := 0
	for _, n := range packages {
		if product > 0 {
			product *= n
		} else {
			product = n
		}
	}
	return product
}

func combinations(input []int, targetWeight int) [][]int {
	var results [][]int
	for i := 0; i < 1<<uint(len(input)); i++ {
		var combination []int
		for ix, in := range input {
			if i&(1<<uint(ix)) > 0 {
				combination = append(combination, in)
			}
		}
		if len(combination) > 0 && sum(combination) == targetWeight {
			results = append(results, combination)
		}
	}
	return results
}

func sum(input []int) int {
	sum := 0
	for _, n := range input {
		sum += n
	}
	return sum
}
