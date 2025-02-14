package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"

	"aoc/2019/intcomputer"
)

func main() {
	input, err := download.ReadInput(2019, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	if err := computer.AddInput(1); err != nil {
		log.Fatalf("adding input failed: %v", err)
	}

	var lastOutput int
	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			lastOutput = result.Value
		case intcomputer.SignalEnd:
			fmt.Println("Part 1", lastOutput)
			return
		case intcomputer.SignalInput:
			log.Fatal("unexpected input request")
		}
	}
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	computer := intcomputer.NewIntcodeComputer(program)
	if err := computer.AddInput(2); err != nil {
		log.Fatalf("adding input failed: %v", err)
	}

	var lastOutput int
	for {
		result, err := computer.Run()
		if err != nil {
			log.Fatalf("running program failed: %v", err)
		}

		switch result.Signal {
		case intcomputer.SignalOutput:
			lastOutput = result.Value
		case intcomputer.SignalEnd:
			fmt.Println("Part 2", lastOutput)
			return
		case intcomputer.SignalInput:
			log.Fatal("unexpected input request")
		}
	}
}
