package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"bufio"
	"fmt"
	"log"
	"slices"
	"strings"
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
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var crabs []int
	for _, s := range splitted {
		n := conv.MustAtoi(s)
		crabs = append(crabs, n)
	}

	leastFuel := -1
	leastPos := -1

	minCrabs, maxCrabs := findMinMax(crabs)
	for p := minCrabs; p <= maxCrabs; p++ {
		fuel := 0
		for _, c := range crabs {
			fuel += abs(c - p)
		}
		if fuel < leastFuel || leastFuel == -1 {
			leastFuel = fuel
			leastPos = p
		}
	}

	fmt.Println("Result: ", leastPos, leastFuel)
}

func part2(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")

	var crabs []int
	for _, s := range splitted {
		n := conv.MustAtoi(s)
		crabs = append(crabs, n)
	}

	leastFuel := -1
	leastPos := -1

	minCrabs, maxCrabs := findMinMax(crabs)
	for p := minCrabs; p <= maxCrabs; p++ {
		fuel := 0
		for _, c := range crabs {
			diff := abs(c - p)
			fuel += diff * (diff + 1) / 2
		}
		if fuel < leastFuel || leastFuel == -1 {
			leastFuel = fuel
			leastPos = p
		}
	}

	fmt.Println("Result: ", leastPos, leastFuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findMinMax(a []int) (int, int) {
	return slices.Min(a), slices.Max(a)
}
