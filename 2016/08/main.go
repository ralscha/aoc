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
	input, err := download.ReadInput(2016, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1and2(lines)
}

func part1and2(lines []string) {
	grid := gridutil.NewGrid2D[bool](false)
	grid.SetMaxRowCol(5, 49) // 6 rows (0-5), 50 columns (0-49)

	for _, line := range lines {
		spaceIndex := strings.Index(line, " ")
		command := line[0:spaceIndex]
		switch command {
		case "rect":
			rect := line[spaceIndex+1:]
			xIndex := strings.Index(rect, "x")
			col := conv.MustAtoi(rect[0:xIndex])
			row := conv.MustAtoi(rect[xIndex+1:])
			for y := 0; y < row; y++ {
				for x := 0; x < col; x++ {
					grid.Set(y, x, true)
				}
			}
		case "rotate":
			rotate := line[spaceIndex+1:]
			equalsIndex := strings.Index(rotate, "=")
			byIndex := strings.Index(rotate, " by ")
			axis := rotate[0 : equalsIndex-2]
			index := conv.MustAtoi(rotate[equalsIndex+1 : byIndex])
			amount := conv.MustAtoi(rotate[byIndex+4:])

			// Create a copy of the grid for rotation
			gridCopy := grid.Copy()

			switch axis {
			case "row":
				for x := 0; x < 50; x++ {
					val, _ := gridCopy.Get(index, x)
					grid.Set(index, (x+amount)%50, val)
				}
			case "column":
				for y := 0; y < 6; y++ {
					val, _ := gridCopy.Get(y, index)
					grid.Set((y+amount)%6, index, val)
				}
			}
		}
	}

	// Count lit pixels
	count := 0
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if val, _ := grid.Get(y, x); val {
				count++
			}
		}
	}

	// Display the grid
	for y := 0; y < 6; y++ {
		for x := 0; x < 50; x++ {
			if val, _ := grid.Get(y, x); val {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println("Part 1", count)
}
