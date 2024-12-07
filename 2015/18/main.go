package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[bool](false)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				grid.Set(i, j, true)
			}
		}
	}

	rounds := 100
	for i := 0; i < rounds; i++ {
		newGrid := grid.Copy()
		minRow, maxRow := grid.GetMinMaxRow()
		minCol, maxCol := grid.GetMinMaxCol()

		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				isOn, _ := grid.Get(row, col)
				neighbors := grid.GetNeighbours8(row, col)
				neighborsOn := 0
				for _, n := range neighbors {
					if n {
						neighborsOn++
					}
				}

				if isOn {
					if neighborsOn != 2 && neighborsOn != 3 {
						newGrid.Set(row, col, false)
					}
				} else if neighborsOn == 3 {
					newGrid.Set(row, col, true)
				}
			}
		}
		grid = newGrid
	}

	fmt.Println(countOn(grid))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[bool](false)
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				grid.Set(i, j, true)
			}
		}
	}

	turnOnCorners(grid)

	rounds := 100
	for i := 0; i < rounds; i++ {
		newGrid := grid.Copy()
		minRow, maxRow := grid.GetMinMaxRow()
		minCol, maxCol := grid.GetMinMaxCol()

		for row := minRow; row <= maxRow; row++ {
			for col := minCol; col <= maxCol; col++ {
				isOn, _ := grid.Get(row, col)
				neighbors := grid.GetNeighbours8(row, col)
				neighborsOn := 0
				for _, n := range neighbors {
					if n {
						neighborsOn++
					}
				}

				if isOn {
					if neighborsOn != 2 && neighborsOn != 3 {
						newGrid.Set(row, col, false)
					}
				} else if neighborsOn == 3 {
					newGrid.Set(row, col, true)
				}
			}
		}

		turnOnCorners(newGrid)
		grid = newGrid
	}

	fmt.Println(countOn(grid))
}

func turnOnCorners(grid gridutil.Grid2D[bool]) {
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	grid.Set(minRow, minCol, true)
	grid.Set(minRow, maxCol, true)
	grid.Set(maxRow, minCol, true)
	grid.Set(maxRow, maxCol, true)
}

func countOn(grid gridutil.Grid2D[bool]) int {
	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if isOn, _ := grid.Get(row, col); isOn {
				count++
			}
		}
	}
	return count
}
