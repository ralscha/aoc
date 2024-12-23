package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2021, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	numbers := conv.ToIntSlice(lines)
	increasing := 0
	lastNum := -1

	for _, num := range numbers {
		if lastNum != -1 && num > lastNum {
			increasing++
		}
		lastNum = num
	}
	fmt.Println("Part 1", increasing)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	numbers := conv.ToIntSlice(lines)
	increasing := 0
	lastWindow := -1

	for i := range len(numbers) - 2 {
		windowSum := numbers[i] + numbers[i+1] + numbers[i+2]
		if lastWindow != -1 && windowSum > lastWindow {
			increasing++
		}
		lastWindow = windowSum
	}
	fmt.Println("Part 2", increasing)
}
