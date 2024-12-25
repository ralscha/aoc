package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"

	"aoc/2019/intcomputer"
)

func main() {
	input, err := download.ReadInput(2019, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	computer.Input = 1
	for !computer.Halted {
		computer.Run()
		if computer.Output != 0 {
			fmt.Println("Part 1", computer.Output)
			return
		}
	}
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	computer.Input = 5
	computer.Run()
	fmt.Println("Part 2", computer.Output)
}
