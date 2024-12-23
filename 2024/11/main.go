package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func simulateBlink(stones map[int]int64) map[int]int64 {
	result := make(map[int]int64)

	for stone, count := range stones {
		if stone == 0 {
			// Rule 1: 0 becomes 1
			result[1] += count
			continue
		}

		stoneStr := strconv.Itoa(stone)
		digits := len(stoneStr)

		if digits%2 == 0 {
			// Rule 2: Split into two stones
			mid := digits / 2
			left := conv.MustAtoi(stoneStr[:mid])
			right := conv.MustAtoi(stoneStr[mid:])
			result[left] += count
			result[right] += count
		} else {
			// Rule 3: Multiply by 2024
			result[stone*2024] += count
		}
	}

	return result
}

func main() {
	input, err := download.ReadInput(2024, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 25)
	part1and2(input, 75)
}

func part1and2(input string, iterations int) {
	initialStones := conv.ToIntSlice(strings.Fields(input))
	stones := make(map[int]int64)
	for _, stone := range initialStones {
		stones[stone]++
	}

	for range iterations {
		stones = simulateBlink(stones)
	}

	var total int64
	for _, count := range stones {
		total += count
	}
	if iterations == 25 {
		fmt.Println("Part 1", total)
	} else {
		fmt.Println("Part 2", total)
	}
}
