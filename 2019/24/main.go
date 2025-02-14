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
	lines := conv.SplitNewline(input)
	initialGrid := gridutil.NewCharGrid2D(lines)

	levels := make(map[int]*container.Set[gridutil.Coordinate])
	levels[0] = container.NewSet[gridutil.Coordinate]()

	for r := range 5 {
		for c := range 5 {
			if v, _ := initialGrid.Get(r, c); v == '#' {
				levels[0].Add(gridutil.Coordinate{Row: r, Col: c})
			}
		}
	}

	for minute := 0; minute < 200; minute++ {
		minLevel, maxLevel := 0, 0
		for level := range levels {
			if level < minLevel {
				minLevel = level
			}
			if level > maxLevel {
				maxLevel = level
			}
		}

		if levels[minLevel].Len() > 0 {
			levels[minLevel-1] = container.NewSet[gridutil.Coordinate]()
		}
		if levels[maxLevel].Len() > 0 {
			levels[maxLevel+1] = container.NewSet[gridutil.Coordinate]()
		}

		newLevels := make(map[int]*container.Set[gridutil.Coordinate])
		for level := minLevel - 1; level <= maxLevel+1; level++ {
			newLevels[level] = container.NewSet[gridutil.Coordinate]()
			for r := range 5 {
				for c := range 5 {
					if r == 2 && c == 2 {
						continue
					}

					coord := gridutil.Coordinate{Row: r, Col: c}
					adjacentBugs := countAdjacentBugs(levels, level, coord)

					if levels[level].Contains(coord) {
						if adjacentBugs == 1 {
							newLevels[level].Add(coord)
						}
					} else {
						if adjacentBugs == 1 || adjacentBugs == 2 {
							newLevels[level].Add(coord)
						}
					}
				}
			}
		}
		levels = newLevels
	}

	totalBugs := 0
	for _, level := range levels {
		totalBugs += len(level.Values())
	}

	fmt.Println("Part 2", totalBugs)
}

func countAdjacentBugs(levels map[int]*container.Set[gridutil.Coordinate], level int, pos gridutil.Coordinate) int {
	count := 0

	for _, dir := range gridutil.Get4Directions() {
		newRow := pos.Row + dir.Row
		newCol := pos.Col + dir.Col

		if newRow == 2 && newCol == 2 {
			continue
		}

		if levels[level] == nil {
			levels[level] = container.NewSet[gridutil.Coordinate]()
		}

		if newRow >= 0 && newRow < 5 && newCol >= 0 && newCol < 5 {
			if levels[level].Contains(gridutil.Coordinate{Row: newRow, Col: newCol}) {
				count++
			}
		}
	}

	if levels[level-1] == nil {
		levels[level-1] = container.NewSet[gridutil.Coordinate]()
	}

	if pos.Row == 0 && levels[level-1].Contains(gridutil.Coordinate{Row: 1, Col: 2}) {
		count++
	}
	if pos.Row == 4 && levels[level-1].Contains(gridutil.Coordinate{Row: 3, Col: 2}) {
		count++
	}
	if pos.Col == 0 && levels[level-1].Contains(gridutil.Coordinate{Row: 2, Col: 1}) {
		count++
	}
	if pos.Col == 4 && levels[level-1].Contains(gridutil.Coordinate{Row: 2, Col: 3}) {
		count++
	}

	if levels[level+1] == nil {
		levels[level+1] = container.NewSet[gridutil.Coordinate]()
	}

	if pos.Row == 1 && pos.Col == 2 {
		for c := range 5 {
			if levels[level+1].Contains(gridutil.Coordinate{Row: 0, Col: c}) {
				count++
			}
		}
	}
	if pos.Row == 3 && pos.Col == 2 {
		for c := range 5 {
			if levels[level+1].Contains(gridutil.Coordinate{Row: 4, Col: c}) {
				count++
			}
		}
	}
	if pos.Row == 2 && pos.Col == 1 {
		for r := range 5 {
			if levels[level+1].Contains(gridutil.Coordinate{Row: r, Col: 0}) {
				count++
			}
		}
	}
	if pos.Row == 2 && pos.Col == 3 {
		for r := range 5 {
			if levels[level+1].Contains(gridutil.Coordinate{Row: r, Col: 4}) {
				count++
			}
		}
	}

	return count
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
