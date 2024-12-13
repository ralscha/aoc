package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/stringutil"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

func main() {
	input, err := download.ReadInput(2023, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		var digits []string
		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, string(c))
			}
		}
		if len(digits) > 0 {
			sum += conv.MustAtoi(digits[0] + digits[len(digits)-1])
		}
	}
	fmt.Println(sum)
}

func part2(input string) {
	numbers := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	lines := conv.SplitNewline(input)
	sum := 0
	for _, line := range lines {
		// Find all digit positions
		var digitPositions []struct {
			pos   int
			value string
		}

		// Find numeric digits
		for i, c := range line {
			if unicode.IsDigit(c) {
				digitPositions = append(digitPositions, struct {
					pos   int
					value string
				}{i, string(c)})
			}
		}

		// Find word digits
		for i, word := range numbers {
			// Use stringutil.FindAllOccurrences to find all instances of each number word
			positions := stringutil.FindAllOccurrences(line, word)
			for _, pos := range positions {
				digitPositions = append(digitPositions, struct {
					pos   int
					value string
				}{pos, strconv.Itoa(i)})
			}
		}

		if len(digitPositions) > 0 {
			// Sort positions by their index to find first and last
			firstPos := digitPositions[0]
			lastPos := digitPositions[0]
			for _, pos := range digitPositions {
				if pos.pos < firstPos.pos {
					firstPos = pos
				}
				if pos.pos > lastPos.pos {
					lastPos = pos
				}
			}

			sum += conv.MustAtoi(firstPos.value + lastPos.value)
		}
	}
	fmt.Println(sum)
}
