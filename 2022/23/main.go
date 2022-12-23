package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2022/23/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

type position struct {
	row, col int
}

type elf struct {
	nextPosition *position
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	elves := make(map[position]*elf)

	for r, line := range lines {
		for i, c := range line {
			if c == '#' {
				pos := position{row: r, col: i}
				elves[pos] = &elf{}
			}
		}

	}

	var directions = []string{"n", "s", "w", "e"}

	round := 10
	for r := 1; r <= round; r++ {

		for pos, elf := range elves {
			var neighbors []bool
			hasNeighbor := false
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					checkPos := position{row: pos.row + i, col: pos.col + j}
					if elves[checkPos] != nil {
						neighbors = append(neighbors, true)
						hasNeighbor = true
					} else {
						neighbors = append(neighbors, false)
					}
				}
			}
			if !hasNeighbor {
				continue
			}
			for _, direction := range directions {
				if direction == "n" && !neighbors[0] && !neighbors[1] && !neighbors[2] {
					elf.nextPosition = &position{row: pos.row - 1, col: pos.col}
					break
				}
				if direction == "s" && !neighbors[5] && !neighbors[6] && !neighbors[7] {
					elf.nextPosition = &position{row: pos.row + 1, col: pos.col}
					break
				}
				if direction == "w" && !neighbors[0] && !neighbors[3] && !neighbors[5] {
					elf.nextPosition = &position{row: pos.row, col: pos.col - 1}
					break
				}
				if direction == "e" && !neighbors[2] && !neighbors[4] && !neighbors[7] {
					elf.nextPosition = &position{row: pos.row, col: pos.col + 1}
					break
				}
			}
		}

		// remove duplicates
		for _, elf := range elves {
			if elf.nextPosition == nil {
				continue
			}
			hasSame := false
			for _, elf2 := range elves {
				if elf != elf2 && elf2.nextPosition != nil && *elf.nextPosition == *elf2.nextPosition {
					hasSame = true
					elf2.nextPosition = nil
				}
			}
			if hasSame {
				elf.nextPosition = nil
			}
		}

		// move
		newElves := make(map[position]*elf)
		for pos, elf := range elves {
			if elf.nextPosition != nil {
				newElves[*elf.nextPosition] = elf
				elf.nextPosition = nil
			} else {
				newElves[pos] = elf
			}
		}
		elves = newElves

		directions = append(directions[1:], directions[0])
	}

	minRow, maxRow, minCol, maxCol := findMinMax(elves)
	fmt.Println((maxRow-minRow+1)*(maxCol-minCol+1) - len(elves))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	elves := make(map[position]*elf)

	for r, line := range lines {
		for i, c := range line {
			if c == '#' {
				pos := position{row: r, col: i}
				elves[pos] = &elf{}
			}
		}

	}

	var directions = []string{"n", "s", "w", "e"}

	round := 1
	for {

		for pos, elf := range elves {
			var neighbors []bool
			hasNeighbor := false
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					checkPos := position{row: pos.row + i, col: pos.col + j}
					if elves[checkPos] != nil {
						neighbors = append(neighbors, true)
						hasNeighbor = true
					} else {
						neighbors = append(neighbors, false)
					}
				}
			}
			if !hasNeighbor {
				continue
			}
			for _, direction := range directions {
				if direction == "n" && !neighbors[0] && !neighbors[1] && !neighbors[2] {
					elf.nextPosition = &position{row: pos.row - 1, col: pos.col}
					break
				}
				if direction == "s" && !neighbors[5] && !neighbors[6] && !neighbors[7] {
					elf.nextPosition = &position{row: pos.row + 1, col: pos.col}
					break
				}
				if direction == "w" && !neighbors[0] && !neighbors[3] && !neighbors[5] {
					elf.nextPosition = &position{row: pos.row, col: pos.col - 1}
					break
				}
				if direction == "e" && !neighbors[2] && !neighbors[4] && !neighbors[7] {
					elf.nextPosition = &position{row: pos.row, col: pos.col + 1}
					break
				}
			}
		}

		// remove duplicates
		for _, elf := range elves {
			if elf.nextPosition == nil {
				continue
			}
			hasSame := false
			for _, elf2 := range elves {
				if elf != elf2 && elf2.nextPosition != nil && *elf.nextPosition == *elf2.nextPosition {
					hasSame = true
					elf2.nextPosition = nil
				}
			}
			if hasSame {
				elf.nextPosition = nil
			}
		}

		// count next positions
		nextPos := 0
		for _, elf := range elves {
			if elf.nextPosition != nil {
				nextPos++
			}
		}

		if nextPos == 0 {
			fmt.Println(round)
			break
		}

		// move
		newElves := make(map[position]*elf)
		for pos, elf := range elves {
			if elf.nextPosition != nil {
				newElves[*elf.nextPosition] = elf
				elf.nextPosition = nil
			} else {
				newElves[pos] = elf
			}
		}
		elves = newElves

		directions = append(directions[1:], directions[0])
		round++
	}
}

func findMinMax(elves map[position]*elf) (int, int, int, int) {
	minRow := 1000
	maxRow := -1000
	minCol := 1000
	maxCol := -1000
	for pos := range elves {
		if pos.row < minRow {
			minRow = pos.row
		}
		if pos.row > maxRow {
			maxRow = pos.row
		}
		if pos.col < minCol {
			minCol = pos.col
		}
		if pos.col > maxCol {
			maxCol = pos.col
		}
	}
	return minRow, maxRow, minCol, maxCol
}
