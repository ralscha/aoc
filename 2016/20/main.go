package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/rangeutil"
	"fmt"
	"log"
	"strings"
)

func parseInput(input string) []rangeutil.Range {
	lines := conv.SplitNewline(input)
	ranges := make([]rangeutil.Range, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "-")
		ranges[i] = rangeutil.NewRange(conv.MustAtoi(parts[0]), conv.MustAtoi(parts[1]))
	}
	return ranges
}

func main() {
	input, err := download.ReadInput(2016, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	ranges := parseInput(input)
	mergedRanges := rangeutil.Merge(ranges)
	for i := range len(mergedRanges) - 1 {
		if mergedRanges[i].End+1 < mergedRanges[i+1].Start {
			fmt.Println("Part 1", mergedRanges[i].End+1)
			return
		}
	}
}

func part2(input string) {
	ranges := parseInput(input)
	mergedRanges := rangeutil.Merge(ranges)

	allowed := 0
	for i := range len(mergedRanges) - 1 {
		allowed += mergedRanges[i+1].Start - (mergedRanges[i].End + 1)
	}

	fmt.Println("Part 2", allowed)
}
