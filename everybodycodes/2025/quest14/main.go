package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/internal/conv"
	"aoc/internal/gridutil"
)

var diagonals = []gridutil.Direction{
	gridutil.DirectionNW,
	gridutil.DirectionNE,
	gridutil.DirectionSW,
	gridutil.DirectionSE,
}

func main() {
	partI()
	partII()
	partIII()
}

func readInput(filename string) gridutil.Grid2D[rune] {
	input, _ := os.ReadFile(filename)
	cleaned := strings.ReplaceAll(string(input), "\r", "")
	lines := conv.SplitNewline(cleaned)
	return gridutil.NewCharGrid2D(lines)
}

func evolveGrid(grid gridutil.Grid2D[rune], minRow, maxRow, minCol, maxCol int) gridutil.Grid2D[rune] {
	newGrid := gridutil.NewGrid2D[rune](false)

	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			current, _ := grid.Get(r, c)

			activeCount := 0
			for _, dir := range diagonals {
				nr, nc := r+dir.Row, c+dir.Col
				if nr >= minRow && nr <= maxRow && nc >= minCol && nc <= maxCol {
					if val, _ := grid.Get(nr, nc); val == '#' {
						activeCount++
					}
				}
			}

			if current == '#' {
				if activeCount%2 == 1 {
					newGrid.Set(r, c, '#')
				} else {
					newGrid.Set(r, c, '.')
				}
			} else {
				if activeCount%2 == 0 {
					newGrid.Set(r, c, '#')
				} else {
					newGrid.Set(r, c, '.')
				}
			}
		}
	}

	return newGrid
}

func countActive(grid gridutil.Grid2D[rune], minRow, maxRow, minCol, maxCol int) int {
	count := 0
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if val, _ := grid.Get(r, c); val == '#' {
				count++
			}
		}
	}
	return count
}

func gridToString(grid gridutil.Grid2D[rune], minRow, maxRow, minCol, maxCol int) string {
	var sb strings.Builder
	for r := minRow; r <= maxRow; r++ {
		for c := minCol; c <= maxCol; c++ {
			if val, _ := grid.Get(r, c); val == '#' {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

func partI() {
	grid := readInput("input1")

	totalActive := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for range 10 {
		grid = evolveGrid(grid, minRow, maxRow, minCol, maxCol)
		totalActive += countActive(grid, minRow, maxRow, minCol, maxCol)
	}

	fmt.Println(totalActive)
}

func partII() {
	grid := readInput("input2")

	totalActive := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for range 2025 {
		grid = evolveGrid(grid, minRow, maxRow, minCol, maxCol)
		totalActive += countActive(grid, minRow, maxRow, minCol, maxCol)
	}

	fmt.Println(totalActive)
}

func patternMatches(grid, pattern gridutil.Grid2D[rune], centerOffset int) bool {
	for r := range 8 {
		for c := range 8 {
			gridVal, _ := grid.Get(centerOffset+r, centerOffset+c)
			patternVal, _ := pattern.Get(r, c)
			if gridVal != patternVal {
				return false
			}
		}
	}
	return true
}

func partIII() {
	pattern := readInput("input3")

	gridSize := 34
	grid := gridutil.NewGrid2D[rune](false)
	for r := range gridSize {
		for c := range gridSize {
			grid.Set(r, c, '.')
		}
	}

	minRow, maxRow := 0, gridSize-1
	minCol, maxCol := 0, gridSize-1

	centerOffset := (gridSize - 8) / 2

	type matchInfo struct {
		round       int
		activeCount int
	}
	seen := make(map[string]int)
	matches := make([]matchInfo, 0)

	targetRounds := 1000000000

	for round := 1; round <= targetRounds; round++ {
		grid = evolveGrid(grid, minRow, maxRow, minCol, maxCol)

		if patternMatches(grid, pattern, centerOffset) {
			activeCount := countActive(grid, minRow, maxRow, minCol, maxCol)
			matches = append(matches, matchInfo{round: round, activeCount: activeCount})
		}

		stateKey := gridToString(grid, minRow, maxRow, minCol, maxCol)

		if prevRound, found := seen[stateKey]; found {
			cycleLen := round - prevRound

			var cycleMatches []matchInfo
			for _, m := range matches {
				if m.round > prevRound && m.round <= round {
					cycleMatches = append(cycleMatches, m)
				}
			}

			totalActive := 0

			for _, m := range matches {
				if m.round <= prevRound {
					totalActive += m.activeCount
				}
			}

			remaining := targetRounds - prevRound
			fullCycles := remaining / cycleLen
			remainder := remaining % cycleLen

			cycleSum := 0
			for _, m := range cycleMatches {
				cycleSum += m.activeCount
			}
			totalActive += fullCycles * cycleSum

			for _, m := range cycleMatches {
				offsetInCycle := m.round - prevRound
				if offsetInCycle <= remainder {
					totalActive += m.activeCount
				}
			}

			fmt.Println(totalActive)
			return
		}
		seen[stateKey] = round
	}

	totalActive := 0
	for _, m := range matches {
		totalActive += m.activeCount
	}
	fmt.Println(totalActive)
}
