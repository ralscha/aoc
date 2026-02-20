package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	grid = tiltGrid(grid, 0, -1)
	totalLoad := calculateLoad(grid)
	fmt.Println(totalLoad)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	directions := []struct {
		dx, dy int
	}{
		{0, -1},
		{-1, 0},
		{0, 1},
		{1, 0},
	}

	seenStates := make(map[string]int)
	var cycleLength, cycleStart int

	for cycle := range 1_000_000_000 {
		for _, dir := range directions {
			grid = tiltGrid(grid, dir.dx, dir.dy)
		}

		gridStr := gridToString(grid)
		if start, ok := seenStates[gridStr]; ok {
			cycleLength = cycle - start
			cycleStart = start
			break
		} else {
			seenStates[gridStr] = cycle
		}
	}

	remainingCycles := (1_000_000_000-cycleStart)%cycleLength - 1
	for range remainingCycles {
		for _, dir := range directions {
			grid = tiltGrid(grid, dir.dx, dir.dy)
		}
	}

	totalLoad := calculateLoad(grid)
	fmt.Println(totalLoad)
}

func tiltGrid(grid [][]byte, dx, dy int) [][]byte {
	height := len(grid)
	width := len(grid[0])
	moved := true

	for moved {
		moved = false

		for y := range height {
			for x := range width {
				if grid[y][x] == 'O' {
					newX, newY := x+dx, y+dy
					if newX >= 0 && newY >= 0 && newX < width && newY < height && grid[newY][newX] == '.' {
						grid[y][x], grid[newY][newX] = '.', 'O'
						moved = true
					}
				}
			}
		}
	}
	return grid
}

func gridToString(grid [][]byte) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return sb.String()
}

func calculateLoad(grid [][]byte) int {
	totalLoad := 0
	for y, row := range grid {
		for _, char := range row {
			if char == 'O' {
				load := len(grid) - y
				totalLoad += load
			}
		}
	}
	return totalLoad
}
