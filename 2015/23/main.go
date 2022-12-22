package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/23/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 0)
	part1and2(input, 1)
}

type instruction struct {
	opcode string
	reg    string
	offset int
}

func part1and2(input string, startA int) {
	lines := conv.SplitNewline(input)
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}
	a, b := startA, 0
	index := 0
	for index < len(instructions) {
		instr := instructions[index]
		switch instr.opcode {
		case "hlf":
			if instr.reg == "a" {
				a /= 2
			} else {
				b /= 2
			}
			index++
		case "tpl":
			if instr.reg == "a" {
				a *= 3
			} else {
				b *= 3
			}
			index++
		case "inc":
			if instr.reg == "a" {
				a++
			} else {
				b++
			}
			index++
		case "jmp":
			index += instr.offset
		case "jie":
			if (instr.reg == "a" && a%2 == 0) || (instr.reg == "b" && b%2 == 0) {
				index += instr.offset
			} else {
				index++
			}
		case "jio":
			if (instr.reg == "a" && a == 1) || (instr.reg == "b" && b == 1) {
				index += instr.offset
			} else {
				index++
			}
		}
	}

	fmt.Println(b)
}

func parseInstruction(line string) instruction {
	splitted := strings.Fields(line)
	opcode := splitted[0]
	reg := splitted[1]
	if opcode == "jmp" {
		return instruction{
			opcode: opcode,
			offset: conv.MustAtoi(splitted[1]),
		}
	} else if len(splitted) == 3 {
		reg = strings.TrimSuffix(reg, ",")
		offset := conv.MustAtoi(splitted[2])
		return instruction{opcode, reg, offset}
	}
	return instruction{opcode, reg, 0}
}
