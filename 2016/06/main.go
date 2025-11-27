package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2016, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1and2(lines)
}

func part1and2(lines []string) {
	bags := make([]*container.Bag[rune], len(lines[0]))
	for i := range bags {
		bags[i] = container.NewBag[rune]()
	}

	for _, line := range lines {
		for i, c := range line {
			bags[i].Add(c)
		}
	}

	resultMax := ""
	resultMin := ""
	for _, bag := range bags {
		maxCount := 0
		maxChar := rune(0)
		minCount := math.MaxInt
		minChar := rune(0)

		for char, count := range bag.Values() {
			if count > maxCount {
				maxCount = count
				maxChar = char
			}
			if count < minCount {
				minCount = count
				minChar = char
			}
		}
		resultMax += string(maxChar)
		resultMin += string(minChar)
	}
	fmt.Println(resultMax)
	fmt.Println(resultMin)
}
