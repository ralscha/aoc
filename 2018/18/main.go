package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2018, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	for range 10 {
		grid = updateGrid(grid)
	}

	treeCount, lumberyardCount := countResources(grid)
	fmt.Println("Part 1", treeCount*lumberyardCount)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	seen := make(map[string]int)
	var history []gridutil.Grid2D[rune]

	for i := range 1000 {
		gridStr := grid.String()
		if prev, ok := seen[gridStr]; ok {
			cycleLen := i - prev
			remaining := 1000000000 - i
			index := prev + (remaining % cycleLen)

			finalGrid := history[index]
			treeCount, lumberyardCount := countResources(finalGrid)
			fmt.Println("Part 2", treeCount*lumberyardCount)
			return
		}
		seen[gridStr] = i
		history = append(history, grid.Copy())

		grid = updateGrid(grid)
	}
}

func updateGrid(grid gridutil.Grid2D[rune]) gridutil.Grid2D[rune] {
	nextGrid := grid.Copy()
	for r := 0; r < grid.Height(); r++ {
		for c := 0; c < grid.Width(); c++ {
			treeCount, lumberyardCount := countNeighbors(grid, r, c)

			current, _ := grid.Get(r, c)
			if current == '.' {
				if treeCount >= 3 {
					nextGrid.Set(r, c, '|')
				}
			} else if current == '|' {
				if lumberyardCount >= 3 {
					nextGrid.Set(r, c, '#')
				}
			} else if current == '#' {
				if lumberyardCount >= 1 && treeCount >= 1 {
					nextGrid.Set(r, c, '#')
				} else {
					nextGrid.Set(r, c, '.')
				}
			}
		}
	}
	return nextGrid
}

func countNeighbors(grid gridutil.Grid2D[rune], r, c int) (int, int) {
	treeCount := 0
	lumberyardCount := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			neighbor, ok := grid.Get(r+dr, c+dc)
			if ok {
				if neighbor == '|' {
					treeCount++
				} else if neighbor == '#' {
					lumberyardCount++
				}
			}
		}
	}
	return treeCount, lumberyardCount
}

func countResources(grid gridutil.Grid2D[rune]) (int, int) {
	treeCount := 0
	lumberyardCount := 0
	for r := 0; r < grid.Height(); r++ {
		for c := 0; c < grid.Width(); c++ {
			val, _ := grid.Get(r, c)
			if val == '|' {
				treeCount++
			} else if val == '#' {
				lumberyardCount++
			}
		}
	}
	return treeCount, lumberyardCount
}
