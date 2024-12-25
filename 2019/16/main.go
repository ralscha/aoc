package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strconv"
)

func main() {
	input, err := download.ReadInput(2019, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	digits := make([]int, len(input))
	for i, r := range input {
		digits[i] = int(r - '0')
	}

	phases := 100
	current := digits
	for range phases {
		next := make([]int, len(current))
		for i := range next {
			pattern := []int{0, 1, 0, -1}
			sum := 0
			for j := range current {
				patternIndex := ((j + 1) / (i + 1)) % len(pattern)
				sum += current[j] * pattern[patternIndex]
			}
			next[i] = mathx.Abs(sum) % 10
		}
		current = next
	}

	result := ""
	for i := range 8 {
		result += strconv.Itoa(current[i])
	}

	fmt.Println("Part 1", result)
}

func part2(input string) {
	digits := make([]int, len(input)*10000)
	for i := range 10000 {
		for j, r := range input {
			digits[i*len(input)+j] = int(r - '0')
		}
	}

	offsetStr := input[:7]
	offset := conv.MustAtoi(offsetStr)

	phases := 100
	current := digits
	n := len(current)

	for range phases {
		next := make([]int, n)
		sum := 0
		for i := n - 1; i >= 0; i-- {
			sum += current[i]
			next[i] = sum % 10
		}
		current = next
	}

	result := ""
	for i := range 8 {
		result += strconv.Itoa(current[offset+i])
	}

	fmt.Println("Part 2", result)
}
