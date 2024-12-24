package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func transform(subjectNumber int, loopSize int) int {
	value := 1
	for range loopSize {
		value *= subjectNumber
		value %= 20201227
	}
	return value
}

func findLoopSize(publicKey int) int {
	subjectNumber := 7
	value := 1
	loopSize := 0
	for {
		loopSize++
		value *= subjectNumber
		value %= 20201227
		if value == publicKey {
			return loopSize
		}
	}
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	cardPublicKey := conv.MustAtoi(lines[0])
	doorPublicKey := conv.MustAtoi(lines[1])
	cardLoopSize := findLoopSize(cardPublicKey)
	encryptionKey := transform(doorPublicKey, cardLoopSize)
	fmt.Println("Part 1", encryptionKey)
}
