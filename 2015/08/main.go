package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2015/08/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var total, totalInMemory int

	for _, line := range lines {
		total += len(line)
		for i := 1; i < len(line)-1; i++ {
			if line[i] == '\\' {
				if i+1 < len(line) && (line[i+1] == '"' || line[i+1] == '\\') {
					totalInMemory++
					i++
				} else if i+1 < len(line) && line[i+1] == 'x' {
					totalInMemory++
					i += 3
				} else {
					totalInMemory++
				}
			} else {
				totalInMemory++
			}
		}
	}

	fmt.Println(total - totalInMemory)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var total, totalEncoded int

	for _, line := range lines {
		total += len(line)

		totalEncoded += len(line) + 4
		for i := 1; i < len(line)-1; i++ {
			if line[i] == '\\' || line[i] == '"' {
				totalEncoded++
			}
		}
	}

	fmt.Println(totalEncoded - total)
}
