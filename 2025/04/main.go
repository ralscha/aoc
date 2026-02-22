package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2025, 4)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			val, ok := grid.Get(row, col)
			if !ok || val != '@' {
				continue
			}
			neighbors := grid.GetNeighbours8(row, col)
			atCount := 0
			for _, n := range neighbors {
				if n == '@' {
					atCount++
				}
			}
			if atCount < 4 {
				count++
			}
		}
	}
	fmt.Println("Part 1:", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	total := 0
	for {
		minRow, maxRow := grid.GetMinMaxRow()
		minCol, maxCol := grid.GetMinMaxCol()
		var toRemove []gridutil.Coordinate
		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				val, ok := grid.Get(row, col)
				if !ok || val != '@' {
					continue
				}
				neighbors := grid.GetNeighbours8(row, col)
				atCount := 0
				for _, n := range neighbors {
					if n == '@' {
						atCount++
					}
				}
				if atCount < 4 {
					toRemove = append(toRemove, gridutil.Coordinate{Row: row, Col: col})
				}
			}
		}
		if len(toRemove) == 0 {
			break
		}
		for _, c := range toRemove {
			grid.SetC(c, '.')
		}
		total += len(toRemove)
	}
	fmt.Println("Part 2:", total)
}
