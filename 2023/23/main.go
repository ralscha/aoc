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

// Part 1 globals
var maxLength int

// Part 2 globals
var (
	dx         = []int{0, 0, 1, -1}
	dy         = []int{1, -1, 0, 0}
	maxLength2 = 0
)

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

type junction struct {
	pos   point
	edges []edge
}

type edge struct {
	to     point
	length int
}

func isValidPoint(p point, grid [][]byte) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(grid) && p.y < len(grid[0]) && grid[p.x][p.y] != '#'
}

func getNeighbors(p point, grid [][]byte) []point {
	neighbors := make([]point, 0, 4)
	for i := range 4 {
		next := point{p.x + dx[i], p.y + dy[i]}
		if isValidPoint(next, grid) {
			neighbors = append(neighbors, next)
		}
	}
	return neighbors
}

func isJunction(p point, grid [][]byte) bool {
	if p.x == 0 || p.x == len(grid)-1 { // Start and end points are junctions
		return true
	}
	neighbors := getNeighbors(p, grid)
	return len(neighbors) > 2 // A junction has more than 2 possible directions
}

func findJunctions(grid [][]byte) map[point]*junction {
	junctions := make(map[point]*junction)

	// Add start point
	for y := 0; y < len(grid[0]); y++ {
		if grid[0][y] == '.' {
			junctions[point{0, y}] = &junction{pos: point{0, y}}
			break
		}
	}

	// Add end point
	for y := 0; y < len(grid[0]); y++ {
		if grid[len(grid)-1][y] == '.' {
			junctions[point{len(grid) - 1, y}] = &junction{pos: point{len(grid) - 1, y}}
			break
		}
	}

	// Find all other junctions
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			p := point{x, y}
			if grid[x][y] != '#' && isJunction(p, grid) {
				if _, exists := junctions[p]; !exists {
					junctions[p] = &junction{pos: p}
				}
			}
		}
	}

	// Find edges between junctions
	for start := range junctions {
		findEdges(start, grid, junctions)
	}

	return junctions
}

func findEdges(start point, grid [][]byte, junctions map[point]*junction) {
	visited := make(map[point]bool)
	queue := []struct {
		pos    point
		length int
	}{{start, 0}}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.length > 0 && junctions[current.pos] != nil {
			// Found path to another junction
			junctions[start].edges = append(junctions[start].edges, edge{
				to:     current.pos,
				length: current.length,
			})
			continue // Don't explore beyond this junction
		}

		for _, next := range getNeighbors(current.pos, grid) {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, struct {
					pos    point
					length int
				}{next, current.length + 1})
			}
		}
	}
}

func dfsJunctions(current point, end point, junctions map[point]*junction, visited map[point]bool, length int) {
	if current == end {
		if length > maxLength2 {
			maxLength2 = length
		}
		return
	}

	visited[current] = true
	for _, edge := range junctions[current].edges {
		if !visited[edge.to] {
			dfsJunctions(edge.to, end, junctions, visited, length+edge.length)
		}
	}
	visited[current] = false
}

func findLongestPath(grid [][]byte) int {
	junctions := findJunctions(grid)

	var start, end point
	// Find start and end points
	for y := 0; y < len(grid[0]); y++ {
		if grid[0][y] == '.' {
			start = point{0, y}
		}
		if grid[len(grid)-1][y] == '.' {
			end = point{len(grid) - 1, y}
		}
	}

	visited := make(map[point]bool)
	dfsJunctions(start, end, junctions, visited, 0)
	return maxLength2
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

func part2(lines []string) {
	grid := make([][]byte, len(lines))
	for i := range grid {
		grid[i] = make([]byte, len(lines[0]))
		for j, ch := range lines[i] {
			grid[i][j] = byte(ch)
		}
	}
	fmt.Println(findLongestPath(grid))
}

func main() {
	input, err := download.ReadInput(2023, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	lines := conv.SplitNewline(input)
	part1(lines)
	part2(lines)
}
