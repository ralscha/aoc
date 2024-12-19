package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
)

type disc struct {
	positions int
	start     int
}

func parseInput(input string) []disc {
	lines := conv.SplitNewline(input)
	discs := make([]disc, len(lines))
	regex := regexp.MustCompile(`Disc #\d+ has (\d+) positions; at time=0, it is at position (\d+).`)

	for i, line := range lines {
		matches := regex.FindStringSubmatch(line)
		positions := conv.MustAtoi(matches[1])
		start := conv.MustAtoi(matches[2])
		discs[i] = disc{positions: positions, start: start}
	}
	return discs
}

func main() {
	input, err := download.ReadInput(2016, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	discs := parseInput(input)
	time := 0
	for {
		valid := true
		for i, disc := range discs {
			if (disc.start+time+i+1)%disc.positions != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println("Part 1", time)
			return
		}
		time++
	}
}

func part2(input string) {
	discs := parseInput(input)
	discs = append(discs, disc{positions: 11, start: 0})
	time := 0
	for {
		valid := true
		for i, disc := range discs {
			if (disc.start+time+i+1)%disc.positions != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println("Part 2", time)
			return
		}
		time++
	}
}
