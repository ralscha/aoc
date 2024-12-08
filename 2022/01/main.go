package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"sort"
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

	sort.Slice(elves, func(i, j int) bool {
		return elves[i] > elves[j]
	})

	fmt.Println("Part 1", elves[0])

	sum := 0
	for i := 0; i < 3; i++ {
		sum += elves[i]
	}
	fmt.Println("Part 2", sum)
}
