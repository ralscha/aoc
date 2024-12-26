package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
)

func main() {
	input, err := download.ReadInput(2019, 24)
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)
	seen := container.NewSet[string]()

	for {
		gridStr := grid.String()
		if seen.Contains(gridStr) {
			biodiversity := calculateBiodiversity(grid)
			fmt.Println("Part 1", biodiversity)
			return
		}
		seen.Add(gridStr)
		nextGrid := gridutil.NewGrid2D[rune](false)
		for r := range 5 {
			for c := range 5 {
				adjacentBugs := 0
				v, _ := grid.Get(r-1, c)
				if r > 0 && v == '#' {
					adjacentBugs++
				}
				v, _ = grid.Get(r+1, c)
				if r < 4 && v == '#' {
					adjacentBugs++
				}
				v, _ = grid.Get(r, c-1)
				if c > 0 && v == '#' {
					adjacentBugs++
				}
				v, _ = grid.Get(r, c+1)
				if c < 4 && v == '#' {
					adjacentBugs++
				}

				v, _ = grid.Get(r, c)
				if v == '#' {
					if adjacentBugs == 1 {
						nextGrid.Set(r, c, '#')
					} else {
						nextGrid.Set(r, c, '.')

					}
				} else {
					if adjacentBugs == 1 || adjacentBugs == 2 {
						nextGrid.Set(r, c, '#')
					} else {
						nextGrid.Set(r, c, '.')
					}
				}
			}
		}
		grid = nextGrid
	}
}

func part2(input string) {
	fmt.Println("Part 2:", "not yet implemented")
}

func calculateBiodiversity(grid gridutil.Grid2D[rune]) int {
	biodiversity := 0
	power := 1
	for r := range 5 {
		for c := range 5 {
			v, _ := grid.Get(r, c)
			if v == '#' {
				biodiversity += power
			}
			power *= 2
		}
	}
	return biodiversity
}
