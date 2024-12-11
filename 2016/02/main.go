package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2016, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	// Create keypad grid
	grid := gridutil.NewGrid2D[int](false)
	grid.SetMaxRowCol(2, 2)
	// Set keypad values
	grid.Set(0, 0, 1)
	grid.Set(0, 1, 2)
	grid.Set(0, 2, 3)
	grid.Set(1, 0, 4)
	grid.Set(1, 1, 5)
	grid.Set(1, 2, 6)
	grid.Set(2, 0, 7)
	grid.Set(2, 1, 8)
	grid.Set(2, 2, 9)

	pos := gridutil.Coordinate{Row: 1, Col: 1} // Start at 5
	for _, line := range lines {
		for _, c := range line {
			var dir gridutil.Direction
			switch c {
			case 'U':
				dir = gridutil.DirectionN
			case 'D':
				dir = gridutil.DirectionS
			case 'L':
				dir = gridutil.DirectionW
			case 'R':
				dir = gridutil.DirectionE
			}
			newPos := gridutil.Coordinate{Row: pos.Row + dir.Row, Col: pos.Col + dir.Col}
			if val, ok := grid.GetC(newPos); ok {
				pos = newPos
				_ = val // Value exists, position is valid
			}
		}
		val, _ := grid.GetC(pos)
		fmt.Print(val)
	}
	fmt.Println()
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	// Create keypad grid
	grid := gridutil.NewGrid2D[rune](false)
	grid.SetMaxRowCol(4, 4)
	// Set keypad values
	grid.Set(0, 2, '1')
	grid.Set(1, 1, '2')
	grid.Set(1, 2, '3')
	grid.Set(1, 3, '4')
	grid.Set(2, 0, '5')
	grid.Set(2, 1, '6')
	grid.Set(2, 2, '7')
	grid.Set(2, 3, '8')
	grid.Set(2, 4, '9')
	grid.Set(3, 1, 'A')
	grid.Set(3, 2, 'B')
	grid.Set(3, 3, 'C')
	grid.Set(4, 2, 'D')

	pos := gridutil.Coordinate{Row: 2, Col: 0} // Start at 5
	for _, line := range lines {
		for _, c := range line {
			var dir gridutil.Direction
			switch c {
			case 'U':
				dir = gridutil.DirectionN
			case 'D':
				dir = gridutil.DirectionS
			case 'L':
				dir = gridutil.DirectionW
			case 'R':
				dir = gridutil.DirectionE
			}
			newPos := gridutil.Coordinate{Row: pos.Row + dir.Row, Col: pos.Col + dir.Col}
			if val, ok := grid.GetC(newPos); ok && val != 0 {
				pos = newPos
			}
		}
		val, _ := grid.GetC(pos)
		fmt.Printf("%c", val)
	}
	fmt.Println()
}
