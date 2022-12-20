package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2017/01/input.txt"
	input, err := download.ReadInput(inputFile, 2017, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	input = input[:len(input)-1]
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			sum += int(input[i] - '0')
		}
	}

	fmt.Println(sum)
}

func part2(input string) {
	input = input[:len(input)-1]
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+len(input)/2)%len(input)] {
			sum += int(input[i] - '0')
		}
	}

	fmt.Println(sum)
}
