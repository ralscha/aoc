package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func countWord(grid [][]rune, word string) int {
	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range [][2]int{
				{0, 1},
				{1, 0},
				{1, 1},
				{1, -1},
				{0, -1},
				{-1, 0},
				{-1, -1},
				{-1, 1},
			} {
				if isWordAt(grid, word, r, c, dir[0], dir[1], wordLen) {
					count++
				}
			}
		}
	}

	return count
}

func isWordAt(grid [][]rune, word string, startRow, startCol, rowDir, colDir, wordLen int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < wordLen; i++ {
		r := startRow + i*rowDir
		c := startCol + i*colDir
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != rune(word[i]) {
			return false
		}
	}

	return true
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	word := "XMAS"
	result := countWord(grid, word)
	fmt.Println("Part 1", result)
}

func countCrossMAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if isCrossMAS(grid, r, c) {
				count++
			}
		}
	}

	return count
}

func isCrossMAS(grid [][]rune, startRow, startCol int) bool {
	rows := len(grid)
	cols := len(grid[0])

	if startRow-1 < 0 || startRow+1 >= rows || startCol-1 < 0 || startCol+1 >= cols {
		return false
	}

	if grid[startRow][startCol] != 'A' {
		return false
	}

	isXMAS1 := grid[startRow-1][startCol-1] == 'M' &&
		grid[startRow-1][startCol+1] == 'S' &&
		grid[startRow+1][startCol-1] == 'M' &&
		grid[startRow+1][startCol+1] == 'S'
	if isXMAS1 {
		return true
	}

	isXMAS2 := grid[startRow-1][startCol-1] == 'S' &&
		grid[startRow-1][startCol+1] == 'M' &&
		grid[startRow+1][startCol-1] == 'S' &&
		grid[startRow+1][startCol+1] == 'M'
	if isXMAS2 {
		return true
	}

	ixXMAS3 := grid[startRow-1][startCol-1] == 'S' &&
		grid[startRow-1][startCol+1] == 'S' &&
		grid[startRow+1][startCol-1] == 'M' &&
		grid[startRow+1][startCol+1] == 'M'
	if ixXMAS3 {
		return true
	}

	ixXMAS4 := grid[startRow-1][startCol-1] == 'M' &&
		grid[startRow-1][startCol+1] == 'M' &&
		grid[startRow+1][startCol-1] == 'S' &&
		grid[startRow+1][startCol+1] == 'S'

	return ixXMAS4
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	result := countCrossMAS(grid)
	fmt.Println("Part 2", result)
}
