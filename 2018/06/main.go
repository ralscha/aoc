package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/grid"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2018/06/input.txt"
	input, err := download.ReadInput(inputFile, 2018, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	places := make([]grid.Coordinate, len(lines))

	minCol, minRow := math.MaxInt32, math.MaxInt32
	maxCol, maxRow := 0, 0

	for i, line := range lines {
		parts := strings.Split(line, ", ")
		x := conv.MustAtoi(parts[0])
		y := conv.MustAtoi(parts[1])
		places[i] = grid.Coordinate{
			Row: y,
			Col: x,
		}
		if x < minCol {
			minCol = x
		}
		if x > maxCol {
			maxCol = x
		}
		if y < minRow {
			minRow = y
		}
		if y > maxRow {
			maxRow = y
		}
	}

	areas := make([]int, len(places))
	infinite := make([]bool, len(places))
	for i := minCol; i <= maxCol; i++ {
		for j := minRow; j <= maxRow; j++ {
			closest := findClosest(i, j, places)
			if closest != -1 {
				areas[closest]++
				if i == minCol || i == maxCol || j == minRow || j == maxRow {
					infinite[closest] = true
				}
			}
		}
	}

	maxArea := 0
	for i, area := range areas {
		if !infinite[i] && area > maxArea {
			maxArea = area
		}
	}
	fmt.Println(maxArea)

	// part 2
	regionSize := 0
	for i := minCol; i <= maxCol; i++ {
		for j := minRow; j <= maxRow; j++ {
			sum := 0
			for _, p := range places {
				sum += mathx.Abs(p.Row-j) + mathx.Abs(p.Col-i)
			}
			if sum < 10000 {
				regionSize++
			}
		}
	}
	fmt.Println(regionSize)
}

func findClosest(c, r int, places []grid.Coordinate) int {
	for i, p := range places {
		if p.Row == r && p.Col == c {
			return i
		}
	}

	minDist := math.MaxInt32
	minIndex := -1
	for i, p := range places {
		dist := mathx.Abs(p.Row-r) + mathx.Abs(p.Col-c)
		if dist < minDist {
			minDist = dist
			minIndex = i
		} else if dist == minDist {
			minIndex = -1
		}
	}

	return minIndex

}
