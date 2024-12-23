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
	input, err := download.ReadInput(2022, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	fullyContain := 0
	partiallyContain := 0
	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, ",")
		if len(splitted) != 2 {
			continue
		}
		firstPair := strings.Split(splitted[0], "-")
		secondPair := strings.Split(splitted[1], "-")

		range1 := rangeutil.NewRange(
			conv.MustAtoi(firstPair[0]),
			conv.MustAtoi(firstPair[1]),
		)
		range2 := rangeutil.NewRange(
			conv.MustAtoi(secondPair[0]),
			conv.MustAtoi(secondPair[1]),
		)

		// For full containment, one range must completely contain the other
		if range1.Contains(range2.Start) && range1.Contains(range2.End) ||
			range2.Contains(range1.Start) && range2.Contains(range1.End) {
			fullyContain++
		}

		// For partial containment, the ranges must overlap
		if range1.Overlaps(range2) {
			partiallyContain++
		}
	}
	fmt.Println(fullyContain)
	fmt.Println(partiallyContain)
}
