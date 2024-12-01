package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var leftList []int
	var rightList []int
	for _, line := range lines {
		fields := strings.Fields(line)
		leftList = append(leftList, conv.MustAtoi(fields[0]))
		rightList = append(rightList, conv.MustAtoi(fields[1]))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += mathx.Abs(leftList[i] - rightList[i])
	}
	fmt.Println("Part 1", totalDistance)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var leftList []int
	var rightList []int
	for _, line := range lines {
		fields := strings.Fields(line)
		leftList = append(leftList, conv.MustAtoi(fields[0]))
		rightList = append(rightList, conv.MustAtoi(fields[1]))
	}

	totalSimilarity := 0
	for i := 0; i < len(leftList); i++ {
		totalSimilarity += count(rightList, leftList[i]) * leftList[i]
	}
	fmt.Println("Part 2", totalSimilarity)
}

func count(list []int, target int) int {
	count := 0
	for _, n := range list {
		if n == target {
			count++
		}
	}
	return count
}
