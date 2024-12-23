package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	earliestTimestamp := conv.MustAtoi(lines[0])
	busIDsStr := strings.Split(lines[1], ",")
	var busIDs []int
	for _, idStr := range busIDsStr {
		if idStr != "x" {
			busIDs = append(busIDs, conv.MustAtoi(idStr))
		}
	}

	minWaitTime := -1
	bestBusID := -1

	for _, busID := range busIDs {
		wait := busID - (earliestTimestamp % busID)
		if minWaitTime == -1 || wait < minWaitTime {
			minWaitTime = wait
			bestBusID = busID
		}
	}

	fmt.Println("Part 1", bestBusID*minWaitTime)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	parts := strings.Split(lines[1], ",")

	type constraint struct {
		remainder int
		mod       int
	}

	var constraints []constraint
	for i, part := range parts {
		if part != "x" {
			id := conv.MustAtoi(part)
			constraints = append(constraints, constraint{remainder: i, mod: id})
		}
	}

	timestamp := 0
	lcm := 1

	for _, c := range constraints {
		for {
			if (timestamp+c.remainder)%c.mod == 0 {
				lcm *= c.mod
				break
			}
			timestamp += lcm
		}
	}

	fmt.Println("Part 2", timestamp)
}
