package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type cpu struct {
	registers map[string]int
	pc        int
	program   []instruction
}

type instruction struct {
	op   string
	args []string
}

func newCPU(initialRegisters map[string]int) *cpu {
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	for reg, val := range initialRegisters {
		registers[reg] = val
	}
	return &cpu{
		registers: registers,
		pc:        0,
	}
}

func (c *cpu) getValue(arg string) int {
	if val, exists := c.registers[arg]; exists {
		return val
	}
	return conv.MustAtoi(arg)
}

func (c *cpu) execute() {
	for c.pc >= 0 && c.pc < len(c.program) {
		inst := c.program[c.pc]

		switch inst.op {
		case "cpy":
			value := c.getValue(inst.args[0])
			c.registers[inst.args[1]] = value
			c.pc++

		case "inc":
			c.registers[inst.args[0]]++
			c.pc++

		case "dec":
			c.registers[inst.args[0]]--
			c.pc++

		case "jnz":
			value := c.getValue(inst.args[0])
			offset := conv.MustAtoi(inst.args[1])
			if value != 0 {
				c.pc += offset
			} else {
				c.pc++
			}
		}
	}
}

func parseProgram(input string) []instruction {
	lines := conv.SplitNewline(input)
	instructions := make([]instruction, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		inst := instruction{
			op:   parts[0],
			args: parts[1:],
		}
		instructions = append(instructions, inst)
	}

	return instructions
}

func solve(input string, initialRegisters map[string]int) int {
	cpu := newCPU(initialRegisters)
	cpu.program = parseProgram(input)
	cpu.execute()
	return cpu.registers["a"]
}

func part1(input string) {
	result := solve(input, map[string]int{})
	fmt.Println("Part 1", result)
}

func part2(input string) {
	result := solve(input, map[string]int{"c": 1})
	fmt.Println("Part 2", result)
}
