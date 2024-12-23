package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"maps"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	// Initial Sue's properties
	realSue := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	part1(input, realSue)
	part2(input, realSue)
}

func parseClues(line string) (int, map[string]int) {
	splitted := strings.Fields(line)
	sueNumber := conv.MustAtoi(splitted[1][:len(splitted[1])-1])
	clues := make(map[string]int)

	for i := 2; i < len(splitted); i += 2 {
		clue := splitted[i][:len(splitted[i])-1]
		no := splitted[i+1]
		if no[len(no)-1] == ',' {
			no = no[:len(no)-1]
		}
		clues[clue] = conv.MustAtoi(no)
	}

	return sueNumber, clues
}

func part1(input string, realSue map[string]int) {
	lines := conv.SplitNewline(input)

	for _, line := range lines {
		sueNumber, clues := parseClues(line)

		match := true
		for key, value := range maps.All(clues) {
			clue, value := key, value
			if value != realSue[clue] {
				match = false
				break
			}
		}

		if match {
			fmt.Println("Part 1", sueNumber)
		}
	}
}

func part2(input string, realSue map[string]int) {
	lines := conv.SplitNewline(input)

	for _, line := range lines {
		sueNumber, clues := parseClues(line)

		match := true
		for key, value := range maps.All(clues) {
			clue, value := key, value
			switch clue {
			case "cats", "trees":
				if value <= realSue[clue] {
					match = false
				}
			case "pomeranians", "goldfish":
				if value >= realSue[clue] {
					match = false
				}
			default:
				if value != realSue[clue] {
					match = false
				}
			}
			if !match {
				break
			}
		}

		if match {
			fmt.Println("Part 2", sueNumber)
		}
	}
}
