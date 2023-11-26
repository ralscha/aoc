package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	sue := make(map[string]int)
	sue["children"] = 3
	sue["cats"] = 7
	sue["samoyeds"] = 2
	sue["pomeranians"] = 3
	sue["akitas"] = 0
	sue["vizslas"] = 0
	sue["goldfish"] = 5
	sue["trees"] = 3
	sue["cars"] = 2
	sue["perfumes"] = 1

	part1(input, sue)
	part2(input, sue)
}

func part1(input string, sue map[string]int) {
	lines := conv.SplitNewline(input)
	for _, line := range lines {
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

		match := true
		for clue, value := range clues {
			if value != sue[clue] {
				match = false
			}
		}
		if match {
			fmt.Println("Part 1: ", sueNumber)
		}
	}
}

func part2(input string, sue map[string]int) {
	lines := conv.SplitNewline(input)
	for _, line := range lines {
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

		match := true
		for clue, value := range clues {
			switch clue {
			case "cats", "trees":
				if value <= sue[clue] {
					match = false
				}
			case "pomeranians", "goldfish":
				if value >= sue[clue] {
					match = false
				}
			default:
				if value != sue[clue] {
					match = false
				}
			}
		}
		if match {
			fmt.Println("Part 2: ", sueNumber)
		}
	}
}
