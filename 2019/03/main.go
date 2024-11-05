package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	grid2 "aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	wire1 := getWirePoints(lines[0])
	wire2 := getWirePoints(lines[1])

	minDist := math.MaxInt
	for coord := range wire1 {
		if wire2[coord] {
			dist := mathx.MannhattanDistance(0, 0, coord.Row, coord.Col)
			if dist < minDist {
				minDist = dist
			}
		}
	}
	fmt.Println("Result 1:", minDist)

}

func getWirePoints(path string) map[grid2.Coordinate]bool {
	wire := make(map[grid2.Coordinate]bool)
	x, y := 0, 0
	coords := strings.Split(path, ",")
	for _, coord := range coords {
		dir := coord[0]
		dist := conv.MustAtoi(coord[1:])
		for j := 0; j < dist; j++ {
			switch dir {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			coord := grid2.Coordinate{Row: x, Col: y}
			wire[coord] = true
		}
	}
	return wire
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	wire1 := getWirePointsWithSteps(lines[0])
	wire2 := getWirePointsWithSteps(lines[1])

	minSteps := math.MaxInt
	for coord, steps1 := range wire1 {
		if steps2, ok := wire2[coord]; ok {
			steps := steps1 + steps2
			if steps < minSteps {
				minSteps = steps
			}
		}
	}
	fmt.Println("Result 2:", minSteps)
}

func getWirePointsWithSteps(path string) map[grid2.Coordinate]int {
	wire := make(map[grid2.Coordinate]int)
	x, y := 0, 0
	steps := 0
	coords := strings.Split(path, ",")
	for _, coord := range coords {
		dir := coord[0]
		dist := conv.MustAtoi(coord[1:])
		for j := 0; j < dist; j++ {
			steps++
			switch dir {
			case 'U':
				y++
			case 'D':
				y--
			case 'L':
				x--
			case 'R':
				x++
			}
			coord := grid2.Coordinate{Row: x, Col: y}
			if _, ok := wire[coord]; !ok {
				wire[coord] = steps
			}
		}
	}
	return wire
}
