package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
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

	pos := gridutil.Coordinate{Col: 0, Row: 0}
	dir := gridutil.DirectionN

	for _, ins := range instructions {
		if ins[0] == 'R' {
			dir = gridutil.TurnRight(dir)
		} else {
			dir = gridutil.TurnLeft(dir)
		}
		steps := conv.MustAtoi(ins[1:])
		pos.Col += dir.Col * steps
		pos.Row += dir.Row * steps
	}

	fmt.Println("Part 1", mathx.Abs(pos.Row)+mathx.Abs(pos.Col))
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

	pos := gridutil.Coordinate{Row: 0, Col: 0}
	dir := gridutil.DirectionN // Start facing north
	visited := container.NewSet[gridutil.Coordinate]()
	firstVisited := gridutil.Coordinate{}
	foundFirst := false

	for _, ins := range instructions {
		if ins[0] == 'R' {
			dir = gridutil.TurnRight(dir)
		} else {
			dir = gridutil.TurnLeft(dir)
		}
		steps := conv.MustAtoi(ins[1:])

		for range steps {
			pos.Col += dir.Col
			pos.Row += dir.Row

			if !foundFirst {
				if visited.Contains(pos) {
					firstVisited = pos
					foundFirst = true
				} else {
					visited.Add(pos)
				}
			}
		}
	}

	fmt.Println("Part 2", mathx.Abs(firstVisited.Row)+mathx.Abs(firstVisited.Col))
}
