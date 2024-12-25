package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	var currentSchematic []string
	var schematics []*schematic

	for _, line := range lines {
		if line == "" && len(currentSchematic) > 0 {
			grid := gridutil.NewCharGrid2D(currentSchematic)
			schematics = append(schematics, &schematic{grid: grid})
			currentSchematic = nil
		} else if line != "" {
			currentSchematic = append(currentSchematic, line)
		}
	}
	if len(currentSchematic) > 0 {
		grid := gridutil.NewCharGrid2D(currentSchematic)
		schematics = append(schematics, &schematic{grid: grid})
	}

	var locks, keys []*schematic
	for _, s := range schematics {
		if isLockSchematic(s) {
			locks = append(locks, s)
		} else {
			keys = append(keys, s)
		}
	}

	validPairs := 0
	for _, lock := range locks {
		lockHeights := lock.getHeights(true)

		for _, key := range keys {
			keyHeights := key.getHeights(false)

			if canFitTogether(lockHeights, keyHeights) {
				validPairs++
			}
		}
	}

	fmt.Println("Part 1", validPairs)
}

type schematic struct {
	grid gridutil.Grid2D[rune]
}

func (s *schematic) getHeights(isLock bool) []int {
	if s.grid.Width() == 0 || s.grid.Height() == 0 {
		return nil
	}

	heights := make([]int, s.grid.Width())
	rows := s.grid.Height()

	for col := range s.grid.Width() {
		if isLock {
			found := false
			for row := range rows {
				val, _ := s.grid.Get(row, col)
				if val != '#' {
					heights[col] = row
					found = true
					break
				}
			}
			if !found {
				heights[col] = rows
			}
		} else {
			count := 0
			for row := rows - 1; row >= 0; row-- {
				val, _ := s.grid.Get(row, col)
				if val == '#' {
					count++
				} else {
					break
				}
			}
			heights[col] = count
		}
	}
	return heights
}

func isLockSchematic(s *schematic) bool {
	for col := range s.grid.Width() {
		val, _ := s.grid.Get(0, col)
		if val != '#' {
			return false
		}
	}

	for col := range s.grid.Width() {
		val, _ := s.grid.Get(s.grid.Height()-1, col)
		if val == '#' {
			return false
		}
	}
	return true
}

func canFitTogether(lockHeights, keyHeights []int) bool {
	for i := range len(lockHeights) {
		if lockHeights[i]+keyHeights[i] > 7 {
			return false
		}
	}
	return true
}
