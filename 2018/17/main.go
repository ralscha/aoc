package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type waterGrid struct {
	grid   gridutil.Grid2D[byte]
	minCol int
	maxCol int
	minRow int
	maxRow int
}

func newWaterGrid() *waterGrid {
	return &waterGrid{
		grid:   gridutil.NewGrid2D[byte](false),
		minCol: 2000,
		maxCol: 0,
		minRow: 2000,
		maxRow: 0,
	}
}

func (w *waterGrid) updateBounds(col, row int) {
	if col < w.minCol {
		w.minCol = col
	}
	if col > w.maxCol {
		w.maxCol = col
	}
	if row < w.minRow {
		w.minRow = row
	}
	if row > w.maxRow {
		w.maxRow = row
	}
}

func (w *waterGrid) isOpen(coord gridutil.Coordinate) bool {
	val, exists := w.grid.GetC(coord)
	return !exists || val == '|'
}

func (w *waterGrid) fill(coord gridutil.Coordinate) {
	if coord.Row > w.maxRow {
		return
	}
	if !w.isOpen(coord) {
		return
	}

	below := gridutil.Coordinate{Row: coord.Row + 1, Col: coord.Col}

	if !w.isOpen(below) {
		leftCol := coord.Col
		leftCoord := gridutil.Coordinate{Row: coord.Row, Col: leftCol}
		for w.isOpen(leftCoord) && !w.isOpen(gridutil.Coordinate{Row: coord.Row + 1, Col: leftCol}) {
			w.grid.SetC(leftCoord, '|')
			leftCol--
			leftCoord.Col = leftCol
		}

		rightCol := coord.Col + 1
		rightCoord := gridutil.Coordinate{Row: coord.Row, Col: rightCol}
		for w.isOpen(rightCoord) && !w.isOpen(gridutil.Coordinate{Row: coord.Row + 1, Col: rightCol}) {
			w.grid.SetC(rightCoord, '|')
			rightCol++
			rightCoord.Col = rightCol
		}

		leftVal, _ := w.grid.Get(coord.Row, leftCol)
		rightVal, _ := w.grid.Get(coord.Row, rightCol)

		if w.isOpen(gridutil.Coordinate{Row: coord.Row + 1, Col: leftCol}) ||
			w.isOpen(gridutil.Coordinate{Row: coord.Row + 1, Col: rightCol}) {
			w.fill(gridutil.Coordinate{Row: coord.Row, Col: leftCol})
			w.fill(gridutil.Coordinate{Row: coord.Row, Col: rightCol})
		} else if leftVal == '#' && rightVal == '#' {
			for col := leftCol + 1; col < rightCol; col++ {
				w.grid.Set(coord.Row, col, '~')
			}
		}
	} else {
		_, exists := w.grid.GetC(coord)
		if !exists {
			w.grid.SetC(coord, '|')
			w.fill(below)
			belowVal, _ := w.grid.GetC(below)
			if belowVal == '~' {
				w.fill(coord)
			}
		}
	}
}

func part1and2(input string) {
	waterGrid := newWaterGrid()
	lines := conv.SplitNewline(input)

	for _, line := range lines {
		parts := strings.Split(line, ", ")
		firstPart := parts[0][2:]
		a := conv.MustAtoi(firstPart)

		rangeParts := strings.Split(parts[1][2:], "..")
		bMin := conv.MustAtoi(rangeParts[0])
		bMax := conv.MustAtoi(rangeParts[1])

		if parts[0][0] == 'x' {
			for row := bMin; row <= bMax; row++ {
				waterGrid.updateBounds(a, row)
				waterGrid.grid.Set(row, a, '#')
			}
		} else {
			for col := bMin; col <= bMax; col++ {
				waterGrid.updateBounds(col, a)
				waterGrid.grid.Set(a, col, '#')
			}
		}
	}

	waterGrid.fill(gridutil.Coordinate{Row: 0, Col: 500})

	water, touched := 0, 0
	for col := waterGrid.minCol - 1; col <= waterGrid.maxCol+1; col++ {
		for row := waterGrid.minRow; row <= waterGrid.maxRow; row++ {
			if val, exists := waterGrid.grid.Get(row, col); exists {
				if val == '|' {
					touched++
				} else if val == '~' {
					water++
				}
			}
		}
	}

	fmt.Println("Part 1", water+touched)
	fmt.Println("Part 2", water)
}
