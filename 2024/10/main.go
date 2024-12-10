package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func findTrailheads(grid gridutil.Grid2D[int]) []gridutil.Coordinate {
	var trailheads []gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if val, exists := grid.Get(row, col); exists && val == 0 {
				trailheads = append(trailheads, gridutil.Coordinate{Row: row, Col: col})
			}
		}
	}
	return trailheads
}

func solvePart1(lines []string) int {
	grid := gridutil.NewNumberGrid2D(lines)
	trailheads := findTrailheads(grid)

	totalScore := 0
	for _, start := range trailheads {
		reachable9s := gridutil.CollectNodes(
			&grid,
			start,
			func(value int) bool { return value == 9 },
			func(current, next int) bool { return next == current+1 },
		)
		totalScore += len(reachable9s)
	}

	return totalScore
}

func solvePart2(lines []string) int {
	grid := gridutil.NewNumberGrid2D(lines)
	trailheads := findTrailheads(grid)

	totalRating := 0
	for _, start := range trailheads {
		paths := gridutil.CountPaths(
			&grid,
			start,
			func(value int) bool { return value == 9 },
			func(current, next int) bool { return next == current+1 },
		)
		totalRating += paths
	}

	return totalRating
}

func main() {
	input, err := download.ReadInput(2024, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	result := solvePart1(lines)
	fmt.Println("Part 1", result)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	result := solvePart2(lines)
	fmt.Println("Part 2", result)
}
