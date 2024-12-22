package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

type button struct {
	x, y int
}

type prize struct {
	x, y int
}

type machine struct {
	buttonA button
	buttonB button
	prize   prize
}

func main() {
	input, err := download.ReadInput(2024, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func (m machine) findSolution() (a, b int, possible bool) {
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if a*m.buttonA.x+b*m.buttonB.x == m.prize.x &&
				a*m.buttonA.y+b*m.buttonB.y == m.prize.y {
				return a, b, true
			}
		}
	}
	return 0, 0, false
}

func (m machine) findSolutionCramer() (a, b int, possible bool) {
	det := m.buttonA.x*m.buttonB.y - m.buttonB.x*m.buttonA.y
	detX := m.prize.x*m.buttonB.y - m.buttonB.x*m.prize.y
	detY := m.buttonA.x*m.prize.y - m.prize.x*m.buttonA.y

	if detX%det != 0 || detY%det != 0 {
		return 0, 0, false
	}

	a = detX / det
	b = detY / det
	return a, b, true
}

func parseMachine(lines []string) machine {
	var m machine

	conv.MustSscanf(lines[0], "Button A: X+%d, Y+%d", &m.buttonA.x, &m.buttonA.y)
	conv.MustSscanf(lines[1], "Button B: X+%d, Y+%d", &m.buttonB.x, &m.buttonB.y)
	conv.MustSscanf(lines[2], "Prize: X=%d, Y=%d", &m.prize.x, &m.prize.y)

	return m
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var machines []machine

	for i := 0; i < len(lines); i += 4 {
		machines = append(machines, parseMachine(lines[i:i+3]))
	}

	totalTokens := 0
	for _, machine := range machines {
		if a, b, possible := machine.findSolution(); possible {
			totalTokens += a*3 + b*1
		}
	}

	fmt.Println("Part 1", totalTokens)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var machines []machine

	const offset = 10000000000000

	for i := 0; i < len(lines); i += 4 {
		machine := parseMachine(lines[i : i+3])
		machine.prize.x += offset
		machine.prize.y += offset
		machines = append(machines, machine)
	}

	totalTokens := 0
	for _, machine := range machines {
		if a, b, possible := machine.findSolutionCramer(); possible {
			totalTokens += a*3 + b*1
		}
	}

	fmt.Println("Part 2", totalTokens)
}
