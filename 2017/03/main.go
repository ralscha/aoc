package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
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

	grid := gridutil.NewGrid2D[int](false)
	grid.SetMinRowCol(-50, -50) // Center the grid
	grid.SetMaxRowCol(50, 50)

	center := gridutil.Coordinate{Col: 0, Row: 0}
	grid.SetC(center, 1)

	pos := gridutil.Coordinate{Col: 1, Row: 0}
	dir := gridutil.Direction{Col: 0, Row: -1} // Start going up

	for {
		// Sum all 8 neighbors
		sum := 0
		for _, val := range grid.GetNeighbours8C(pos) {
			sum += val
		}
		grid.SetC(pos, sum)

		if sum > n {
			fmt.Println(sum)
			return
		}

		// Try to turn left
		leftDir := gridutil.TurnLeft(dir)
		leftPos := gridutil.Coordinate{Col: pos.Col + leftDir.Col, Row: pos.Row + leftDir.Row}
		if val, _ := grid.GetC(leftPos); val == 0 {
			// Can turn left
			pos = leftPos
			dir = leftDir
		} else {
			// Continue straight
			pos = gridutil.Coordinate{Col: pos.Col + dir.Col, Row: pos.Row + dir.Row}
		}
	}
}
