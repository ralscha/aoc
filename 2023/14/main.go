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

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	tiltGrid(&grid, 0, -1)
	totalLoad := calculateLoad(grid)
	fmt.Println(totalLoad)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
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

	for cycle := 0; cycle < 1000000000; cycle++ {
		for _, dir := range directions {
			tiltGrid(&grid, dir.dx, dir.dy)
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

	remainingCycles := (1000000000-cycleStart)%cycleLength - 1
	for cycle := 0; cycle < remainingCycles; cycle++ {
		for _, dir := range directions {
			tiltGrid(&grid, dir.dx, dir.dy)
		}
	}

	totalLoad := calculateLoad(grid)
	fmt.Println(totalLoad)
}

func tiltGrid(grid *[][]rune, dx, dy int) {
	height := len(*grid)
	width := len((*grid)[0])
	moved := true

	for moved {
		moved = false
		newGrid := make([][]rune, height)
		for i := range newGrid {
			newGrid[i] = make([]rune, width)
			copy(newGrid[i], (*grid)[i])
		}

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if (*grid)[y][x] == 'O' {
					newX, newY := x+dx, y+dy
					if newX >= 0 && newY >= 0 && newX < width && newY < height && (*grid)[newY][newX] == '.' {
						newGrid[y][x], newGrid[newY][newX] = '.', 'O'
						moved = true
					}
				}
			}
		}
		*grid = newGrid
	}
}

func gridToString(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteRune('\n')
	}
	return sb.String()
}

func calculateLoad(grid [][]rune) int {
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
