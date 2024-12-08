package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2021, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	crabs := conv.ToIntSliceComma(input)

	minCrabs, maxCrabs := math.MaxInt, math.MinInt
	for _, c := range crabs {
		if c < minCrabs {
			minCrabs = c
		}
		if c > maxCrabs {
			maxCrabs = c
		}
	}

	leastFuel := math.MaxInt

	for p := minCrabs; p <= maxCrabs; p++ {
		fuel := 0
		for _, c := range crabs {
			fuel += mathx.Abs(c - p)
		}
		if fuel < leastFuel {
			leastFuel = fuel
		}
	}

	fmt.Println("Part 1", leastFuel)
}

func part2(input string) {
	crabs := conv.ToIntSliceComma(input)

	minCrabs, maxCrabs := math.MaxInt, math.MinInt
	for _, c := range crabs {
		if c < minCrabs {
			minCrabs = c
		}
		if c > maxCrabs {
			maxCrabs = c
		}
	}

	leastFuel := math.MaxInt

	for p := minCrabs; p <= maxCrabs; p++ {
		fuel := 0
		for _, c := range crabs {
			diff := mathx.Abs(c - p)
			// Sum of arithmetic sequence: n * (n + 1) / 2
			fuel += diff * (diff + 1) / 2
		}
		if fuel < leastFuel {
			leastFuel = fuel
		}
	}

	fmt.Println("Part 2", leastFuel)
}
