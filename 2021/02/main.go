package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	horizontalPosition := 0
	depth := 0

	for _, line := range lines {
		split := strings.Fields(line)
		num := conv.MustAtoi(split[1])
		switch split[0] {
		case "down":
			depth += num
		case "up":
			depth -= num
		case "forward":
			horizontalPosition += num
		}
	}

	fmt.Println("Part 1", depth*horizontalPosition)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		split := strings.Fields(line)
		num := conv.MustAtoi(split[1])
		switch split[0] {
		case "down":
			aim += num
		case "up":
			aim -= num
		case "forward":
			horizontalPosition += num
			depth += aim * num
		}
	}

	fmt.Println("Part 2", depth*horizontalPosition)
}
