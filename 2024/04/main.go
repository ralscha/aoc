package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func countWord(grid gridutil.Grid2D[rune], word string) int {
	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	directions := gridutil.Get8Directions()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			pos := gridutil.Coordinate{Col: col, Row: row}
			for _, dir := range directions {
				if isWordAt(grid, word, pos, dir) {
					count++
				}
			}
		}
	}

	return count
}

func isWordAt(grid gridutil.Grid2D[rune], word string, start gridutil.Coordinate, dir gridutil.Direction) bool {
	pos := start
	for _, char := range word {
		val, ok := grid.GetC(pos)
		if !ok || val != char {
			return false
		}
		pos.Col += dir.Col
		pos.Row += dir.Row
	}
	return true
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	grid := gridutil.NewCharGrid2D(lines)
	word := "XMAS"
	result := countWord(grid, word)
	fmt.Println("Part 1", result)
}

func countCrossMAS(grid gridutil.Grid2D[rune]) int {
	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			pos := gridutil.Coordinate{Col: col, Row: row}
			if isCrossMAS(grid, pos) {
				count++
			}
		}
	}

	return count
}

func isCrossMAS(grid gridutil.Grid2D[rune], center gridutil.Coordinate) bool {
	// Check center is 'A'
	centerVal, ok := grid.GetC(center)
	if !ok || centerVal != 'A' {
		return false
	}

	// Get diagonal positions
	topLeft := gridutil.Coordinate{Col: center.Col - 1, Row: center.Row - 1}
	topRight := gridutil.Coordinate{Col: center.Col + 1, Row: center.Row - 1}
	bottomLeft := gridutil.Coordinate{Col: center.Col - 1, Row: center.Row + 1}
	bottomRight := gridutil.Coordinate{Col: center.Col + 1, Row: center.Row + 1}

	// Get values at diagonal positions
	tl, oktl := grid.GetC(topLeft)
	tr, oktr := grid.GetC(topRight)
	bl, okbl := grid.GetC(bottomLeft)
	br, okbr := grid.GetC(bottomRight)

	if !oktl || !oktr || !okbl || !okbr {
		return false
	}

	// Check all possible valid patterns
	return (tl == 'M' && tr == 'S' && bl == 'M' && br == 'S') || // Pattern 1
		(tl == 'S' && tr == 'M' && bl == 'S' && br == 'M') || // Pattern 2
		(tl == 'S' && tr == 'S' && bl == 'M' && br == 'M') || // Pattern 3
		(tl == 'M' && tr == 'M' && bl == 'S' && br == 'S') // Pattern 4
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	grid := gridutil.NewCharGrid2D(lines)
	result := countCrossMAS(grid)
	fmt.Println("Part 2", result)
}
