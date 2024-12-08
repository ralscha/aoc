package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
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
	grid := gridutil.NewCharGrid2D(lines)
	visited := guardRoute(grid)
	fmt.Println("Part 1", visited.Len())
}

func guardRoute(grid gridutil.Grid2D[rune]) *container.Set[gridutil.Coordinate] {
	start := gridutil.Coordinate{}
	direction := gridutil.Direction{Col: 0, Row: -1} // North

outerLoop:
	for row := range grid.Height() {
		for col := range grid.Width() {
			if val, _ := grid.Get(row, col); val == '^' {
				start = gridutil.Coordinate{Col: col, Row: row}
				break outerLoop
			}
		}
	}

	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(start)

	currentPos := start
	for {
		nextCell, outside := grid.PeekC(currentPos, direction)
		if outside {
			break
		}

		if nextCell == '#' {
			direction = gridutil.TurnRight(direction)
		} else {
			currentPos = gridutil.Coordinate{
				Col: currentPos.Col + direction.Col,
				Row: currentPos.Row + direction.Row,
			}
			visited.Add(currentPos)
		}
	}
	return visited
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	originalRoute := guardRoute(grid)

	start := gridutil.Coordinate{}
	direction := gridutil.Direction{Col: 0, Row: -1} // North

outerLoop:
	for row := range grid.Height() {
		for col := range grid.Width() {
			if val, _ := grid.Get(row, col); val == '^' {
				start = gridutil.Coordinate{Col: col, Row: row}
				break outerLoop
			}
		}
	}

	var possibleObstructions []gridutil.Coordinate
	for row := range grid.Height() {
		for col := range grid.Width() {
			pos := gridutil.Coordinate{Col: col, Row: row}
			if !originalRoute.Contains(pos) {
				continue
			}
			if val, _ := grid.GetC(pos); val == '.' && pos != start {
				possibleObstructions = append(possibleObstructions, pos)
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
	coord     gridutil.Coordinate
	direction gridutil.Direction
}

func causesLoop(grid gridutil.Grid2D[rune], start gridutil.Coordinate, direction gridutil.Direction, obstruction gridutil.Coordinate) bool {
	modifiedGrid := grid.Copy()
	modifiedGrid.SetC(obstruction, '#')

	visited := container.NewSet[coordinateWithDirection]()
	currentPos := coordinateWithDirection{coord: start, direction: direction}
	visited.Add(currentPos)

	for {
		nextCell, outside := modifiedGrid.PeekC(currentPos.coord, currentPos.direction)
		if outside {
			break
		}

		if nextCell == '#' {
			currentPos.direction = gridutil.TurnRight(currentPos.direction)
		} else {
			currentPos.coord = gridutil.Coordinate{
				Col: currentPos.coord.Col + currentPos.direction.Col,
				Row: currentPos.coord.Row + currentPos.direction.Row,
			}
			if visited.Contains(currentPos) {
				return true
			}
			visited.Add(currentPos)
		}
	}

	return false
}
