package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	partI()
	partII()
	partIII()
}

func partI() {
	cleanLines := readAndCleanInput("input1")
	grid := gridutil.NewNumberGrid2D(cleanLines)

	start := gridutil.Coordinate{Row: 0, Col: 0}
	reachable := computeReachable(grid, []gridutil.Coordinate{start})

	fmt.Println(reachable.Len())
}

func partII() {
	cleanLines := readAndCleanInput("input2")
	grid := gridutil.NewNumberGrid2D(cleanLines)

	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	start1 := gridutil.Coordinate{Row: minRow, Col: minCol}
	start2 := gridutil.Coordinate{Row: maxRow, Col: maxCol}

	reachable := computeReachable(grid, []gridutil.Coordinate{start1, start2})

	fmt.Println(reachable.Len())
}

func partIII() {
	cleanLines := readAndCleanInput("input3")
	grid := gridutil.NewNumberGrid2D(cleanLines)

	var candidates []gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			coord := gridutil.Coordinate{Row: r, Col: c}
			currVal, ok := grid.GetC(coord)
			if !ok {
				continue
			}

			isLocalMax := true
			for _, dir := range gridutil.Get4Directions() {
				next := gridutil.Coordinate{Row: r + dir.Row, Col: c + dir.Col}
				if val, ok := grid.GetC(next); ok {
					if val > currVal {
						isLocalMax = false
						break
					}
				}
			}

			if isLocalMax {
				candidates = append(candidates, coord)
			}
		}
	}

	maxSize1 := 0
	var r1 *container.Set[gridutil.Coordinate]

	for _, c := range candidates {
		r := computeReachable(grid, []gridutil.Coordinate{c})
		if r.Len() > maxSize1 {
			maxSize1 = r.Len()
			r1 = r
		}
	}

	maxSize2 := 0
	var r2 *container.Set[gridutil.Coordinate]

	for _, c := range candidates {
		r := computeReachable(grid, []gridutil.Coordinate{c})
		newCount := 0
		for _, coord := range r.Values() {
			if !r1.Contains(coord) {
				newCount++
			}
		}
		if newCount > maxSize2 {
			maxSize2 = newCount
			r2 = r
		}
	}

	maxSize3 := 0
	var r3 *container.Set[gridutil.Coordinate]

	r1r2 := unionSets(r1, r2)

	for _, c := range candidates {
		r := computeReachable(grid, []gridutil.Coordinate{c})
		newCount := 0
		for _, coord := range r.Values() {
			if !r1r2.Contains(coord) {
				newCount++
			}
		}
		if newCount > maxSize3 {
			maxSize3 = newCount
			r3 = r
		}
	}

	finalSet := unionSets(r1, r2, r3)

	fmt.Println(finalSet.Len())
}

func readAndCleanInput(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("reading %s failed: %v", filename, err)
	}

	lines := conv.SplitNewline(string(input))
	var cleanLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			cleanLines = append(cleanLines, trimmed)
		}
	}
	return cleanLines
}

func computeReachable(grid gridutil.Grid2D[int], starts []gridutil.Coordinate) *container.Set[gridutil.Coordinate] {
	queue := make([]gridutil.Coordinate, len(starts))
	copy(queue, starts)
	visited := container.NewSet[gridutil.Coordinate]()
	for _, start := range starts {
		visited.Add(start)
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		currVal, _ := grid.GetC(curr)

		for _, dir := range gridutil.Get4Directions() {
			next := gridutil.Coordinate{Row: curr.Row + dir.Row, Col: curr.Col + dir.Col}
			if val, ok := grid.GetC(next); ok {
				if !visited.Contains(next) && currVal >= val {
					visited.Add(next)
					queue = append(queue, next)
				}
			}
		}
	}

	return visited
}

func unionSets(sets ...*container.Set[gridutil.Coordinate]) *container.Set[gridutil.Coordinate] {
	result := container.NewSet[gridutil.Coordinate]()
	for _, set := range sets {
		for _, coord := range set.Values() {
			result.Add(coord)
		}
	}
	return result
}
