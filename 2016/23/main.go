package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type instructions struct {
	opcode string
	args   []string
}

func parseInstructions(input string) []instructions {
	lines := conv.SplitNewline(input)
	ins := make([]instructions, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		ins[i] = instructions{opcode: parts[0], args: parts[1:]}
	}
	return ins
}

func getValue(arg string, registers map[string]int) int {
	if val, err := strconv.Atoi(arg); err == nil {
		return val
	}
	return registers[arg]
}

func executeOptimized(instructions []instructions, registers map[string]int) {
	ip := 0
	for ip >= 0 && ip < len(instructions) {
		instruction := instructions[ip]
		opcode := instruction.opcode
		args := instruction.args

		switch opcode {
		case "cpy":
			if len(args) == 2 {
				val := getValue(args[0], registers)
				if _, ok := registers[args[1]]; ok {
					registers[args[1]] = val
				}
			}
			ip++
		case "inc":
			if len(args) == 1 {
				if _, ok := registers[args[0]]; ok {
					registers[args[0]]++
				}
			}
			ip++
		case "dec":
			if len(args) == 1 {
				if _, ok := registers[args[0]]; ok {
					registers[args[0]]--
				}
			}
			ip++
		case "jnz":
			if len(args) == 2 {
				val := getValue(args[0], registers)
				offset := getValue(args[1], registers)
				if val != 0 {
					ip += offset
				} else {
					ip++
				}
			} else {
				ip++
			}
		case "tgl":
			if len(args) == 1 {
				offset := getValue(args[0], registers)
				target := ip + offset
				if target >= 0 && target < len(instructions) {
					toggle(&instructions[target])
				}
			}
			ip++
		default:
			ip++
		}
	}
}

func executeOptimizedPart2(instructions []instructions, registers map[string]int) {
	ip := 0
	for ip >= 0 && ip < len(instructions) {
		instruction := instructions[ip]
		opcode := instruction.opcode
		args := instruction.args

		if ip == 4 && instructions[ip].opcode == "cpy" && instructions[ip+1].opcode == "inc" && instructions[ip+2].opcode == "dec" && instructions[ip+3].opcode == "jnz" && instructions[ip+4].opcode == "dec" && instructions[ip+5].opcode == "jnz" {
			if args[1] == "c" && instructions[ip+2].args[0] == "c" && instructions[ip+3].args[1] == "-2" && instructions[ip+4].args[0] == "d" && instructions[ip+5].args[1] == "-5" {
				b := registers["b"]
				d := registers["d"]
				registers["a"] += b * d
				registers["c"] = 0
				registers["d"] = 0
				ip = 10
				continue
			}
		}

		if ip == 11 && instructions[ip].opcode == "cpy" && instructions[ip+1].opcode == "inc" && instructions[ip+2].opcode == "dec" && instructions[ip+3].opcode == "jnz" && instructions[ip+4].opcode == "dec" && instructions[ip+5].opcode == "jnz" {
			if args[1] == "c" && instructions[ip+2].args[0] == "c" && instructions[ip+3].args[1] == "-2" && instructions[ip+4].args[0] == "b" && instructions[ip+5].args[1] == "-5" {
				f := registers["f"]
				g := registers["g"]
				registers["a"] += f * g
				registers["b"] = 0
				registers["c"] = 0
				ip = 18
				continue
			}
		}

		switch opcode {
		case "cpy":
			if len(args) == 2 {
				var val int
				if v, err := strconv.Atoi(args[0]); err == nil {
					val = v
				} else {
					val = registers[args[0]]
				}
				if _, ok := registers[args[1]]; ok {
					registers[args[1]] = val
				}
			}
			ip++
		case "inc":
			if len(args) == 1 {
				if _, ok := registers[args[0]]; ok {
					registers[args[0]]++
				}
			}
			ip++
		case "dec":
			if len(args) == 1 {
				if _, ok := registers[args[0]]; ok {
					registers[args[0]]--
				}
			}
			ip++
		case "jnz":
			if len(args) == 2 {
				val := getValue(args[0], registers)
				offset := getValue(args[1], registers)
				if val != 0 {
					ip += offset
				} else {
					ip++
				}
			} else {
				ip++
			}
		case "tgl":
			if len(args) == 1 {
				offset := getValue(args[0], registers)
				target := ip + offset
				if target >= 0 && target < len(instructions) {
					toggle(&instructions[target])
				}
			}
			ip++
		default:
			ip++
		}
	}
}

func toggle(instruction *instructions) {
	switch len(instruction.args) {
	case 1:
		switch instruction.opcode {
		case "inc":
			instruction.opcode = "dec"
		default:
			instruction.opcode = "inc"
		}
	case 2:
		switch instruction.opcode {
		case "jnz":
			instruction.opcode = "cpy"
		default:
			instruction.opcode = "jnz"
		}
	}
}
func part1(input string) {
	registers := map[string]int{"a": 7, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0}
	instructions := parseInstructions(input)

	executeOptimized(instructions, registers)

	fmt.Println("Part 1", registers["a"])
}

func part2(input string) {
	registersPart2 := map[string]int{"a": 12, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0}
	instructionsPart2 := parseInstructions(input)
	executeOptimizedPart2(instructionsPart2, registersPart2)
	fmt.Println("Part 2", registersPart2["a"])
}
