package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	orbits := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ")")
		orbits[split[1]] = split[0]
	}

	count := 0
	for k := range orbits {
		count++
		for orbits[k] != "COM" {
			count++
			k = orbits[k]
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	orbits := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ")")
		orbits[split[1]] = split[0]
	}

	youPath := container.NewSet[string]()
	sanPath := container.NewSet[string]()

	// Build path from YOU to COM
	for k := "YOU"; k != "COM"; k = orbits[k] {
		youPath.Add(k)
	}

	// Build path from SAN to COM
	for k := "SAN"; k != "COM"; k = orbits[k] {
		sanPath.Add(k)
	}

	// Find first common ancestor
	minTransfers := len(orbits)
	for k := "YOU"; k != "COM"; k = orbits[k] {
		if sanPath.Contains(k) {
			youSteps := 0
			for temp := "YOU"; temp != k; temp = orbits[temp] {
				youSteps++
			}
			sanSteps := 0
			for temp := "SAN"; temp != k; temp = orbits[temp] {
				sanSteps++
			}
			transfers := youSteps + sanSteps - 2 // subtract 2 because we don't count YOU and SAN
			if transfers < minTransfers {
				minTransfers = transfers
			}
		}
	}

	fmt.Println("Part 2", minTransfers)
}
