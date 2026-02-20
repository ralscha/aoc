package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var crateRegex = regexp.MustCompile(`\[([A-Z])]`)

func main() {
	input, err := download.ReadInput(2022, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	var stacks [][]string

	for line := range strings.SplitSeq(input, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "1") {
			break
		}
		cratesIndex := crateRegex.FindAllStringIndex(line, -1)
		for _, crateIndex := range cratesIndex {
			createNo := crateIndex[0] / 4
			if len(stacks) <= createNo {
				for i := len(stacks); i <= createNo; i++ {
					stacks = append(stacks, []string{})
				}
			}
			crate := line[crateIndex[0]+1 : crateIndex[1]-1]
			stacks[createNo] = append(stacks[createNo], crate)
		}
	}

	for line := range strings.SplitSeq(input, "\n") {
		if !strings.HasPrefix(line, "move") {
			continue
		}
		move, from, to := parseLine(line)
		from -= 1
		to -= 1
		for range move {
			stacks[to] = append([]string{stacks[from][0]}, stacks[to]...)
			stacks[from] = stacks[from][1:]
		}
	}

	var top strings.Builder
	for _, stack := range stacks {
		if len(stack) > 0 {
			top.WriteString(stack[0])
		}
	}
	fmt.Println(top.String())
}

func part2(input string) {
	var stacks [][]string

	for line := range strings.SplitSeq(input, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "1") {
			break
		}
		cratesIndex := crateRegex.FindAllStringIndex(line, -1)
		for _, crateIndex := range cratesIndex {
			createNo := crateIndex[0] / 4
			if len(stacks) <= createNo {
				for i := len(stacks); i <= createNo; i++ {
					stacks = append(stacks, []string{})
				}
			}
			crate := line[crateIndex[0]+1 : crateIndex[1]-1]
			stacks[createNo] = append(stacks[createNo], crate)
		}
	}

	for line := range strings.SplitSeq(input, "\n") {
		if !strings.HasPrefix(line, "move") {
			continue
		}
		move, from, to := parseLine(line)
		from -= 1
		to -= 1

		moveCrates := stacks[from][0:move]
		moveCratesCopy := make([]string, len(moveCrates))
		copy(moveCratesCopy, moveCrates)

		stacks[to] = append(moveCratesCopy, stacks[to]...)
		stacks[from] = stacks[from][move:]
	}

	var top strings.Builder
	for _, stack := range stacks {
		if len(stack) > 0 {
			top.WriteString(stack[0])
		}
	}
	fmt.Println(top.String())
}

func parseLine(line string) (int, int, int) {
	// move 6 from 4 to 3
	splitted := strings.Fields(line)
	return conv.MustAtoi(splitted[1]), conv.MustAtoi(splitted[3]), conv.MustAtoi(splitted[5])
}
