package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
)

func main() {
	input, err := download.ReadInput(2021, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var ones [12]int32
	var zeros [12]int32

	for _, line := range lines {
		for ix, c := range line {
			if c-'0' == 0 {
				zeros[ix]++
			} else {
				ones[ix]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for ix := range ones {
		if ones[ix] > zeros[ix] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaNumber, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", gamma, err)
	}
	epsilonNumber, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", epsilon, err)
	}
	fmt.Println("Part 1", gammaNumber*epsilonNumber)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	input1 := make([]string, len(lines))
	copy(input1, lines)
	input2 := make([]string, len(lines))
	copy(input2, lines)

	for i := 0; i < 12; i++ {
		if len(input1) > 1 {
			onesWinner := count(input1, i, true)
			input1 = filter(input1, onesWinner, i)
		}

		if len(input2) > 1 {
			zerosWinner := count(input2, i, false)
			input2 = filter(input2, zerosWinner, i)
		}
	}

	if len(input1) != 1 || len(input2) != 1 {
		log.Fatalf("Wrong calculation")
	}

	oxygen, err := strconv.ParseInt(input1[0], 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", input1[0], err)
	}
	co2, err := strconv.ParseInt(input2[0], 2, 64)
	if err != nil {
		log.Fatalf("ParseInt failed: %s %v", input2[0], err)
	}

	fmt.Println("Part 2", oxygen*co2)
}

func count(lines []string, pos int, one bool) int {
	var ones, zeros int
	for _, line := range lines {
		if line[pos] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if one {
		if zeros > ones {
			return 0
		}
		return 1
	} else {
		if ones < zeros {
			return 1
		}
		return 0
	}
}

func filter(lines []string, winner, pos int) []string {
	var filtered []string
	for _, line := range lines {
		if int(line[pos]-'0') == winner {
			filtered = append(filtered, line)
		}
	}
	return filtered
}
