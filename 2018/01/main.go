package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2018/01/input.txt"
	input, err := download.ReadInput(inputFile, 2018, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	frequency := 0
	for _, line := range lines {
		frequency += conv.MustAtoi(line)
	}
	fmt.Println(frequency)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	frequency := 0
	seen := container.NewSet[int]()
	for {
		for _, line := range lines {
			frequency += conv.MustAtoi(line)
			if seen.Contains(frequency) {
				fmt.Println(frequency)
				return
			}
			seen.Add(frequency)
		}
	}

}
