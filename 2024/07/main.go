package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func compute(operands []int, operators []string) int {
	result := operands[0]
	for i, op := range operators {
		switch op {
		case "+":
			result += operands[i+1]
		case "*":
			result *= operands[i+1]
		case "||":
			result = conv.MustAtoi(fmt.Sprintf("%d%d", result, operands[i+1]))
		}
	}
	return result
}

func equalsToTarget(target int, numbers []int, includeConcat bool) bool {
	n := len(numbers) - 1
	var operators []string
	if includeConcat {
		operators = []string{"+", "*", "||"}
	} else {
		operators = []string{"+", "*"}
	}
	combinations := mathx.CartesianProductSelf(n, operators)
	for _, operators := range combinations {
		if compute(numbers, operators) == target {
			return true
		}
	}
	return false
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	result := 0

	for _, input := range lines {
		parts := strings.Split(input, ":")
		target := conv.MustAtoi(strings.TrimSpace(parts[0]))
		numbers := conv.ToIntSlice(strings.Fields(strings.TrimSpace(parts[1])))
		if equalsToTarget(target, numbers, false) {
			result += target
		}
	}

	fmt.Println("Part 1", result)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	result := 0
	for _, input := range lines {
		parts := strings.Split(input, ":")
		target := conv.MustAtoi(strings.TrimSpace(parts[0]))
		numbers := conv.ToIntSlice(strings.Fields(strings.TrimSpace(parts[1])))
		if equalsToTarget(target, numbers, true) {
			result += target
		}
	}

	fmt.Println("Part 2", result)
}
