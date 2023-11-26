package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2016, 1)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	splitted := strings.Fields(input)
	var instructions []string
	for _, s := range splitted {
		ins := s
		if strings.HasSuffix(s, ",") {
			ins = s[:len(s)-1]
		}
		instructions = append(instructions, ins)
	}

	direction := 0
	x, y := 0, 0

	for _, ins := range instructions {
		if ins[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}
		steps := ins[1:]
		switch direction {
		case 0:
			y += conv.MustAtoi(steps)
		case 1:
			x += conv.MustAtoi(steps)
		case 2:
			y -= conv.MustAtoi(steps)
		case 3:
			x -= conv.MustAtoi(steps)
		}

	}

	fmt.Printf("Part 1: %d\n", mathx.Abs(x)+mathx.Abs(y))
}

func part2(input string) {
	splitted := strings.Fields(input)
	var instructions []string
	for _, s := range splitted {
		ins := s
		if strings.HasSuffix(s, ",") {
			ins = s[:len(s)-1]
		}
		instructions = append(instructions, ins)
	}

	direction := 0
	x, y := 0, 0
	firstVisitedX := 0
	firstVisitedY := 0
	visited := make(map[string]bool)

	for _, ins := range instructions {
		if ins[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}
		steps := ins[1:]

		for i := 0; i < conv.MustAtoi(steps); i++ {
			switch direction {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			}

			if firstVisitedX == 0 && firstVisitedY == 0 {
				key := fmt.Sprintf("%d,%d", x, y)
				if !visited[key] {
					visited[key] = true
				} else {
					firstVisitedX = x
					firstVisitedY = y
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", mathx.Abs(firstVisitedX)+mathx.Abs(firstVisitedY))
}
