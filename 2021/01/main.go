package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputFile := "./2021/01/input.txt"
	input, err := download.ReadInput(inputFile, 2021, 1)
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
		num, err := strconv.Atoi(nums)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", nums, err)
		}

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
		num, err := strconv.Atoi(nums)
		if err != nil {
			log.Fatalf("conversion failed: %s %v", nums, err)
		}
		windowSum += num
		if windowCount == 3 {
			if windowSum > lastWindow && lastWindow != -1 {
				increasing += 1
			}
			lastWindow = windowSum

			firstNumOfWindow, err := strconv.Atoi(numbers[ix-2])
			if err != nil {
				log.Fatalf("conversion failed: %s %v", nums, err)
			}
			windowSum = windowSum - firstNumOfWindow
		} else {
			windowCount += 1
			windowSum += num
		}

	}
	fmt.Printf("increasing %d\n", increasing)
}
