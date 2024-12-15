package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 15)
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func move(grid *gridutil.Grid2D[rune], pos, newPos gridutil.Coordinate, dir gridutil.Direction) gridutil.Coordinate {
	currentVal, ok := grid.GetC(pos)
	if !ok {
		return pos
	}

	val, ok := grid.GetC(newPos)
	if !ok {
		return pos
	}

	switch val {
	case '.':
		grid.SetC(newPos, currentVal)
		grid.SetC(pos, '.')
		return newPos
	case '#':
		return pos
	case 'O':
		nextPos := gridutil.Coordinate{
			Row: newPos.Row + dir.Row,
			Col: newPos.Col + dir.Col,
		}
		if newPos != move(grid, newPos, nextPos, dir) {
			grid.SetC(newPos, currentVal)
			grid.SetC(pos, '.')
			return newPos
		}
		return pos
	case '[', ']':
		if dir == gridutil.DirectionW || dir == gridutil.DirectionE {
			nextPos := gridutil.Coordinate{
				Row: newPos.Row + dir.Row,
				Col: newPos.Col + dir.Col,
			}
			if newPos != move(grid, newPos, nextPos, dir) {
				grid.SetC(newPos, currentVal)
				grid.SetC(pos, '.')
				return newPos
			}
			return pos
		} else {
			var otherHalf gridutil.Coordinate
			if val == '[' {
				otherHalf = gridutil.Coordinate{Row: newPos.Row, Col: newPos.Col + 1}
			} else {
				otherHalf = gridutil.Coordinate{Row: newPos.Row, Col: newPos.Col - 1}
			}

			nextPos := gridutil.Coordinate{
				Row: newPos.Row + dir.Row,
				Col: newPos.Col + dir.Col,
			}
			otherNextPos := gridutil.Coordinate{
				Row: otherHalf.Row + dir.Row,
				Col: otherHalf.Col + dir.Col,
			}

			if canMove(grid, newPos, nextPos) && canMove(grid, otherHalf, otherNextPos) {
				move(grid, newPos, nextPos, dir)
				move(grid, otherHalf, otherNextPos, dir)
				grid.SetC(pos, '.')
				grid.SetC(newPos, currentVal)
				grid.SetC(otherHalf, '.')
				return newPos
			}
			return pos
		}
	}
	return pos
}

func canMove(grid *gridutil.Grid2D[rune], pos, newPos gridutil.Coordinate) bool {
	val, ok := grid.GetC(newPos)
	if !ok {
		return false
	}

	switch val {
	case '.':
		return true
	case '#':
		return false
	case 'O':
		nextPos := gridutil.Coordinate{
			Row: newPos.Row + (newPos.Row - pos.Row),
			Col: newPos.Col + (newPos.Col - pos.Col),
		}
		return pos != move(grid, newPos, nextPos, gridutil.Direction{
			Row: newPos.Row - pos.Row,
			Col: newPos.Col - pos.Col,
		})
	case '[', ']':
		dir := gridutil.Direction{
			Row: newPos.Row - pos.Row,
			Col: newPos.Col - pos.Col,
		}
		if dir == gridutil.DirectionW || dir == gridutil.DirectionE {
			nextPos := gridutil.Coordinate{
				Row: newPos.Row + dir.Row,
				Col: newPos.Col + dir.Col,
			}
			return canMove(grid, newPos, nextPos)
		} else {
			var otherHalf gridutil.Coordinate
			if val == '[' {
				otherHalf = gridutil.Coordinate{Row: newPos.Row, Col: newPos.Col + 1}
			} else {
				otherHalf = gridutil.Coordinate{Row: newPos.Row, Col: newPos.Col - 1}
			}

			nextPos := gridutil.Coordinate{
				Row: newPos.Row + dir.Row,
				Col: newPos.Col + dir.Col,
			}
			otherNextPos := gridutil.Coordinate{
				Row: otherHalf.Row + dir.Row,
				Col: otherHalf.Col + dir.Col,
			}

			return canMove(grid, newPos, nextPos) && canMove(grid, otherHalf, otherNextPos)
		}
	}
	return false
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	gridStr := parts[0]
	moves := strings.ReplaceAll(parts[1], "\n", "")

	lines := conv.SplitNewline(gridStr)
	grid := gridutil.NewCharGrid2D(lines)

	robotPos := gridutil.Coordinate{}
outerLoop:
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if ch, ok := grid.Get(row, col); ok && ch == '@' {
				robotPos = gridutil.Coordinate{Row: row, Col: col}
				break outerLoop
			}
		}
	}

	dirMap := map[rune]gridutil.Direction{
		'^': gridutil.DirectionN,
		'v': gridutil.DirectionS,
		'<': gridutil.DirectionW,
		'>': gridutil.DirectionE,
	}

	for _, m := range moves {
		dir := dirMap[m]
		newPos := gridutil.Coordinate{
			Row: robotPos.Row + dir.Row,
			Col: robotPos.Col + dir.Col,
		}
		robotPos = move(&grid, robotPos, newPos, dir)
	}

	sum := 0
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if val, ok := grid.Get(row, col); ok && val == 'O' {
				gps := 100*row + col
				sum += gps
			}
		}
	}
	fmt.Println("Part 1", sum)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	gridStr := parts[0]
	moves := strings.ReplaceAll(parts[1], "\n", "")

	lines := conv.SplitNewline(gridStr)
	doubledLines := make([]string, len(lines))
	for i, line := range lines {
		var newLine strings.Builder
		for _, ch := range line {
			switch ch {
			case '#':
				newLine.WriteString("##")
			case 'O':
				newLine.WriteString("[]")
			case '.':
				newLine.WriteString("..")
			case '@':
				newLine.WriteString("@.")
			}
		}
		doubledLines[i] = newLine.String()
	}

	grid := gridutil.NewCharGrid2D(doubledLines)

	robotPos := gridutil.Coordinate{}
outerLoop:
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if ch, ok := grid.Get(row, col); ok && ch == '@' {
				robotPos = gridutil.Coordinate{Row: row, Col: col}
				break outerLoop
			}
		}
	}

	dirMap := map[rune]gridutil.Direction{
		'^': gridutil.DirectionN,
		'v': gridutil.DirectionS,
		'<': gridutil.DirectionW,
		'>': gridutil.DirectionE,
	}

	for _, m := range moves {
		dir := dirMap[m]
		newPos := gridutil.Coordinate{
			Row: robotPos.Row + dir.Row,
			Col: robotPos.Col + dir.Col,
		}
		robotPos = move(&grid, robotPos, newPos, dir)
	}

	sum := 0
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if val, ok := grid.Get(row, col); ok && val == '[' {
				gps := 100*row + col
				sum += gps
			}
		}
	}

	fmt.Println("Part 2", sum)
}
