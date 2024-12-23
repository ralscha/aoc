package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	registers := make(map[string]int)
	lastSound := 0
	pc := 0

	getValue := func(arg string) int {
		if val, err := strconv.Atoi(arg); err == nil {
			return val
		}
		return registers[arg]
	}

	for pc >= 0 && pc < len(lines) {
		parts := strings.Split(lines[pc], " ")
		opcode := parts[0]

		switch opcode {
		case "snd":
			lastSound = getValue(parts[1])
			pc++
		case "set":
			registers[parts[1]] = getValue(parts[2])
			pc++
		case "add":
			registers[parts[1]] += getValue(parts[2])
			pc++
		case "mul":
			registers[parts[1]] *= getValue(parts[2])
			pc++
		case "mod":
			registers[parts[1]] %= getValue(parts[2])
			pc++
		case "rcv":
			if getValue(parts[1]) != 0 {
				fmt.Println("Part 1", lastSound)
				return
			}
			pc++
		case "jgz":
			if getValue(parts[1]) > 0 {
				pc += getValue(parts[2])
			} else {
				pc++
			}
		}
	}
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	type program struct {
		id        int
		registers map[string]int
		pc        int
		queue     *container.Queue[int]
		sendCount int
	}

	program0 := &program{
		id:        0,
		registers: map[string]int{"p": 0},
		queue:     container.NewQueue[int](),
	}
	program1 := &program{
		id:        1,
		registers: map[string]int{"p": 1},
		queue:     container.NewQueue[int](),
	}

	getValue := func(p *program, arg string) int {
		if val, err := strconv.Atoi(arg); err == nil {
			return val
		}
		return p.registers[arg]
	}

	execute := func(p *program, other *program) bool {
		if p.pc < 0 || p.pc >= len(lines) {
			return false
		}

		parts := strings.Split(lines[p.pc], " ")
		opcode := parts[0]

		switch opcode {
		case "snd":
			val := getValue(p, parts[1])
			other.queue.Push(val)
			if p.id == 1 {
				p.sendCount++
			}
			p.pc++
		case "set":
			p.registers[parts[1]] = getValue(p, parts[2])
			p.pc++
		case "add":
			p.registers[parts[1]] += getValue(p, parts[2])
			p.pc++
		case "mul":
			p.registers[parts[1]] *= getValue(p, parts[2])
			p.pc++
		case "mod":
			p.registers[parts[1]] %= getValue(p, parts[2])
			p.pc++
		case "rcv":
			if p.queue.IsEmpty() {
				return false
			}
			p.registers[parts[1]] = p.queue.Pop()
			p.pc++
		case "jgz":
			if getValue(p, parts[1]) > 0 {
				p.pc += getValue(p, parts[2])
			} else {
				p.pc++
			}
		}
		return true
	}

	for {
		progress0 := execute(program0, program1)
		progress1 := execute(program1, program0)

		if !progress0 && !progress1 && program0.queue.IsEmpty() && program1.queue.IsEmpty() {
			break
		}
	}

	fmt.Println("Part 2", program1.sendCount)
}
