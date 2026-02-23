package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/rangeutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rangeLines := conv.SplitNewline(sections[0])
	idLines := conv.SplitNewline(sections[1])

	var ranges []rangeutil.Range
	for _, line := range rangeLines {
		parts := strings.SplitN(line, "-", 2)
		start := conv.MustAtoi(parts[0])
		end := conv.MustAtoi(parts[1])
		ranges = append(ranges, rangeutil.NewRange(start, end))
	}

	fresh := 0
	for _, line := range idLines {
		id := conv.MustAtoi(line)
		for _, r := range ranges {
			if r.Contains(id) {
				fresh++
				break
			}
		}
	}
	fmt.Println("Part 1", fresh)
}

func part2(input string) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rangeLines := conv.SplitNewline(sections[0])

	var ranges []rangeutil.Range
	for _, line := range rangeLines {
		parts := strings.SplitN(line, "-", 2)
		start := conv.MustAtoi(parts[0])
		end := conv.MustAtoi(parts[1])
		ranges = append(ranges, rangeutil.NewRange(start, end))
	}

	merged := rangeutil.Merge(ranges)
	total := 0
	for _, r := range merged {
		total += r.Length()
	}
	fmt.Println("Part 2", total)
}
