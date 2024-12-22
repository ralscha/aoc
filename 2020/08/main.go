package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type instruction struct {
	op  string
	arg int
}

func parseInput(input string) []instruction {
	lines := conv.SplitNewline(input)
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		parts := strings.Fields(line)
		arg := conv.MustAtoi(parts[1])
		instructions[i] = instruction{op: parts[0], arg: arg}
	}
	return instructions
}

func part1(input string) {
	instructions := parseInput(input)
	accumulator := 0
	pc := 0
	executed := container.NewSet[int]()

	for {
		if executed.Contains(pc) {
			fmt.Println("Part 1", accumulator)
			return
		}
		executed.Add(pc)
		instr := instructions[pc]
		switch instr.op {
		case "acc":
			accumulator += instr.arg
			pc++
		case "jmp":
			pc += instr.arg
		case "nop":
			pc++
		}
	}
}

func run(instructions []instruction) (int, bool) {
	accumulator := 0
	pc := 0
	executed := container.NewSet[int]()

	for pc < len(instructions) {
		if executed.Contains(pc) {
			return accumulator, false
		}
		executed.Add(pc)
		instr := instructions[pc]
		switch instr.op {
		case "acc":
			accumulator += instr.arg
			pc++
		case "jmp":
			pc += instr.arg
		case "nop":
			pc++
		}
	}
	return accumulator, true
}

func part2(input string) {
	originalInstructions := parseInput(input)

	for i := range originalInstructions {
		modifiedInstructions := make([]instruction, len(originalInstructions))
		copy(modifiedInstructions, originalInstructions)

		if modifiedInstructions[i].op == "jmp" {
			modifiedInstructions[i].op = "nop"
			if acc, terminated := run(modifiedInstructions); terminated {
				fmt.Println("Part 2", acc)
				return
			}
		} else if modifiedInstructions[i].op == "nop" {
			modifiedInstructions[i].op = "jmp"
			if acc, terminated := run(modifiedInstructions); terminated {
				fmt.Println("Part 2", acc)
				return
			}
		}
	}
}
