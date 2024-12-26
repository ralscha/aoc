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
	input, err := download.ReadInput(2018, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var ipReg int
	var program []instruction

	for _, line := range lines {
		if strings.HasPrefix(line, "#ip") {
			ipReg = conv.MustAtoi(line[4:])
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 4 {
			op := fields[0]
			a := conv.MustAtoi(fields[1])
			b := conv.MustAtoi(fields[2])
			c := conv.MustAtoi(fields[3])
			program = append(program, instruction{op, a, b, c})
		}
	}

	reg := [6]int{}
	seen := make(map[int]bool)

	for ip := 0; ip >= 0 && ip < len(program); {
		reg[ipReg] = ip
		inst := program[ip]

		if ip == 28 {
			if len(seen) == 0 {
				fmt.Println("Part 1", reg[program[28].a])
				return
			}
			seen[reg[program[28].a]] = true
		}

		reg = exec(inst, reg)
		ip = reg[ipReg]
		ip++
	}
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var ipReg int
	var program []instruction

	for _, line := range lines {
		if strings.HasPrefix(line, "#ip") {
			ipReg = conv.MustAtoi(line[4:])
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 4 {
			op := fields[0]
			a := conv.MustAtoi(fields[1])
			b := conv.MustAtoi(fields[2])
			c := conv.MustAtoi(fields[3])
			program = append(program, instruction{op, a, b, c})
		}
	}

	reg := [6]int{}
	seen := container.NewSet[int]()
	var lastSeen int
	ip := 0

	for ip >= 0 && ip < len(program) {
		reg[ipReg] = ip
		inst := program[ip]

		if ip == 28 {
			val := reg[program[28].a]
			if seen.Contains(val) {
				fmt.Println("Part 2", lastSeen)
				return
			}
			seen.Add(val)
			lastSeen = val
		}

		reg = exec(inst, reg)
		ip = reg[ipReg]
		ip++
	}
}

type instruction struct {
	op      string
	a, b, c int
}

func exec(inst instruction, reg [6]int) [6]int {
	switch inst.op {
	case "addr":
		reg[inst.c] = reg[inst.a] + reg[inst.b]
	case "addi":
		reg[inst.c] = reg[inst.a] + inst.b
	case "mulr":
		reg[inst.c] = reg[inst.a] * reg[inst.b]
	case "muli":
		reg[inst.c] = reg[inst.a] * inst.b
	case "banr":
		reg[inst.c] = reg[inst.a] & reg[inst.b]
	case "bani":
		reg[inst.c] = reg[inst.a] & inst.b
	case "borr":
		reg[inst.c] = reg[inst.a] | reg[inst.b]
	case "bori":
		reg[inst.c] = reg[inst.a] | inst.b
	case "setr":
		reg[inst.c] = reg[inst.a]
	case "seti":
		reg[inst.c] = inst.a
	case "gtir":
		if inst.a > reg[inst.b] {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	case "gtri":
		if reg[inst.a] > inst.b {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	case "gtrr":
		if reg[inst.a] > reg[inst.b] {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	case "eqir":
		if inst.a == reg[inst.b] {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	case "eqri":
		if reg[inst.a] == inst.b {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	case "eqrr":
		if reg[inst.a] == reg[inst.b] {
			reg[inst.c] = 1
		} else {
			reg[inst.c] = 0
		}
	}
	return reg
}
