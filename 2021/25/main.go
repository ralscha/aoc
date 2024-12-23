package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
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
	empty = '.'
	east  = '>'
	south = 'v'
)

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[rune](false)

	// Build grid
	for row, line := range lines {
		for col, char := range line {
			grid.Set(row, col, char)
		}
	}

	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	height := maxRow - minRow + 1
	width := maxCol - minCol + 1

	turn := 1
	for {
		moved := false

		// Move east-facing cucumbers
		newGrid := grid.Copy()
		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				if val, _ := grid.Get(row, col); val == east {
					nextCol := (col-minCol+1)%width + minCol
					if nextVal, _ := grid.Get(row, nextCol); nextVal == empty {
						newGrid.Set(row, col, empty)
						newGrid.Set(row, nextCol, east)
						moved = true
					}
				}
			}
		}
		grid = newGrid

		// Move south-facing cucumbers
		newGrid = grid.Copy()
		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				if val, _ := grid.Get(row, col); val == south {
					nextRow := (row-minRow+1)%height + minRow
					if nextVal, _ := grid.Get(nextRow, col); nextVal == empty {
						newGrid.Set(row, col, empty)
						newGrid.Set(nextRow, col, south)
						moved = true
					}
				}
			}
		}
		grid = newGrid

		if !moved {
			break
		}
		turn++
	}

	fmt.Println("Part 1", turn)
}
