package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	fuelNeeded := 0
	for _, line := range lines {
		mass := conv.MustAtoi(line)
		mass = mass / 3
		mass -= 2
		fuelNeeded += mass
	}
	fmt.Println(fuelNeeded)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	fuelNeeded := 0
	for _, line := range lines {
		mass := conv.MustAtoi(line)
		for {
			mass = mass / 3
			mass -= 2
			if mass <= 0 {
				break
			}
			fuelNeeded += mass
		}
	}
	fmt.Println(fuelNeeded)
}
