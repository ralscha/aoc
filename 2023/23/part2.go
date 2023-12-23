package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

var (
	dx         = []int{0, 0, 1, -1}
	dy         = []int{1, -1, 0, 0}
	maxLength2 = 0
)

func isValid(x, y int, grid [][]byte, visited [][]bool) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) && grid[x][y] != '#' && !visited[x][y]
}

func dfs2(x, y int, grid [][]byte, visited [][]bool, length int) {
	if grid[x][y] == '.' && x == len(grid)-1 {
		if length > maxLength2 {
			maxLength2 = length
		}
		return
	}

	visited[x][y] = true

	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if isValid(nx, ny, grid, visited) {
			dfs2(nx, ny, grid, visited, length+1)
		}
	}

	visited[x][y] = false
}

func findLongestPath(grid [][]byte) int {
	startX, startY := 0, 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	for y, cell := range grid[0] {
		if cell == '.' {
			startY = y
			break
		}
	}

	dfs2(startX, startY, grid, visited, 0)
	return maxLength2
}

func part2() {
	input, err := download.ReadInput(2023, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	lines := conv.SplitNewline(input)

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	fmt.Println(findLongestPath(grid))
}
