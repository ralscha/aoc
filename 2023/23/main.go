package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

type point struct {
	x, y int
}

var slopeDirections = map[byte]point{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

var maxLength int

func dfs(start, end point, grid map[point]byte, visited map[point]bool, length int) {
	if start == end {
		if length > maxLength {
			maxLength = length
		}
		return
	}

	visited[start] = true

	for _, delta := range slopeDirections {
		next := point{start.x + delta.x, start.y + delta.y}
		if tile, ok := grid[next]; ok && !visited[next] && (tile == '.' || slopeDirections[tile] == delta) {
			dfs(next, end, grid, visited, length+1)
		}
	}

	visited[start] = false
}

func main() {
	input, err := download.ReadInput(2023, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	lines := conv.SplitNewline(input)
	part1(lines)
	part2()
}

func part1(lines []string) {
	grid := make(map[point]byte)
	for y, line := range lines {
		for x, tile := range line {
			grid[point{x, y}] = byte(tile)
		}
	}

	start := findStart(grid)
	end := findEnd(len(lines)-1, grid)

	visited := make(map[point]bool)
	dfs(start, end, grid, visited, 0)
	fmt.Println(maxLength)
}

func findStart(grid map[point]byte) point {
	y := 0
	for k, v := range grid {
		if v == '.' && k.y == y {
			return point{k.x, y}
		}
	}
	panic("start not found")
}

func findEnd(lastRow int, grid map[point]byte) point {
	y := lastRow
	for k, v := range grid {
		if v == '.' && k.y == y {
			return point{k.x, y}
		}
	}
	panic("end not found")
}
