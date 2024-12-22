package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1Result := part1(input)
	part2(input, part1Result)
}

func isValid(target int, preamble []int) bool {
	for i := range len(preamble) {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == target {
				return true
			}
		}
	}
	return false
}

func part1(input string) int {
	lines := conv.SplitNewline(input)
	numbers := conv.ToIntSlice(lines)
	preambleLength := 25

	for i := preambleLength; i < len(numbers); i++ {
		preamble := numbers[i-preambleLength : i]
		if !isValid(numbers[i], preamble) {
			fmt.Println("Part 1", numbers[i])
			return numbers[i]
		}
	}
	return -1
}

func part2(input string, invalidNumber int) {
	lines := conv.SplitNewline(input)
	numbers := conv.ToIntSlice(lines)

	for i := range len(numbers) {
		for j := i + 2; j <= len(numbers); j++ {
			subArray := numbers[i:j]
			sum := 0
			for _, num := range subArray {
				sum += num
			}
			if sum == invalidNumber {
				minValue := subArray[0]
				maxValue := subArray[0]
				for _, num := range subArray {
					if num < minValue {
						minValue = num
					}
					if num > maxValue {
						maxValue = num
					}
				}
				fmt.Println("Part 2", minValue+maxValue)
				return
			}
		}
	}
}
