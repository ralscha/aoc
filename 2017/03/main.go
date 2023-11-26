package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2017, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	n := conv.MustAtoi(input[:len(input)-1])
	side := int(math.Ceil(math.Sqrt(float64(n))))
	prevSideMax := side*side - (side - 1)
	distanceToCenter := mathx.Abs(n - prevSideMax - side/2)
	steps := side/2 + distanceToCenter
	fmt.Println(steps)
}

func part2(input string) {
	n := conv.MustAtoi(input[:len(input)-1])

	grid := make([][]int, 100)
	for i := range grid {
		grid[i] = make([]int, 100)
	}
	grid[50][50] = 1
	x, y := 51, 50
	direction := 0
	for {
		grid[y][x] = sumNeighbors(grid, x, y)
		if grid[y][x] > n {
			fmt.Println(grid[y][x])
			return
		}
		x, y, direction = nextPosition(x, y, direction, grid)
	}

}

func sumNeighbors(grid [][]int, x, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += grid[y+i][x+j]
		}
	}
	return sum
}

func nextPosition(x, y, direction int, grid [][]int) (int, int, int) {
	switch direction {
	case 0:
		if grid[y-1][x] == 0 {
			return x, y - 1, 1
		}
		return x + 1, y, 0
	case 1:
		if grid[y][x-1] == 0 {
			return x - 1, y, 2
		}
		return x, y - 1, 1
	case 2:
		if grid[y+1][x] == 0 {
			return x, y + 1, 3
		}
		return x - 1, y, 2
	case 3:
		if grid[y][x+1] == 0 {
			return x + 1, y, 0
		}
		return x, y + 1, 3
	default:
		log.Fatalln("invalid direction")
		return 0, 0, 0
	}
}
