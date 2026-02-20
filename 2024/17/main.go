package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strings"
)

type computer struct {
	a, b, c        int
	ip             int
	program        []int
	output         []int
	instructionLen int
}

func newComputer(program []int) *computer {
	return &computer{
		program:        program,
		instructionLen: len(program),
	}
}

func (cmp *computer) reset(a, b, c int) {
	cmp.a = a
	cmp.b = b
	cmp.c = c
	cmp.ip = 0
	cmp.output = cmp.output[:0]
}

func (cmp *computer) run() {
	for cmp.ip < cmp.instructionLen {
		opcode := cmp.program[cmp.ip]
		operand := cmp.program[cmp.ip+1]

		switch opcode {
		case 0:
			cmp.a /= 1 << cmp.getComboValue(operand)
		case 1:
			cmp.b ^= operand
		case 2:
			cmp.b = cmp.getComboValue(operand) % 8
		case 3:
			if cmp.a != 0 {
				cmp.ip = operand
				continue
			}
		case 4:
			cmp.b ^= cmp.c
		case 5:
			cmp.output = append(cmp.output, cmp.getComboValue(operand)%8)
		case 6:
			cmp.b = cmp.a / (1 << cmp.getComboValue(operand))
		case 7:
			cmp.c = cmp.a / (1 << cmp.getComboValue(operand))
		default:
			log.Fatalf("Unknown opcode: %d", opcode)
		}
		cmp.ip += 2
	}
}

func (cmp *computer) getComboValue(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return cmp.a
	case 5:
		return cmp.b
	case 6:
		return cmp.c
	default:
		log.Fatalf("Invalid combo operand: %d", operand)
		return 0
	}
}

func (cmp *computer) GetOutput() string {
	var parts []string
	for _, v := range cmp.output {
		parts = append(parts, fmt.Sprintf("%d", v))
	}
	return strings.Join(parts, ",")
}

func (cmp *computer) outputMatches(target []int) bool {
	if len(cmp.output) != len(target) {
		return false
	}
	for i, v := range cmp.output {
		if v != target[i] {
			return false
		}
	}
	return true
}

func parseInput(input string) (a, b, c int, program []int) {
	lines := strings.Split(input, "\n")

	conv.MustSscanf(lines[0], "Register A: %d", &a)
	conv.MustSscanf(lines[1], "Register B: %d", &b)
	conv.MustSscanf(lines[2], "Register C: %d", &c)

	programStr := strings.TrimPrefix(lines[4], "Program: ")
	programParts := strings.Split(programStr, ",")
	program = make([]int, len(programParts))
	for i, p := range programParts {
		conv.MustSscanf(p, "%d", &program[i])
	}

	return
}

func main() {
	input, err := download.ReadInput(2024, 17)
	if err != nil {
		panic(err)
	}
	part1(input)
	part2(input)
}

func part1(input string) {
	a, b, c, program := parseInput(input)

	computer := newComputer(program)
	computer.a = a
	computer.b = b
	computer.c = c
	computer.run()

	fmt.Println("Part 1", computer.GetOutput())
}

type solution []int

func recombine(nums []int) int {
	result := nums[0]
	d := 10

	for _, c := range nums[1:] {
		result += (c >> 7) << d
		d += 3
	}

	return result
}

// https://github.com/yahya-tamur/advent/blob/main/python/a2024/17.py
func part2(input string) {
	_, b, c, program := parseInput(input)

	steps := make([]int, 1024)
	computer := newComputer(program)
	for a := range 1024 {
		computer.reset(a, b, c)
		computer.run()
		if len(computer.output) > 0 {
			steps[a] = computer.output[0]
		}
	}

	var solutions []solution
	for i := range 1024 {
		if steps[i] == program[0] {
			solutions = append(solutions, solution{i})
		}
	}

	for _, k := range program[1:] {
		var newSolutions []solution
		for _, s := range solutions {
			current := s[len(s)-1] >> 3
			for i := range 8 {
				if steps[(i<<7)+current] == k {
					newSol := make(solution, len(s)+1)
					copy(newSol, s)
					newSol[len(s)] = (i << 7) + current
					newSolutions = append(newSolutions, newSol)
				}
			}
		}
		solutions = newSolutions
	}

	minResult := math.MaxInt
	for _, s := range solutions {
		result := recombine(s)
		computer.reset(result, b, c)
		computer.run()
		if computer.outputMatches(program) {
			if result < minResult {
				minResult = result
			}
		}
	}

	fmt.Println("Part 2", minResult)
}
