package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/geomutil"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2023, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	galaxies := container.NewSet[gridutil.Coordinate]()

	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	// Find galaxies
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if val, exists := grid.Get(row, col); exists && val == '#' {
				galaxies.Add(gridutil.Coordinate{Row: row, Col: col})
			}
		}
	}

	// Find empty rows and columns
	emptyRows := container.NewSet[int]()
	emptyCols := container.NewSet[int]()

	for row := minRow; row <= maxRow; row++ {
		rowEmpty := true
		for col := minCol; col <= maxCol; col++ {
			if galaxies.Contains(gridutil.Coordinate{Row: row, Col: col}) {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows.Add(row)
		}
	}

	for col := minCol; col <= maxCol; col++ {
		colEmpty := true
		for row := minRow; row <= maxRow; row++ {
			if galaxies.Contains(gridutil.Coordinate{Row: row, Col: col}) {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols.Add(col)
		}
	}

	// Convert galaxies to slice for easier iteration
	var galaxyPoints []gridutil.Coordinate
	for _, p := range galaxies.Values() {
		galaxyPoints = append(galaxyPoints, p)
	}

	sumPart1 := 0
	sumPart2 := 0
	for i := range len(galaxyPoints) - 1 {
		for j := i + 1; j < len(galaxyPoints); j++ {
			start := galaxyPoints[i]
			end := galaxyPoints[j]

			// Calculate Manhattan distance
			distance := geomutil.ManhattanDistance(start, end)
			sumPart1 += distance
			sumPart2 += distance

			// Count empty rows and columns between points
			minCol, maxCol := min(start.Col, end.Col), max(start.Col, end.Col)
			minRow, maxRow := min(start.Row, end.Row), max(start.Row, end.Row)

			for col := minCol; col <= maxCol; col++ {
				if emptyCols.Contains(col) {
					sumPart1 += 1
					sumPart2 += 1000000 - 1
				}
			}
			for row := minRow; row <= maxRow; row++ {
				if emptyRows.Contains(row) {
					sumPart1 += 1
					sumPart2 += 1000000 - 1
				}
			}
		}
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}
