package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

type fold struct {
	axis byte // 'x' or 'y'
	pos  int
}

func main() {
	input, err := download.ReadInput(2021, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	dots, folds := parseInput(input)

	// Only perform first fold
	if len(folds) > 0 {
		dots = performFold(dots, folds[0])
	}

	fmt.Println("Part 1", dots.Len())
}

func part2(input string) {
	dots, folds := parseInput(input)

	// Perform all folds
	for _, f := range folds {
		dots = performFold(dots, f)
	}

	// Create grid for visualization
	grid := gridutil.NewGrid2D[bool](false)
	maxX, maxY := 0, 0

	for _, dot := range dots.Values() {
		if dot.Col > maxX {
			maxX = dot.Col
		}
		if dot.Row > maxY {
			maxY = dot.Row
		}
		grid.SetC(dot, true)
	}

	// Print result
	fmt.Println("Final pattern:")
	for row := 0; row <= maxY; row++ {
		for col := 0; col <= maxX; col++ {
			if val, _ := grid.Get(row, col); val {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput(input string) (*container.Set[gridutil.Coordinate], []fold) {
	lines := conv.SplitNewline(input)
	dots := container.NewSet[gridutil.Coordinate]()
	var folds []fold

	// Parse input into dots and folds
	parsingDots := true
	for _, line := range lines {
		if line == "" {
			parsingDots = false
			continue
		}

		if parsingDots {
			parts := strings.Split(line, ",")
			x := conv.MustAtoi(parts[0])
			y := conv.MustAtoi(parts[1])
			dots.Add(gridutil.Coordinate{Row: y, Col: x})
		} else {
			parts := strings.Split(line[11:], "=") // Remove "fold along "
			folds = append(folds, fold{
				axis: parts[0][0], // 'x' or 'y'
				pos:  conv.MustAtoi(parts[1]),
			})
		}
	}

	return dots, folds
}

func performFold(dots *container.Set[gridutil.Coordinate], f fold) *container.Set[gridutil.Coordinate] {
	newDots := container.NewSet[gridutil.Coordinate]()

	for _, dot := range dots.Values() {
		var newDot gridutil.Coordinate
		if f.axis == 'x' && dot.Col > f.pos {
			// Fold left
			newDot = gridutil.Coordinate{
				Row: dot.Row,
				Col: f.pos - (dot.Col - f.pos),
			}
			newDots.Add(newDot)
		} else if f.axis == 'y' && dot.Row > f.pos {
			// Fold up
			newDot = gridutil.Coordinate{
				Row: f.pos - (dot.Row - f.pos),
				Col: dot.Col,
			}
			newDots.Add(newDot)
		} else {
			// Point stays the same
			newDots.Add(dot)
		}
	}

	return newDots
}
