package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	maxSeatID := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		rowStr := line[:7]
		colStr := line[7:]

		minRow := 0
		maxRow := 127
		for _, char := range rowStr {
			if char == 'F' {
				maxRow = (minRow + maxRow) / 2
			} else if char == 'B' {
				minRow = (minRow + maxRow + 1) / 2
			}
		}
		row := minRow

		minCol := 0
		maxCol := 7
		for _, char := range colStr {
			if char == 'L' {
				maxCol = (minCol + maxCol) / 2
			} else if char == 'R' {
				minCol = (minCol + maxCol + 1) / 2
			}
		}
		col := minCol

		seatID := row*8 + col
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println("Part 1", maxSeatID)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	seatIDs := container.NewSet[int]()
	minSeatID := 127*8 + 7
	maxSeatID := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		rowStr := line[:7]
		colStr := line[7:]

		minRow := 0
		maxRow := 127
		for _, char := range rowStr {
			if char == 'F' {
				maxRow = (minRow + maxRow) / 2
			} else if char == 'B' {
				minRow = (minRow + maxRow + 1) / 2
			}
		}
		row := minRow

		minCol := 0
		maxCol := 7
		for _, char := range colStr {
			if char == 'L' {
				maxCol = (minCol + maxCol) / 2
			} else if char == 'R' {
				minCol = (minCol + maxCol + 1) / 2
			}
		}
		col := minCol

		seatID := row*8 + col
		seatIDs.Add(seatID)
		if seatID < minSeatID {
			minSeatID = seatID
		}
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	for i := minSeatID; i <= maxSeatID; i++ {
		if !seatIDs.Contains(i) {
			fmt.Println("Part 2", i)
			return
		}
	}
}
