package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2016, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func isWall(row, col, favorite int) bool {
	val := col*col + 3*col + 2*col*row + row + row*row + favorite
	count := 0
	for val > 0 {
		count += val & 1
		val >>= 1
	}
	return count%2 != 0
}

func part1(input string) {
	favorite := conv.MustAtoi(input)
	start := gridutil.Coordinate{Row: 1, Col: 1}
	target := gridutil.Coordinate{Row: 39, Col: 31}

	grid := gridutil.NewGrid2D[bool](false)
	for row := 0; row <= target.Row; row++ {
		for col := 0; col <= target.Col; col++ {
			grid.SetC(gridutil.Coordinate{Row: row, Col: col}, isWall(row, col, favorite))
		}
	}

	isObstacle := func(currentPos gridutil.Coordinate, current bool) bool {
		return current
	}

	isGoal := func(currentPos gridutil.Coordinate, current bool) bool {
		return currentPos == target
	}

	directions := gridutil.Get4Directions()

	pathResult, found := grid.ShortestPathWithBFS(start, isGoal, isObstacle, directions)

	if found {
		fmt.Println("Part 1", len(pathResult.Path)-1)
	}
}

func part2(input string) {
	favorite := conv.MustAtoi(input)

	queue := container.NewQueue[[]int]()
	queue.Push([]int{1, 1, 0})
	visited := container.NewSet[gridutil.Coordinate]()
	visited.Add(gridutil.Coordinate{Row: 1, Col: 1})
	count := 0

	for !queue.IsEmpty() {
		curr := queue.Pop()
		x := curr[0]
		y := curr[1]
		steps := curr[2]

		if steps <= 50 {
			count++
		}

		if steps == 50 {
			continue
		}

		for _, move := range gridutil.Get4Directions() {
			nx := x + move.Col
			ny := y + move.Row
			if nx >= 0 && ny >= 0 && !isWall(nx, ny, favorite) {
				coord := gridutil.Coordinate{Row: nx, Col: ny}
				if !visited.Contains(coord) {
					visited.Add(coord)
					queue.Push([]int{nx, ny, steps + 1})
				}
			}
		}
	}
	fmt.Println("Part 2", count)
}
