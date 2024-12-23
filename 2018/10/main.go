package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type point struct {
	pos      gridutil.Coordinate
	velocity gridutil.Direction
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	points := make([]point, len(lines))
	for i, line := range lines {
		splitted := strings.Split(line, "=")
		last := strings.LastIndex(splitted[1], ">")
		pos := splitted[1][1:last]
		velocity := splitted[2][1 : len(splitted[2])-1]

		posSplitted := strings.Split(pos, ",")
		velocitySplitted := strings.Split(velocity, ",")

		points[i].pos = gridutil.Coordinate{
			Col: conv.MustAtoi(strings.TrimSpace(posSplitted[0])),
			Row: conv.MustAtoi(strings.TrimSpace(posSplitted[1])),
		}
		points[i].velocity = gridutil.Direction{
			Col: conv.MustAtoi(strings.TrimSpace(velocitySplitted[0])),
			Row: conv.MustAtoi(strings.TrimSpace(velocitySplitted[1])),
		}
	}

	minBoundingBox := math.MaxInt
	minBoundingBoxIndex := 0

	for i := range 100000 {
		minX := math.MaxInt
		maxX := math.MinInt
		minY := math.MaxInt
		maxY := math.MinInt

		for _, p := range points {
			x := p.pos.Col + i*p.velocity.Col
			y := p.pos.Row + i*p.velocity.Row

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}

		if maxX-minX+maxY-minY < minBoundingBox {
			minBoundingBox = maxX - minX + maxY - minY
			minBoundingBoxIndex = i
		}
	}

	fmt.Println("Part2:", minBoundingBoxIndex)

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for _, p := range points {
		x := p.pos.Col + minBoundingBoxIndex*p.velocity.Col
		y := p.pos.Row + minBoundingBoxIndex*p.velocity.Row

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	grid := gridutil.NewGrid2D[rune](false)
	grid.SetMinRowCol(minY, minX)
	grid.SetMaxRowCol(maxY, maxX)

	// Initialize grid with spaces
	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			grid.Set(row, col, ' ')
		}
	}

	// Plot points
	for _, p := range points {
		x := p.pos.Col + minBoundingBoxIndex*p.velocity.Col
		y := p.pos.Row + minBoundingBoxIndex*p.velocity.Row
		grid.Set(y, x, '*')
	}

	// Print grid
	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			if val, ok := grid.Get(row, col); ok {
				fmt.Printf("%c", val)
			}
		}
		fmt.Println()
	}
}
