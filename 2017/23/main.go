package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2()
}

func part1(input string) {
	program := parseProgram(input)
	processor := newProcessor(program)
	processor.execute()
	fmt.Println("Part 1", processor.mulCount)
}

func part2() {
	b := 57*100 + 100000
	c := b + 17000
	h := 0

	for n := b; n <= c; n += 17 {
		if len(mathx.Factors(n)) > 2 {
			h++
		}
	}

	fmt.Println("Part 2", h)
}

type instruction struct {
	op   string
	x, y string
}

type processor struct {
	registers map[string]int
	pc        int
	mulCount  int
	program   []instruction
}

func newProcessor(program []instruction) *processor {
	return &processor{
		registers: make(map[string]int),
		program:   program,
	}
}

func (p *processor) getValue(s string) int {
	if val, ok := p.registers[s]; ok {
		return val
	}
	if num, err := strconv.Atoi(s); err == nil {
		return num
	}
	return 0
}

func (p *processor) execute() {
	for p.pc >= 0 && p.pc < len(p.program) {
		inst := p.program[p.pc]

		switch inst.op {
		case "set":
			p.registers[inst.x] = p.getValue(inst.y)
			p.pc++
		case "sub":
			p.registers[inst.x] -= p.getValue(inst.y)
			p.pc++
		case "mul":
			p.registers[inst.x] *= p.getValue(inst.y)
			p.mulCount++
			p.pc++
		case "jnz":
			if p.getValue(inst.x) != 0 {
				p.pc += p.getValue(inst.y)
			} else {
				p.pc++
			}
		}
	}
}

func parseProgram(input string) []instruction {
	lines := conv.SplitNewline(input)
	program := make([]instruction, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		inst := instruction{op: parts[0], x: parts[1]}
		if len(parts) > 2 {
			inst.y = parts[2]
		}
		program[i] = inst
	}

	return program
}
