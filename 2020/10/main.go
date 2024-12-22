package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2020, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	adapters := conv.ToIntSlice(lines)
	slices.Sort(adapters)

	diff1 := 0
	diff3 := 0

	current := 0
	for _, adapter := range adapters {
		diff := adapter - current
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}
		current = adapter
	}

	diff3++

	fmt.Println("Part 1", diff1*diff3)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	adapters := conv.ToIntSlice(lines)
	slices.Sort(adapters)

	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	n := len(adapters)
	dp := make([]int64, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		for j := range i {
			if adapters[i]-adapters[j] <= 3 {
				dp[i] += dp[j]
			}
		}
	}

	fmt.Println("Part 2", dp[n-1])
}
