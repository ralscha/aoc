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
	input, err := download.ReadInput(2016, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func run(instructions []string, initialRegisters map[string]int) (string, bool) {
	registers := make(map[string]int)
	for k, v := range initialRegisters {
		registers[k] = v
	}
	pc := 0
	output := ""
	for len(output) < 20 {
		if pc >= len(instructions) {
			return output, false
		}
		instruction := instructions[pc]
		parts := strings.Split(instruction, " ")
		switch parts[0] {
		case "cpy":
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				val = registers[parts[1]]
			}
			registers[parts[2]] = val
			pc++
		case "inc":
			registers[parts[1]]++
			pc++
		case "dec":
			registers[parts[1]]--
			pc++
		case "jnz":
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				val = registers[parts[1]]
			}
			jump, err := strconv.Atoi(parts[2])
			if err != nil {
				jump = registers[parts[2]]
			}
			if val != 0 {
				pc += jump
			} else {
				pc++
			}
		case "out":
			val := registers[parts[1]]
			output += strconv.Itoa(val)
			pc++
		}
	}
	return output, true
}

func part1(input string) {
	instructions := conv.SplitNewline(input)
	for i := 0; ; i++ {
		initialRegisters := map[string]int{
			"a": i,
			"b": 0,
			"c": 0,
			"d": 0,
		}
		output, valid := run(instructions, initialRegisters)
		if valid {
			validOutput := true
			for j := 0; j < len(output); j++ {
				if j%2 == 0 && output[j] != '0' {
					validOutput = false
					break
				}
				if j%2 == 1 && output[j] != '1' {
					validOutput = false
					break
				}
			}
			if validOutput {
				fmt.Println("Part 1", i)
				return
			}
		}
	}
}
