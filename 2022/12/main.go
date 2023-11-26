package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
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

type point struct {
	row, col int
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

	steps := bfs(grid, point{startRow, startCol}, point{endRow, endCol})
	fmt.Println(steps)
}

func bfs(grid [][]int, start, end point) int {
	noRows := len(grid)
	noCols := len(grid[0])

	queue := list.New()
	queue.PushBack(start)
	visited := make(map[point]bool)
	visited[start] = true
	steps := make(map[point]int)

	dx := [4]int{0, 0, -1, 1}
	dy := [4]int{-1, 1, 0, 0}

	for queue.Len() > 0 {
		front := queue.Front()
		head := front.Value.(point)
		queue.Remove(front)
		if head.row == end.row && head.col == end.col {
			break
		}
		for i := 0; i < 4; i++ {
			newRow := head.row + dx[i]
			newCol := head.col + dy[i]
			if newRow >= 0 && newRow < noRows && newCol >= 0 && newCol < noCols && !visited[point{newRow, newCol}] &&
				grid[newRow][newCol] < grid[head.row][head.col]+2 {
				steps[point{newRow, newCol}] = steps[head] + 1
				visited[point{newRow, newCol}] = true
				queue.PushBack(point{newRow, newCol})
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

	var startPoints []point
	var endRow, endCol int

	for r := range lines {
		for c := range lines[r] {
			if lines[r][c] == 'S' {
				startPoints = append(startPoints, point{r, c})
				grid[r][c] = 1
			} else if lines[r][c] == 'E' {
				endRow, endCol = r, c
				grid[r][c] = 'z' - 'a' + 1
			} else {
				if lines[r][c] == 'a' {
					startPoints = append(startPoints, point{r, c})
				}
				grid[r][c] = int(lines[r][c] - 'a' + 1)
			}
		}
	}

	minSteps := math.MaxInt
	for _, start := range startPoints {
		steps := bfs(grid, start, point{endRow, endCol})
		if steps != 0 && steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println(minSteps)
}
