package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2025, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		max := 0
		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				joltage := conv.MustAtoi(line[i:i+1])*10 + conv.MustAtoi(line[j:j+1])
				if joltage > max {
					max = joltage
				}
			}
		}
		sum += max
	}
	fmt.Println("Part 1", sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	k := 12
	var sum int64
	for _, line := range lines {
		n := len(line)
		start := 0
		var result int64
		for i := range k {
			end := n - (k - i)
			maxDigit := byte('0')
			maxPos := start
			for p := start; p <= end; p++ {
				if line[p] > maxDigit {
					maxDigit = line[p]
					maxPos = p
				}
			}
			result = result*10 + conv.MustAtoi64(string(maxDigit))
			start = maxPos + 1
		}
		sum += result
	}
	fmt.Println("Part 2", sum)
}
