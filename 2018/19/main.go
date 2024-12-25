package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type instruction struct {
	opcode string
	args   []int
}

func parseInput(input string) (int, []instruction) {
	lines := conv.SplitNewline(input)
	var ipBind int
	conv.MustSscanf(lines[0], "#ip %d", &ipBind)

	instructions := make([]instruction, len(lines)-1)
	for i, line := range lines[1:] {
		parts := strings.Split(line, " ")
		opcode := parts[0]
		args := make([]int, len(parts)-1)
		for j := 1; j < len(parts); j++ {
			args[j-1] = conv.MustAtoi(parts[j])
		}
		instructions[i] = instruction{opcode: opcode, args: args}
	}
	return ipBind, instructions
}

func execute(registers []int, instruction instruction) []int {
	registersCopy := make([]int, len(registers))
	copy(registersCopy, registers)

	a, b, c := -1, -1, -1
	if len(instruction.args) >= 1 {
		a = instruction.args[0]
	}
	if len(instruction.args) >= 2 {
		b = instruction.args[1]
	}
	if len(instruction.args) >= 3 {
		c = instruction.args[2]
	}

	switch instruction.opcode {
	case "addr":
		registersCopy[c] = registers[a] + registers[b]
	case "addi":
		registersCopy[c] = registers[a] + b
	case "mulr":
		registersCopy[c] = registers[a] * registers[b]
	case "muli":
		registersCopy[c] = registers[a] * b
	case "banr":
		registersCopy[c] = registers[a] & registers[b]
	case "bani":
		registersCopy[c] = registers[a] & b
	case "borr":
		registersCopy[c] = registers[a] | registers[b]
	case "bori":
		registersCopy[c] = registers[a] | b
	case "setr":
		registersCopy[c] = registers[a]
	case "seti":
		registersCopy[c] = a
	case "gtir":
		if a > registers[b] {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	case "gtri":
		if registers[a] > b {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	case "gtrr":
		if registers[a] > registers[b] {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	case "eqir":
		if a == registers[b] {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	case "eqri":
		if registers[a] == b {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	case "eqrr":
		if registers[a] == registers[b] {
			registersCopy[c] = 1
		} else {
			registersCopy[c] = 0
		}
	}
	return registersCopy
}

func part1(input string) {
	ipBind, instructions := parseInput(input)
	registers := make([]int, 6)
	ip := 0

	for ip >= 0 && ip < len(instructions) {
		registers[ipBind] = ip
		instruction := instructions[ip]
		registers = execute(registers, instruction)
		ip = registers[ipBind]
		ip++
	}

	fmt.Println("Part 1", registers[0])
}

func part2(input string) {
	ipBind, instructions := parseInput(input)
	registers := make([]int, 6)
	registers[0] = 1
	ip := 0

	for ip >= 0 && ip < len(instructions) {
		registers[ipBind] = ip
		instruction := instructions[ip]
		registers = execute(registers, instruction)

		if ip == 32 {
			break
		}

		ip = registers[ipBind]
		ip++
	}

	numberToFactorize := registers[2] + registers[5]
	result := sumOfDivisors(numberToFactorize)
	fmt.Println("Part 2", result)
}

func sumOfDivisors(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
