package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
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
	numbers := strings.Fields(input)
	increasing := 0
	lastNum := -1

	for _, nums := range numbers {
		num := conv.MustAtoi(nums)

		if lastNum != -1 && num > lastNum {
			increasing += 1
		}
		lastNum = num
	}
	fmt.Printf("increasing %d\n", increasing)
}

func part2(input string) {
	numbers := strings.Fields(input)
	increasing := 0
	lastWindow := -1
	windowCount := 0
	windowSum := 0

	for ix, nums := range numbers {
		num := conv.MustAtoi(nums)
		windowSum += num
		if windowCount == 3 {
			if windowSum > lastWindow && lastWindow != -1 {
				increasing += 1
			}
			lastWindow = windowSum

			firstNumOfWindow := conv.MustAtoi(numbers[ix-2])
			windowSum = windowSum - firstNumOfWindow
		} else {
			windowCount += 1
			windowSum += num
		}

	}
	fmt.Printf("increasing %d\n", increasing)
}
