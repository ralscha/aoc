package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2022, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type elf struct {
	nextPosition *gridutil.Coordinate
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	elves := make(map[gridutil.Coordinate]*elf)

	for r, line := range lines {
		for i, c := range line {
			if c == '#' {
				pos := gridutil.Coordinate{Row: r, Col: i}
				elves[pos] = &elf{}
			}
		}
	}

	directions := []string{"n", "s", "w", "e"}
	neighborOffsets := getNeighborOffsets()

	round := 10
	for r := 1; r <= round; r++ {
		proposeMovements(elves, directions, neighborOffsets)
		removeDuplicateProposals(elves)
		elves = moveElves(elves)
		directions = append(directions[1:], directions[0])
	}

	minRow, maxRow, minCol, maxCol := findMinMax(elves)
	fmt.Println((maxRow-minRow+1)*(maxCol-minCol+1) - len(elves))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	elves := make(map[gridutil.Coordinate]*elf)

	for r, line := range lines {
		for i, c := range line {
			if c == '#' {
				pos := gridutil.Coordinate{Row: r, Col: i}
				elves[pos] = &elf{}
			}
		}
	}

	directions := []string{"n", "s", "w", "e"}
	neighborOffsets := getNeighborOffsets()

	round := 1
	for {
		proposeMovements(elves, directions, neighborOffsets)
		removeDuplicateProposals(elves)

		// Count next positions
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

		elves = moveElves(elves)
		directions = append(directions[1:], directions[0])
		round++
	}
}

func getNeighborOffsets() []gridutil.Direction {
	return []gridutil.Direction{
		{Row: -1, Col: -1}, {Row: -1, Col: 0}, {Row: -1, Col: 1}, // NW, N, NE
		{Row: 0, Col: -1}, {Row: 0, Col: 1}, // W, E
		{Row: 1, Col: -1}, {Row: 1, Col: 0}, {Row: 1, Col: 1}, // SW, S, SE
	}
}

func proposeMovements(elves map[gridutil.Coordinate]*elf, directions []string, neighborOffsets []gridutil.Direction) {
	for pos, elf := range elves {
		var neighbors []bool
		hasNeighbor := false
		for _, offset := range neighborOffsets {
			checkPos := gridutil.Coordinate{Row: pos.Row + offset.Row, Col: pos.Col + offset.Col}
			if elves[checkPos] != nil {
				neighbors = append(neighbors, true)
				hasNeighbor = true
			} else {
				neighbors = append(neighbors, false)
			}
		}
		if !hasNeighbor {
			continue
		}
		for _, direction := range directions {
			if direction == "n" && !neighbors[0] && !neighbors[1] && !neighbors[2] {
				elf.nextPosition = &gridutil.Coordinate{Row: pos.Row - 1, Col: pos.Col}
				break
			}
			if direction == "s" && !neighbors[5] && !neighbors[6] && !neighbors[7] {
				elf.nextPosition = &gridutil.Coordinate{Row: pos.Row + 1, Col: pos.Col}
				break
			}
			if direction == "w" && !neighbors[0] && !neighbors[3] && !neighbors[5] {
				elf.nextPosition = &gridutil.Coordinate{Row: pos.Row, Col: pos.Col - 1}
				break
			}
			if direction == "e" && !neighbors[2] && !neighbors[4] && !neighbors[7] {
				elf.nextPosition = &gridutil.Coordinate{Row: pos.Row, Col: pos.Col + 1}
				break
			}
		}
	}
}

func removeDuplicateProposals(elves map[gridutil.Coordinate]*elf) {
	proposalCount := make(map[gridutil.Coordinate]int)
	for _, elf := range elves {
		if elf.nextPosition != nil {
			proposalCount[*elf.nextPosition]++
		}
	}

	for _, elf := range elves {
		if elf.nextPosition != nil && proposalCount[*elf.nextPosition] > 1 {
			elf.nextPosition = nil
		}
	}
}

func moveElves(elves map[gridutil.Coordinate]*elf) map[gridutil.Coordinate]*elf {
	newElves := make(map[gridutil.Coordinate]*elf)
	for pos, elf := range elves {
		if elf.nextPosition != nil {
			newElves[*elf.nextPosition] = elf
			elf.nextPosition = nil
		} else {
			newElves[pos] = elf
		}
	}
	return newElves
}

func findMinMax(elves map[gridutil.Coordinate]*elf) (int, int, int, int) {
	minRow, maxRow := 1000, -1000
	minCol, maxCol := 1000, -1000
	for pos := range elves {
		if pos.Row < minRow {
			minRow = pos.Row
		}
		if pos.Row > maxRow {
			maxRow = pos.Row
		}
		if pos.Col < minCol {
			minCol = pos.Col
		}
		if pos.Col > maxCol {
			maxCol = pos.Col
		}
	}
	return minRow, maxRow, minCol, maxCol
}
