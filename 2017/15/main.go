package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2017, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var startA, startB int
	conv.MustSscanf(lines[0], "Generator A starts with %d", &startA)
	conv.MustSscanf(lines[1], "Generator B starts with %d", &startB)

	factorA := 16807
	factorB := 48271
	divisor := 2147483647

	count := 0
	valueA := startA
	valueB := startB

	for range 40000000 {
		valueA = (valueA * factorA) % divisor
		valueB = (valueB * factorB) % divisor

		if valueA&0xFFFF == valueB&0xFFFF {
			count++
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var startA, startB int
	conv.MustSscanf(lines[0], "Generator A starts with %d", &startA)
	conv.MustSscanf(lines[1], "Generator B starts with %d", &startB)

	factorA := 16807
	factorB := 48271
	divisor := 2147483647

	count := 0
	valueA := startA
	valueB := startB

	for range 5000000 {
		for {
			valueA = (valueA * factorA) % divisor
			if valueA%4 == 0 {
				break
			}
		}
		for {
			valueB = (valueB * factorB) % divisor
			if valueB%8 == 0 {
				break
			}
		}

		if valueA&0xFFFF == valueB&0xFFFF {
			count++
		}
	}

	fmt.Println("Part 2", count)
}
