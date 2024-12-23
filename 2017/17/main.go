package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2017, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	steps := conv.MustAtoi(input)

	buffer := []int{0}
	currentPosition := 0

	for i := 1; i <= 2017; i++ {
		nextPosition := (currentPosition + steps) % len(buffer)
		buffer = append(buffer[:nextPosition+1], buffer[nextPosition:]...)
		buffer[nextPosition+1] = i
		currentPosition = nextPosition + 1
	}

	var result int
	for i := range len(buffer) {
		if buffer[i] == 2017 {
			result = buffer[(i+1)%len(buffer)]
			break
		}
	}

	fmt.Println("Part 1", result)
}

func part2(input string) {
	steps := conv.MustAtoi(input)

	length := 1
	currentPosition := 0
	valueAfterZero := 0

	for i := 1; i <= 50000000; i++ {
		nextPosition := (currentPosition + steps) % length
		if nextPosition == 0 {
			valueAfterZero = i
		}
		currentPosition = nextPosition + 1
		length++
	}

	fmt.Println("Part 2", valueAfterZero)
}
