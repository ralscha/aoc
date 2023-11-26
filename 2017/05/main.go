package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2017, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1(conv.ToIntSlice(lines))
	part2(conv.ToIntSlice(lines))
}

func part1(instructions []int) {
	currentInstruction := 0
	steps := 0
	for currentInstruction >= 0 && currentInstruction < len(instructions) {
		steps++
		newCurrentInstruction := currentInstruction + instructions[currentInstruction]
		instructions[currentInstruction]++
		currentInstruction = newCurrentInstruction
	}
	fmt.Println(steps)
}

func part2(instructions []int) {
	currentInstruction := 0
	steps := 0
	for currentInstruction >= 0 && currentInstruction < len(instructions) {
		steps++
		newCurrentInstruction := currentInstruction + instructions[currentInstruction]
		if instructions[currentInstruction] >= 3 {
			instructions[currentInstruction]--
		} else {
			instructions[currentInstruction]++
		}
		currentInstruction = newCurrentInstruction
	}
	fmt.Println(steps)
}
