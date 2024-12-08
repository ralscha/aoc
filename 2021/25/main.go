package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2021, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

const (
	east  = 0
	south = 1
)

type pos struct {
	row, col int
}

type cucumber int

func part1(input string) {
	lines := conv.SplitNewline(input)
	width, height := len(lines[0]), len(lines)
	cucumbers := make(map[pos]cucumber)

	for row, line := range lines {
		for col, char := range line {
			if char == '>' {
				cucumbers[pos{row, col}] = east
			} else if char == 'v' {
				cucumbers[pos{row, col}] = south
			}
		}
	}

	directions := [2]pos{{0, 1}, {1, 0}}

	turn := 1
	for {
		moved := false

		for dir, direction := range directions {
			allMovableCucumbers := make(map[pos]pos)
			for p, c := range cucumbers {
				if c == cucumber(dir) {
					newCol := (p.col + direction.col) % width
					newRow := (p.row + direction.row) % height
					newPos := pos{newRow, newCol}
					if _, ok := cucumbers[newPos]; !ok {
						allMovableCucumbers[p] = newPos
					}
				}
			}
			for oldPos, newPos := range allMovableCucumbers {
				c := cucumbers[oldPos]
				delete(cucumbers, oldPos)
				cucumbers[newPos] = c
				moved = true
			}
		}

		if !moved {
			break
		}
		turn++
	}
	fmt.Println("Part 1", turn)
}
