package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[int](false)

	for _, line := range lines {
		splitted := strings.Split(line, " -> ")
		startSplitted := strings.Split(splitted[0], ",")
		endSplitted := strings.Split(splitted[1], ",")

		start := gridutil.Coordinate{
			Col: conv.MustAtoi(startSplitted[0]),
			Row: conv.MustAtoi(startSplitted[1]),
		}
		end := gridutil.Coordinate{
			Col: conv.MustAtoi(endSplitted[0]),
			Row: conv.MustAtoi(endSplitted[1]),
		}

		if start.Col == end.Col {
			startY, endY := start.Row, end.Row
			if startY > endY {
				startY, endY = endY, startY
			}
			for y := startY; y <= endY; y++ {
				val, _ := grid.Get(y, start.Col)
				grid.Set(y, start.Col, val+1)
			}
		} else if start.Row == end.Row {
			startX, endX := start.Col, end.Col
			if startX > endX {
				startX, endX = endX, startX
			}
			for x := startX; x <= endX; x++ {
				val, _ := grid.Get(start.Row, x)
				grid.Set(start.Row, x, val+1)
			}
		}
	}

	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if val, ok := grid.Get(row, col); ok && val >= 2 {
				count++
			}
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[int](false)

	for _, line := range lines {
		splitted := strings.Split(line, " -> ")
		startSplitted := strings.Split(splitted[0], ",")
		endSplitted := strings.Split(splitted[1], ",")

		start := gridutil.Coordinate{
			Col: conv.MustAtoi(startSplitted[0]),
			Row: conv.MustAtoi(startSplitted[1]),
		}
		end := gridutil.Coordinate{
			Col: conv.MustAtoi(endSplitted[0]),
			Row: conv.MustAtoi(endSplitted[1]),
		}

		if start.Col == end.Col {
			startY, endY := start.Row, end.Row
			if startY > endY {
				startY, endY = endY, startY
			}
			for y := startY; y <= endY; y++ {
				val, _ := grid.Get(y, start.Col)
				grid.Set(y, start.Col, val+1)
			}
		} else if start.Row == end.Row {
			startX, endX := start.Col, end.Col
			if startX > endX {
				startX, endX = endX, startX
			}
			for x := startX; x <= endX; x++ {
				val, _ := grid.Get(start.Row, x)
				grid.Set(start.Row, x, val+1)
			}
		} else {
			// Diagonal lines
			dx := 1
			if start.Col > end.Col {
				dx = -1
			}
			dy := 1
			if start.Row > end.Row {
				dy = -1
			}

			x, y := start.Col, start.Row
			for {
				val, _ := grid.Get(y, x)
				grid.Set(y, x, val+1)
				if x == end.Col {
					break
				}
				x += dx
				y += dy
			}
		}
	}

	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if val, ok := grid.Get(row, col); ok && val >= 2 {
				count++
			}
		}
	}

	fmt.Println("Part 2", count)
}
