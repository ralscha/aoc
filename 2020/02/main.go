package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	valid := 0
	for _, line := range lines {
		var minOccurrences, maxOccurrences int
		var char byte
		var password string
		conv.MustSscanf(line, "%d-%d %c: %s", &minOccurrences, &maxOccurrences, &char, &password)
		count := 0
		for i := range password {
			if password[i] == char {
				count++
			}
		}

		if count >= minOccurrences && count <= maxOccurrences {
			valid++
		}
	}

	fmt.Println("Part 1", valid)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	valid := 0
	for _, line := range lines {
		var pos1, pos2 int
		var char byte
		var password string
		conv.MustSscanf(line, "%d-%d %c: %s", &pos1, &pos2, &char, &password)

		if (password[pos1-1] == char) != (password[pos2-1] == char) {
			valid++
		}
	}

	fmt.Println("Part 2", valid)
}
