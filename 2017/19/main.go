package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2017, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	letters, steps := solve(input)
	fmt.Println("Part 1", letters)
	fmt.Println("Part 2", steps)
}

func findStart(grid *gridutil.Grid2D[rune]) gridutil.Coordinate {
	for col := 0; col < grid.Width(); col++ {
		if val, ok := grid.Get(0, col); ok && val == '|' {
			return gridutil.Coordinate{Row: 0, Col: col}
		}
	}
	return gridutil.Coordinate{}
}

func getPossibleTurns(dir gridutil.Direction) []gridutil.Direction {
	return []gridutil.Direction{
		{Row: -dir.Col, Col: dir.Row},
		{Row: dir.Col, Col: -dir.Row},
	}
}

func solve(input string) (string, int) {
	grid := gridutil.NewCharGrid2D(conv.SplitNewline(input))
	pos := findStart(&grid)
	dir := gridutil.DirectionS
	letters := ""
	steps := 1

	for {
		nextPos := gridutil.Coordinate{
			Row: pos.Row + dir.Row,
			Col: pos.Col + dir.Col,
		}
		val, ok := grid.GetC(nextPos)
		if !ok || val == ' ' {
			break
		}

		pos = nextPos
		steps++

		if 'A' <= val && val <= 'Z' {
			letters += string(val)
		} else if val == '+' {
			turned := false
			for _, turn := range getPossibleTurns(dir) {
				newPos := gridutil.Coordinate{
					Row: pos.Row + turn.Row,
					Col: pos.Col + turn.Col,
				}
				if nextVal, ok := grid.GetC(newPos); ok && nextVal != ' ' {
					dir = turn
					turned = true
					break
				}
			}
			if !turned {
				break
			}
		}
	}

	return letters, steps
}
