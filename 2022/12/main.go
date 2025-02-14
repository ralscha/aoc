package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"container/list"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2022, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[i]))
	}

	var startRow, startCol int
	var endRow, endCol int

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'S' {
				startRow, startCol = i, j
				grid[i][j] = 1
			} else if lines[i][j] == 'E' {
				endRow, endCol = i, j
				grid[i][j] = 'z' - 'a' + 1
			} else {
				grid[i][j] = int(lines[i][j] - 'a' + 1)
			}
		}
	}

	steps := bfs(grid, gridutil.Coordinate{Row: startRow, Col: startCol}, gridutil.Coordinate{Row: endRow, Col: endCol})
	fmt.Println("Part 1", steps)
}

func bfs(grid [][]int, start, end gridutil.Coordinate) int {
	noRows := len(grid)
	noCols := len(grid[0])

	queue := list.New()
	queue.PushBack(start)
	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(start)
	steps := make(map[gridutil.Coordinate]int)

	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{-1, 1, 0, 0}

	for queue.Len() > 0 {
		front := queue.Front()
		head := front.Value.(gridutil.Coordinate)
		queue.Remove(front)
		if head.Row == end.Row && head.Col == end.Col {
			break
		}
		for i := range 4 {
			newRow := head.Row + dx[i]
			newCol := head.Col + dy[i]
			if newRow >= 0 && newRow < noRows && newCol >= 0 && newCol < noCols && !visited.Contains(gridutil.Coordinate{Row: newRow, Col: newCol}) &&
				grid[newRow][newCol] < grid[head.Row][head.Col]+2 {
				steps[gridutil.Coordinate{Row: newRow, Col: newCol}] = steps[head] + 1
				visited.Add(gridutil.Coordinate{Row: newRow, Col: newCol})
				queue.PushBack(gridutil.Coordinate{Row: newRow, Col: newCol})
			}
		}
	}

	return steps[end]
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[i]))
	}

	var startPoints []gridutil.Coordinate
	var endRow, endCol int

	for r := range lines {
		for c := range lines[r] {
			if lines[r][c] == 'S' {
				startPoints = append(startPoints, gridutil.Coordinate{Row: r, Col: c})
				grid[r][c] = 1
			} else if lines[r][c] == 'E' {
				endRow, endCol = r, c
				grid[r][c] = 'z' - 'a' + 1
			} else {
				if lines[r][c] == 'a' {
					startPoints = append(startPoints, gridutil.Coordinate{Row: r, Col: c})
				}
				grid[r][c] = int(lines[r][c] - 'a' + 1)
			}
		}
	}

	minSteps := math.MaxInt
	for _, start := range startPoints {
		steps := bfs(grid, start, gridutil.Coordinate{Row: endRow, Col: endCol})
		if steps != 0 && steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println("Part 2", minSteps)
}
