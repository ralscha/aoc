package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	width := len(lines[0])
	height := len(lines)
	row := 0
	col := 0
	treeCount := 0
	for row < height-1 {
		row++
		col += 3
		if lines[row][col%width] == '#' {
			treeCount++
		}
	}

	fmt.Println("Part 1", treeCount)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	width := len(lines[0])
	height := len(lines)

	slopes := []struct {
		right int
		down  int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeCount := 1
	for _, slope := range slopes {
		row := 0
		col := 0
		trees := 0
		for row < height-1 {
			row += slope.down
			col += slope.right
			if lines[row][col%width] == '#' {
				trees++
			}
		}
		treeCount *= trees
	}

	fmt.Println("Part 2", treeCount)
}
