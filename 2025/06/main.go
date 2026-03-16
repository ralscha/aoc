package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/slicesx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	opRow := lines[len(lines)-1]
	numRows := lines[:len(lines)-1]

	numbers := make([][]int, len(numRows))
	for i, line := range numRows {
		numberStrings := strings.Fields(line)
		nums := make([]int, len(numberStrings))
		for n, s := range numberStrings {
			nums[n] = conv.MustAtoi(s)
		}
		numbers[i] = nums
	}

	operations := strings.Fields(opRow)

	total := 0
	for col := range numbers[0] {
		var nums []int
		for row := range numbers {
			nums = append(nums, numbers[row][col])
		}
		op := operations[col]
		total += applyOp(op, nums)
	}
	fmt.Println("Part 1", total)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	opRow := lines[len(lines)-1]
	numRows := lines[:len(lines)-1]

	maxCol := 0
	for _, line := range numRows {
		if len(line) > maxCol {
			maxCol = len(line)
		}
	}

	var numStrings []string
	currentIndex := 0
	op := opRow[currentIndex]
	total := 0

	for {
		var numString strings.Builder
		for _, row := range numRows {
			ch := row[currentIndex]
			if ch != ' ' {
				numString.WriteString(string(ch))
			}
		}
		if numString.String() != "" {
			numStrings = append(numStrings, numString.String())
		}

		currentIndex++
		if currentIndex >= maxCol || allSpaces(numRows, currentIndex) {
			nums := make([]int, len(numStrings))
			for i, s := range numStrings {
				nums[i] = conv.MustAtoi(s)
			}
			total += applyOp(string(op), nums)
			currentIndex++
			if currentIndex >= maxCol {
				break
			}
			op = opRow[currentIndex]
			numStrings = nil
		}
	}

	fmt.Println("Part 2", total)
}

func applyOp(op string, nums []int) int {
	if op == "*" {
		result := 1
		for _, n := range nums {
			result *= n
		}
		return result
	}
	return slicesx.Sum(nums)
}

func allSpaces(s []string, col int) bool {
	for _, str := range s {
		if str[col] != ' ' {
			return false
		}
	}
	return true
}
