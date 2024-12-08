package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	splitted := strings.Split(input, "-")
	start := conv.MustAtoi(splitted[0])
	end := conv.MustAtoi(splitted[1])

	count := 0

	for i := start; i <= end; i++ {
		if isValidPassword(i) {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func isValidPassword(pw int) bool {
	pwStr := fmt.Sprintf("%d", pw)
	if !stringutil.HasRepeatedChar(pwStr) {
		return false
	}
	
	// Check if digits are non-decreasing
	for i := 0; i < len(pwStr)-1; i++ {
		if pwStr[i] > pwStr[i+1] {
			return false
		}
	}

	return true
}

func part2(input string) {
	splitted := strings.Split(input, "-")
	start := conv.MustAtoi(splitted[0])
	end := conv.MustAtoi(splitted[1])

	count := 0

	for i := start; i <= end; i++ {
		if isValidPassword2(i) {
			count++
		}
	}

	fmt.Println("Part 2:", count)
}

func isValidPassword2(pw int) bool {
	pwStr := fmt.Sprintf("%d", pw)
	hasExactDouble := false
	groupSize := 1

	for i := 0; i < len(pwStr)-1; i++ {
		if pwStr[i] > pwStr[i+1] {
			return false
		}
		if pwStr[i] == pwStr[i+1] {
			groupSize++
		} else {
			if groupSize == 2 {
				hasExactDouble = true
			}
			groupSize = 1
		}
	}

	if groupSize == 2 {
		hasExactDouble = true
	}

	return hasExactDouble
}
