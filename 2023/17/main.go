package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2023, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, char := range line {
			n := int(char - '0')
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	leastHeatLoss := minHeatLoss(grid, 1, 3)
	fmt.Println(leastHeatLoss)

	leastHeatLoss = minHeatLoss(grid, 4, 10)
	fmt.Println(leastHeatLoss)
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

type position struct {
	x, y int
	dir  direction
}

type state struct {
	pos      position
	heatLoss int
}

func minHeatLoss(grid [][]int, minDistance, maxDistance int) int {
	rows := len(grid)
	cols := len(grid[0])
	target := position{x: cols - 1, y: rows - 1}

	pq := container.NewPriorityQueue[state]()

	start := state{pos: position{x: 0, y: 0, dir: -1}, heatLoss: 0}
	pq.Push(start, start.heatLoss)

	visited := make(map[position]bool)

	for pq.Len() > 0 {
		currentState := pq.Pop()
		currentPos := currentState.pos

		if currentPos.x == target.x && currentPos.y == target.y {
			return currentState.heatLoss
		}

		if visited[currentPos] {
			continue
		}
		visited[currentPos] = true

		for _, d := range []direction{up, right, down, left} {
			if currentPos.dir == up && d == down || currentPos.dir == right && d == left || currentPos.dir == down && d == up || currentPos.dir == left && d == right {
				continue
			}
			if currentPos.dir == d {
				continue
			}

			incrementHeat := 0
			for distance := 1; distance <= maxDistance; distance++ {
				nextX, nextY := currentPos.x, currentPos.y
				switch d {
				case up:
					nextY = nextY - distance
				case right:
					nextX = nextX + distance
				case down:
					nextY = nextY + distance
				case left:
					nextX = nextX - distance
				}

				if nextX < 0 || nextX >= cols || nextY < 0 || nextY >= rows {
					continue
				}

				incrementHeat += grid[nextY][nextX]
				if distance < minDistance {
					continue
				}
				nextPos := position{x: nextX, y: nextY, dir: d}
				nextHeatLoss := currentState.heatLoss + incrementHeat
				nextState := state{pos: nextPos, heatLoss: nextHeatLoss}
				pq.Push(nextState, nextState.heatLoss)
			}
		}
	}

	return math.MaxInt32
}
