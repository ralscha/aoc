package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	grid2 "aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := grid2.NewCharGrid2D(lines)

	start := grid2.Coordinate{}
	direction := grid2.DirectionN

outerLoop:
	for row := range grid.Height() {
		for col := range grid.Width() {
			cell, _ := grid.Get(row, col)
			if cell == '^' {
				start = grid2.Coordinate{Col: col, Row: row}
				break outerLoop
			}
		}
	}

	visited := make(map[grid2.Coordinate]bool)
	visited[start] = true

	currentPos := start
	for {
		nextCell, outside := grid.Peek(currentPos.Row, currentPos.Col, direction)
		if outside {
			break
		}

		if nextCell == '#' {
			direction = grid2.TurnRight(direction)
		} else {
			currentPos = grid2.Coordinate{
				Row: currentPos.Row + direction.Row,
				Col: currentPos.Col + direction.Col,
			}
			visited[currentPos] = true
		}
	}

	fmt.Println("Part 1", len(visited))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := grid2.NewCharGrid2D(lines)

	start := grid2.Coordinate{}
	direction := grid2.DirectionN

outerLoop:
	for row := range grid.Height() {
		for col := range grid.Width() {
			cell, _ := grid.Get(row, col)
			if cell == '^' {
				start = grid2.Coordinate{Col: col, Row: row}
				break outerLoop
			}
		}
	}

	var possibleObstructions []grid2.Coordinate
	for row := range grid.Height() {
		for col := range grid.Width() {
			cell, _ := grid.Get(row, col)
			if cell == '.' && !(col == start.Col && row == start.Row) {
				possibleObstructions = append(possibleObstructions, grid2.Coordinate{Col: col, Row: row})
			}
		}
	}

	loopCount := 0
	for _, obstruction := range possibleObstructions {
		if causesLoop(grid, start, direction, obstruction) {
			loopCount++
		}
	}

	fmt.Println("Part 2", loopCount)
}

type coordinateWithDirection struct {
	coord     grid2.Coordinate
	direction grid2.Direction
}

func causesLoop(grid grid2.Grid2D[rune], start grid2.Coordinate, direction grid2.Direction, obstruction grid2.Coordinate) bool {
	modifiedGrid := grid.Copy()
	modifiedGrid.Set(obstruction.Row, obstruction.Col, '#')

	visited := make(map[coordinateWithDirection]bool)
	currentPos := coordinateWithDirection{coord: start, direction: direction}
	visited[currentPos] = true

	for {
		nextCell, outside := modifiedGrid.Peek(currentPos.coord.Row, currentPos.coord.Col, currentPos.direction)
		if outside {
			break
		}

		if nextCell == '#' {
			currentPos.direction = grid2.TurnRight(currentPos.direction)
		} else {
			currentPos.coord = grid2.Coordinate{
				Row: currentPos.coord.Row + currentPos.direction.Row,
				Col: currentPos.coord.Col + currentPos.direction.Col,
			}
			if visited[currentPos] {
				return true
			}
			visited[currentPos] = true
		}
	}

	return false
}
