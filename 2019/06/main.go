package main

import (
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

	fmt.Println("Part 1:", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	orbits := make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, ")")
		orbits[split[1]] = split[0]
	}

	you := make([]string, 0)
	san := make([]string, 0)

	for k := "YOU"; k != "COM"; k = orbits[k] {
		you = append(you, k)
	}

	for k := "SAN"; k != "COM"; k = orbits[k] {
		san = append(san, k)
	}

	for i, y := range you {
		for j, s := range san {
			if y == s {
				fmt.Println("Part 2:", i+j-2)
				return
			}
		}
	}
}
