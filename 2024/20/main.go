package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func solve(grid gridutil.Grid2D[rune], maxCheatDistance int) int {
	var start, end gridutil.Coordinate
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if val, ok := grid.Get(row, col); ok {
				if val == 'S' {
					start = gridutil.Coordinate{Row: row, Col: col}
				} else if val == 'E' {
					end = gridutil.Coordinate{Row: row, Col: col}
				}
			}
		}
	}

	mainPath, found := grid.ShortestPathWithBFS(start,
		func(pos gridutil.Coordinate, val rune) bool {
			return pos == end
		},
		func(pos gridutil.Coordinate, val rune) bool {
			return val == '#'
		},
		gridutil.Get4Directions())

	if !found {
		log.Fatal("No path found")
	}

	pathPositions := mainPath.Path

	cheatsCount := 0
	for i := 0; i < len(pathPositions)-1; i++ {
		for j := i + 1; j < len(pathPositions); j++ {
			savedBySkipping := j - i

			distance := mathx.Abs(pathPositions[i].Row-pathPositions[j].Row) +
				mathx.Abs(pathPositions[i].Col-pathPositions[j].Col)

			if distance <= maxCheatDistance {
				timeSaved := savedBySkipping - distance

				if timeSaved >= 100 {
					cheatsCount++
				}
			}
		}
	}

	return cheatsCount
}

func main() {
	input, err := download.ReadInput(2024, 20)
	if err != nil {
		log.Fatal(err)
	}
	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	fmt.Println("Part 1", solve(grid, 2))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	fmt.Println("Part 2", solve(grid, 20))
}
