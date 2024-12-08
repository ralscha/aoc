package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2021, 9)
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

	var lowPoints []int
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			currentValue, _ := grid.Get(row, col)
			isLower := true

			neighbors := grid.GetNeighbours4C(gridutil.Coordinate{Row: row, Col: col})
			for _, neighbor := range neighbors {
				if currentValue >= neighbor {
					isLower = false
					break
				}
			}

			if isLower {
				lowPoints = append(lowPoints, currentValue)
			}
		}
	}

	total := 0
	for _, lp := range lowPoints {
		total += lp + 1
	}

	fmt.Println("Part 1", total)
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

	var basinSizes []int
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if isLowPoint(grid, coord) {
				size := findBasinSize(grid, coord)
				if size > 0 {
					basinSizes = append(basinSizes, size)
				}
			}
		}
	}

	slices.SortFunc(basinSizes, func(i, j int) int {
		return j - i
	})
	result := basinSizes[0] * basinSizes[1] * basinSizes[2]
	fmt.Println("Part 2", result)
}

func isLowPoint(grid gridutil.Grid2D[int], coord gridutil.Coordinate) bool {
	currentValue, _ := grid.GetC(coord)
	neighbors := grid.GetNeighbours4C(coord)

	for _, neighbor := range neighbors {
		if currentValue >= neighbor {
			return false
		}
	}
	return true
}

func findBasinSize(grid gridutil.Grid2D[int], start gridutil.Coordinate) int {
	visited := container.NewSet[gridutil.Coordinate]()
	return crawl(grid, start, visited)
}

func crawl(grid gridutil.Grid2D[int], coord gridutil.Coordinate, visited *container.Set[gridutil.Coordinate]) int {
	if visited.Contains(coord) {
		return 0
	}

	value, ok := grid.GetC(coord)
	if !ok || value == 9 {
		return 0
	}

	visited.Add(coord)
	size := 1

	// Get valid neighbors (up, down, left, right)
	for _, dir := range []gridutil.Direction{
		{Row: -1, Col: 0}, // up
		{Row: 1, Col: 0},  // down
		{Row: 0, Col: -1}, // left
		{Row: 0, Col: 1},  // right
	} {
		next := gridutil.Coordinate{
			Row: coord.Row + dir.Row,
			Col: coord.Col + dir.Col,
		}
		size += crawl(grid, next, visited)
	}

	return size
}
