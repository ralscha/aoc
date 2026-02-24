package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"sort"
)

func main() {
	input, err := download.ReadInput(2025, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parse(input string) (int, int, int, map[int][]int) {
	grid := gridutil.NewCharGrid2D(conv.SplitNewline(input))
	_, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	startCol := 0
	splitterRows := make(map[int][]int)
	for r := 0; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if v, ok := grid.Get(r, c); ok {
				switch v {
				case 'S':
					startCol = c
				case '^':
					splitterRows[c] = append(splitterRows[c], r)
				}
			}
		}
	}
	return startCol, minCol, maxCol, splitterRows
}

func part1(input string) {
	startCol, minCol, maxCol, splitterRows := parse(input)

	hitSplitters := container.NewSet[gridutil.Coordinate]()
	queue := []gridutil.Coordinate{{Row: 0, Col: startCol}}

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]

		if b.Col < minCol || b.Col > maxCol {
			continue
		}

		splRow := nextSplitter(b.Col, b.Row, splitterRows)
		if splRow != -1 {
			splitterKey := gridutil.Coordinate{Row: splRow, Col: b.Col}
			if !hitSplitters.Contains(splitterKey) {
				hitSplitters.Add(splitterKey)
				queue = append(queue, gridutil.Coordinate{Row: splRow, Col: b.Col - 1}, gridutil.Coordinate{Row: splRow, Col: b.Col + 1})
			}
		}
	}

	fmt.Println("Part 1", hitSplitters.Len())
}

func part2(input string) {
	startCol, minCol, maxCol, splitterRows := parse(input)
	memo := make(map[gridutil.Coordinate]int64)
	fmt.Println("Part 2", countTimelines(startCol, 0, minCol, maxCol, memo, splitterRows))
}

func nextSplitter(col, fromRow int, splitterRows map[int][]int) int {
	rows := splitterRows[col]
	idx := sort.SearchInts(rows, fromRow)
	if idx < len(rows) {
		return rows[idx]
	}
	return -1
}

func countTimelines(col, fromRow, minCol, maxCol int, memo map[gridutil.Coordinate]int64, splitterRows map[int][]int) int64 {
	if col < minCol || col > maxCol {
		return 0
	}
	s := gridutil.Coordinate{Row: fromRow, Col: col}
	if v, ok := memo[s]; ok {
		return v
	}

	splRow := nextSplitter(col, fromRow, splitterRows)
	if splRow == -1 {
		return 1
	}

	result := countTimelines(col-1, splRow, minCol, maxCol, memo, splitterRows) +
		countTimelines(col+1, splRow, minCol, maxCol, memo, splitterRows)
	memo[s] = result
	return result
}
