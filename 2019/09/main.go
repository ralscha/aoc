package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2019, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type intCodeComputer struct {
	memory       map[int64]int64
	ip           int64
	relativeBase int64
}

func newIntcodeComputer(program []int64) *intCodeComputer {
	memory := make(map[int64]int64)
	for i, v := range program {
		memory[int64(i)] = v
	}
	return &intCodeComputer{memory: memory}
}

func (comp *intCodeComputer) getValue(addr int64) int64 {
	return comp.memory[addr]
}

func (comp *intCodeComputer) setValue(addr int64, value int64) {
	comp.memory[addr] = value
}

func decodeInstruction(instruction int64) (int64, [3]int64) {
	opcode := instruction % 100
	modes := [3]int64{
		(instruction / 100) % 10,
		(instruction / 1000) % 10,
		(instruction / 10000) % 10,
	}
	return opcode, modes
}

func (comp *intCodeComputer) getParameter(mode int64, offset int64) int64 {
	switch mode {
	case 0:
		return comp.getValue(comp.getValue(comp.ip + offset))
	case 1:
		return comp.getValue(comp.ip + offset)
	case 2:
		return comp.getValue(comp.relativeBase + comp.getValue(comp.ip+offset))
	default:
		log.Fatalf("Invalid parameter mode: %d", mode)
		return 0
	}
}

func (comp *intCodeComputer) getAddress(mode int64, offset int64) int64 {
	switch mode {
	case 0:
		return comp.getValue(comp.ip + offset)
	case 2:
		return comp.relativeBase + comp.getValue(comp.ip+offset)
	default:
		log.Fatalf("Invalid address mode: %d", mode)
		return 0
	}
}

func (comp *intCodeComputer) Run(input int64) (output int64) {
	for {
		opcode, modes := decodeInstruction(comp.getValue(comp.ip))

		switch opcode {
		case 1: // Addition
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			comp.setValue(dst, a+b)
			comp.ip += 4
		case 2: // Multiplication
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			comp.setValue(dst, a*b)
			comp.ip += 4
		case 3: // Input
			dst := comp.getAddress(modes[0], 1)
			comp.setValue(dst, input)
			comp.ip += 2
		case 4: // Output
			a := comp.getParameter(modes[0], 1)
			output = a
			comp.ip += 2
		case 5: // Jump-if-true
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			if a != 0 {
				comp.ip = b
			} else {
				comp.ip += 3
			}
		case 6: // Jump-if-false
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			if a == 0 {
				comp.ip = b
			} else {
				comp.ip += 3
			}
		case 7: // Less than
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			if a < b {
				comp.setValue(dst, 1)
			} else {
				comp.setValue(dst, 0)
			}
			comp.ip += 4
		case 8: // Equals
			a := comp.getParameter(modes[0], 1)
			b := comp.getParameter(modes[1], 2)
			dst := comp.getAddress(modes[2], 3)
			if a == b {
				comp.setValue(dst, 1)
			} else {
				comp.setValue(dst, 0)
			}
			comp.ip += 4
		case 9: // Adjust relative base
			a := comp.getParameter(modes[0], 1)
			comp.relativeBase += a
			comp.ip += 2
		case 99: // Halt
			return output
		default:
			log.Fatalf("Unknown opcode: %d", opcode)
		}
	}
}

func part1(input string) {
	code := conv.ToInt64SliceComma(input)
	computer := newIntcodeComputer(code)
	output := computer.Run(1)
	fmt.Println("Part 1", output)
}

func part2(input string) {
	code := conv.ToInt64SliceComma(input)
	computer := newIntcodeComputer(code)
	output := computer.Run(2)
	fmt.Println("Part 2", output)
}
