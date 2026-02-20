package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2022, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	total := 0
	for line := range strings.SplitSeq(input, "\n") {
		half := len(line) / 2
		first, second := line[:half], line[half:]
		for _, c := range first {
			if strings.Contains(second, string(c)) {
				total += int(getPriority(c))
				break
			}
		}
	}
	fmt.Println(total)
}

func part2(input string) {
	total := 0
	splittedInput := strings.Split(input, "\n")
	for i := 0; i < len(splittedInput); i += 3 {
		if i+2 >= len(splittedInput) {
			break
		}
		first, second, third := splittedInput[i], splittedInput[i+1], splittedInput[i+2]
		for _, c := range first {
			if strings.Contains(second, string(c)) && strings.Contains(third, string(c)) {
				total += int(getPriority(c))
				break
			}
		}
	}
	fmt.Println(total)
}

func getPriority(input rune) rune {
	/*
		Lowercase item types a through z have priorities 1 through 26.
		Uppercase item types A through Z have priorities 27 through 52.
	*/
	if input >= 65 && input <= 90 {
		return input - 64 + 26
	} else {
		return input - 96
	}
}
