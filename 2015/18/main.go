package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]bool, len(lines))
	for i := range grid {
		grid[i] = make([]bool, len(lines[i]))
	}
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				grid[i][j] = true
			}
		}
	}

	rounds := 100
	for i := 0; i < rounds; i++ {
		newGrid := make([][]bool, len(grid))
		for i := range grid {
			newGrid[i] = make([]bool, len(grid[i]))
			for j := range grid[i] {
				newGrid[i][j] = grid[i][j]
				neighbors := countNeighborsOn(grid, i, j)
				if grid[i][j] {
					if neighbors != 2 && neighbors != 3 {
						newGrid[i][j] = false
					}
				} else if neighbors == 3 {
					newGrid[i][j] = true
				}
			}
		}

		grid = newGrid
	}

	fmt.Println(countOn(grid))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]bool, len(lines))
	for i := range grid {
		grid[i] = make([]bool, len(lines[i]))
	}
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				grid[i][j] = true
			}
		}
	}

	turnOnCorners(grid)

	rounds := 100
	for i := 0; i < rounds; i++ {
		newGrid := make([][]bool, len(grid))
		for i := range grid {
			newGrid[i] = make([]bool, len(grid[i]))
			for j := range grid[i] {
				newGrid[i][j] = grid[i][j]
				neighbors := countNeighborsOn(grid, i, j)
				if grid[i][j] {
					if neighbors != 2 && neighbors != 3 {
						newGrid[i][j] = false
					}
				} else if neighbors == 3 {
					newGrid[i][j] = true
				}
			}
		}

		turnOnCorners(newGrid)

		grid = newGrid
	}

	fmt.Println(countOn(grid))
}

func turnOnCorners(newGrid [][]bool) {
	newGrid[0][0] = true
	newGrid[0][len(newGrid[0])-1] = true
	newGrid[len(newGrid)-1][0] = true
	newGrid[len(newGrid)-1][len(newGrid[0])-1] = true
}

func countNeighborsOn(grid [][]bool, r, c int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if r+i < 0 || r+i >= len(grid) {
				continue
			}
			if c+j < 0 || c+j >= len(grid[0]) {
				continue
			}
			if grid[r+i][c+j] {
				count++
			}
		}
	}
	return count
}

func countOn(grid [][]bool) int {
	count := 0
	for _, row := range grid {
		for _, c := range row {
			if c {
				count++
			}
		}
	}
	return count
}
