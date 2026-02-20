package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2021, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[int](false)

	// Build grid
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, int(c-'0'))
		}
	}

	totalFlashes := 0
	for range 100 {
		totalFlashes += simulateStep(&grid)
	}

	fmt.Println("Part 1", totalFlashes)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[int](false)

	// Build grid
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, int(c-'0'))
		}
	}

	step := 1
	for {
		flashes := simulateStep(&grid)
		if flashes == 100 { // All octopuses flashed
			fmt.Println("Part 2", step)
			break
		}
		step++
	}
}

func simulateStep(grid *gridutil.Grid2D[int]) int {
	// First, increase all energy levels by 1
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			val, _ := grid.Get(row, col)
			grid.Set(row, col, val+1)
		}
	}

	// Track flashed octopuses
	flashed := gridutil.NewGrid2D[bool](false)
	flashes := 0

	// Keep flashing until no more flashes occur
	for {
		newFlashes := false
		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				coord := gridutil.Coordinate{Row: row, Col: col}
				hasFlashed, _ := flashed.GetC(coord)
				energy, _ := grid.GetC(coord)

				if energy > 9 && !hasFlashed {
					// Flash this octopus
					flashed.SetC(coord, true)
					flashes++
					newFlashes = true

					// Increase energy of all neighbors
					for _, dir := range gridutil.Get8Directions() {
						neighborCoord := gridutil.Coordinate{
							Row: coord.Row + dir.Row,
							Col: coord.Col + dir.Col,
						}
						if val, ok := grid.GetC(neighborCoord); ok {
							grid.SetC(neighborCoord, val+1)
						}
					}
				}
			}
		}
		if !newFlashes {
			break
		}
	}

	// Reset energy levels of flashed octopuses
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if hasFlashed, _ := flashed.GetC(coord); hasFlashed {
				grid.SetC(coord, 0)
			}
		}
	}

	return flashes
}
