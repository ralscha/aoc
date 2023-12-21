package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2023, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type pos struct {
	x, y int
}

var directions = []pos{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type state struct {
	pos   pos
	steps int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)

	var garden [][]bool
	var startX, startY int

	row := 0
	for _, line := range lines {
		gardenRow := make([]bool, len(line))
		for col, char := range line {
			if char == 'S' {
				startX, startY = row, col
				gardenRow[col] = true
			} else {
				gardenRow[col] = char == '.'
			}
		}
		garden = append(garden, gardenRow)
		row++
	}

	ways := make(map[state]int)
	ways[state{pos{startX, startY}, 0}] = 1

	for step := 1; step <= 64; step++ {
		nextWays := make(map[state]int)
		for s, count := range ways {
			for _, dir := range directions {
				nextPos := pos{s.pos.x + dir.x, s.pos.y + dir.y}
				if nextPos.x >= 0 && nextPos.x < len(garden) && nextPos.y >= 0 && nextPos.y < len(garden[0]) && garden[nextPos.x][nextPos.y] {
					nextState := state{nextPos, step}
					nextWays[nextState] += count
				}
			}
		}
		ways = nextWays
	}

	var count uint64 = 0
	for state := range ways {
		if state.steps == 64 {
			count++
		}
	}

	fmt.Println(count)

}
