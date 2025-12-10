package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2025, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	rotations := conv.SplitNewline(input)
	position := 50
	zeroPositions := 0
	for _, rotation := range rotations {
		direction := rotation[0]
		steps := conv.MustAtoi(rotation[1:])
		switch direction {
		case 'R':
			position += steps
		case 'L':
			position -= steps
		}
		position = ((position % 100) + 100) % 100
		if position == 0 {
			zeroPositions++
		}
	}
	fmt.Println("Part 1", zeroPositions)
}

func part2(input string) {

	rotations := conv.SplitNewline(input)
	position := 50
	totalHits := 0
	for _, rotation := range rotations {
		direction := rotation[0]
		steps := conv.MustAtoi(rotation[1:])

		var first int
		switch direction {
		case 'R':
			first = (100 - position) % 100
		case 'L':
			first = position % 100
		}
		if first == 0 {
			first = 100
		}

		if first <= steps {
			totalHits += 1 + (steps-first)/100
		}

		switch direction {
		case 'R':
			position += steps
		case 'L':
			position -= steps
		}
		position = ((position % 100) + 100) % 100

	}
	fmt.Println("Part 2", totalHits)

}
