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
	input, err := download.ReadInput(2021, 15)
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
		for col, val := range line {
			grid.Set(row, col, int(val-'0'))
		}
	}

	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	start := gridutil.Coordinate{Row: minRow, Col: minCol}
	target := gridutil.Coordinate{Row: maxRow, Col: maxCol}

	risk := findLowestRisk(grid, start, target)
	fmt.Println("Part 1", risk)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	baseGrid := gridutil.NewGrid2D[int](false)

	// Build base grid
	for row, line := range lines {
		for col, val := range line {
			baseGrid.Set(row, col, int(val-'0'))
		}
	}

	// Create expanded grid
	minRow, maxRow := baseGrid.GetMinMaxRow()
	minCol, maxCol := baseGrid.GetMinMaxCol()
	height := maxRow - minRow + 1
	width := maxCol - minCol + 1

	expandedGrid := gridutil.NewGrid2D[int](false)
	for row := 0; row < height*5; row++ {
		for col := 0; col < width*5; col++ {
			baseValue, _ := baseGrid.Get(row%height, col%width)
			increase := row/height + col/width
			value := baseValue + increase
			if value > 9 {
				value = value - 9
			}
			expandedGrid.Set(row, col, value)
		}
	}

	start := gridutil.Coordinate{Row: 0, Col: 0}
	target := gridutil.Coordinate{Row: height*5 - 1, Col: width*5 - 1}

	risk := findLowestRisk(expandedGrid, start, target)
	fmt.Println("Part 2", risk)
}

func findLowestRisk(grid gridutil.Grid2D[int], start, target gridutil.Coordinate) int {
	lowestRiskAt := make(map[gridutil.Coordinate]int)
	pq := container.NewPriorityQueue[gridutil.Coordinate]()
	pq.Push(start, 0)

	directions := []gridutil.Direction{
		{Row: -1, Col: 0}, // up
		{Row: 1, Col: 0},  // down
		{Row: 0, Col: -1}, // left
		{Row: 0, Col: 1},  // right
	}

	for !pq.IsEmpty() {
		current := pq.Pop()

		for _, dir := range directions {
			next := gridutil.Coordinate{
				Row: current.Row + dir.Row,
				Col: current.Col + dir.Col,
			}

			// Check if next position is valid
			if risk, ok := grid.GetC(next); ok {
				nextRisk := lowestRiskAt[current] + risk
				if existingRisk, exists := lowestRiskAt[next]; !exists || nextRisk < existingRisk {
					lowestRiskAt[next] = nextRisk
					pq.Push(next, nextRisk)
				}
			}
		}
	}

	return lowestRiskAt[target]
}
