package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	simulate := func(current gridutil.Grid2D[rune]) gridutil.Grid2D[rune] {
		next := current.Copy()
		changed := false
		for r := range grid.Height() {
			for c := range grid.Width() {
				if val, ok := current.Get(r, c); ok {
					if val == '.' {
						continue
					}
					occupiedNeighbors := 0
					for _, neighbor := range current.GetNeighbours8(r, c) {
						if neighbor == '#' {
							occupiedNeighbors++
						}
					}

					if val == 'L' && occupiedNeighbors == 0 {
						next.Set(r, c, '#')
						changed = true
					} else if val == '#' && occupiedNeighbors >= 4 {
						next.Set(r, c, 'L')
						changed = true
					}
				}
			}
		}
		if !changed {
			return gridutil.Grid2D[rune]{}
		}
		return next
	}

	current := grid
	for {
		next := simulate(current)
		if next.Count() == 0 {
			break
		}
		current = next
	}

	occupiedCount := 0
	for r := range grid.Height() {
		for c := range grid.Width() {
			if val, ok := current.Get(r, c); ok && val == '#' {
				occupiedCount++
			}
		}
	}

	fmt.Println("Part 1", occupiedCount)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	visibleOccupied := func(g gridutil.Grid2D[rune], r, c int) int {
		occupied := 0
		for _, dir := range gridutil.Get8Directions() {
			row, col := r+dir.Row, c+dir.Col
			for {
				if val, ok := g.Get(row, col); ok {
					if val == '#' {
						occupied++
						break
					} else if val == 'L' {
						break
					}
				} else {
					break
				}
				row += dir.Row
				col += dir.Col
			}
		}
		return occupied
	}

	simulate := func(current gridutil.Grid2D[rune]) gridutil.Grid2D[rune] {
		next := current.Copy()
		changed := false
		for r := range grid.Height() {
			for c := range grid.Width() {
				if val, ok := current.Get(r, c); ok {
					if val == '.' {
						continue
					}
					occupiedVisible := visibleOccupied(current, r, c)

					if val == 'L' && occupiedVisible == 0 {
						next.Set(r, c, '#')
						changed = true
					} else if val == '#' && occupiedVisible >= 5 {
						next.Set(r, c, 'L')
						changed = true
					}
				}
			}
		}
		if !changed {
			return gridutil.Grid2D[rune]{}
		}
		return next
	}

	current := grid
	for {
		next := simulate(current)
		if next.Count() == 0 {
			break
		}
		current = next
	}

	occupiedCount := 0
	for r := range grid.Height() {
		for c := range grid.Width() {
			if val, ok := current.Get(r, c); ok && val == '#' {
				occupiedCount++
			}
		}
	}

	fmt.Println("Part 2", occupiedCount)
}
