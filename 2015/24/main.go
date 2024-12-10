package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2015, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 3)
	part1and2(input, 4)
}

func part1and2(input string, groups int) {
	lines := conv.SplitNewline(input)
	packages := make([]int, len(lines))
	sum := 0
	for i, line := range lines {
		packages[i] = conv.MustAtoi(line)
		sum += packages[i]
	}

	slices.SortFunc(packages, func(i, j int) int {
		return j - i
	})
	targetWeight := sum / groups

	var minQE uint64
	found := false

	for size := 1; size <= len(packages) && !found; size++ {
		minQE = findMinQE(packages, targetWeight, size, 0, 0, 0, 1)
		if minQE > 0 {
			found = true
		}
	}

	if found {
		if groups == 3 {
			fmt.Println("Part 1", minQE)
		} else {
			fmt.Println("Part 2", minQE)
		}
	}
}

func findMinQE(packages []int, target, size, pos, sum int, count int, qe uint64) uint64 {
	if sum > target || count > size {
		return 0
	}
	if count == size {
		if sum == target {
			return qe
		}
		return 0
	}
	if pos >= len(packages) {
		return 0
	}

	minQE := findMinQE(packages, target, size, pos+1, sum, count, qe)

	if newQE := findMinQE(packages, target, size, pos+1, sum+packages[pos], count+1, qe*uint64(packages[pos])); newQE > 0 {
		if minQE == 0 || newQE < minQE {
			minQE = newQE
		}
	}

	return minQE
}
