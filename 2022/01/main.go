package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2022, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	var elves []int
	totalCalories := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elves = append(elves, totalCalories)
			totalCalories = 0
		} else {
			calories := conv.MustAtoi(line)
			totalCalories += calories
		}
	}
	slices.SortFunc(elves, func(a, b int) int {
		return b - a
	})

	fmt.Println("Part 1", elves[0])

	sum := 0
	for i := range 3 {
		sum += elves[i]
	}
	fmt.Println("Part 2", sum)
}
