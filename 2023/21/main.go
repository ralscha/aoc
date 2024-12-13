package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"math/big"
)

func main() {
	input, err := download.ReadInput(2023, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type state struct {
	pos   gridutil.Coordinate
	steps int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	// Find start position
	var start gridutil.Coordinate
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if val, exists := grid.Get(row, col); exists && val == 'S' {
				start = gridutil.Coordinate{Row: row, Col: col}
				grid.Set(row, col, '.') // Replace 'S' with '.' for consistent movement
				break
			}
		}
	}

	ways := make(map[state]int)
	ways[state{start, 0}] = 1

	directions := gridutil.Get4Directions()

	for step := 1; step <= 64; step++ {
		nextWays := make(map[state]int)
		for s, count := range ways {
			for _, dir := range directions {
				nextPos := gridutil.Coordinate{
					Row: s.pos.Row + dir.Row,
					Col: s.pos.Col + dir.Col,
				}
				if val, exists := grid.GetC(nextPos); exists && val == '.' {
					nextState := state{nextPos, step}
					nextWays[nextState] += count
				}
			}
		}
		ways = nextWays
	}

	var count uint64 = 0
	for state := range ways {
		if state.steps == 64 {
			count++
		}
	}

	fmt.Println(count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[bool](false)

	// Convert input to boolean grid (true for walkable, false for rocks)
	for row, line := range lines {
		for col, char := range line {
			grid.Set(row, col, char != '#')
		}
	}

	// Find start position
	locs := container.NewSet[gridutil.Coordinate]()
	for row, line := range lines {
		for col, char := range line {
			if char == 'S' {
				locs.Add(gridutil.Coordinate{Row: row, Col: col})
			}
		}
	}

	directions := gridutil.Get4Directions()
	var steps [3]*big.Int
	s := 0
	gridSize := len(lines)

	for i := 1; i <= gridSize*3+65; i++ {
		nlocs := container.NewSet[gridutil.Coordinate]()
		for _, loc := range locs.Values() {
			for _, dir := range directions {
				nextPos := gridutil.Coordinate{
					Row: loc.Row + dir.Row,
					Col: loc.Col + dir.Col,
				}
				if isValid(&grid, nextPos) {
					nlocs.Add(nextPos)
				}
			}
		}

		if i%gridSize == gridSize/2 {
			steps[s] = big.NewInt(int64(nlocs.Len()))
			s++
			if s == 3 {
				break
			}
		}
		locs = nlocs
	}

	result := calculateResult(big.NewInt(int64(26501365/len(lines))), steps)
	fmt.Println(result)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func isValid(grid *gridutil.Grid2D[bool], pos gridutil.Coordinate) bool {
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()
	height := maxRow - minRow + 1
	width := maxCol - minCol + 1

	// Map position to grid bounds using modulo
	mappedRow := mod(pos.Row-minRow, height) + minRow
	mappedCol := mod(pos.Col-minCol, width) + minCol

	if val, exists := grid.Get(mappedRow, mappedCol); exists {
		return val
	}
	return false
}

func calculateResult(n *big.Int, steps [3]*big.Int) *big.Int {
	a0 := steps[0]
	a1 := steps[1]
	a2 := steps[2]

	b0 := big.NewInt(0).Set(a0)
	b1 := big.NewInt(0).Sub(a1, a0)
	b2 := big.NewInt(0).Sub(a2, a1)

	// Calculate quadratic formula components
	c := big.NewInt(0).Mul(n, big.NewInt(0).Sub(n, big.NewInt(1)))
	c = big.NewInt(0).Div(c, big.NewInt(2))
	c = big.NewInt(0).Mul(c, big.NewInt(0).Sub(b2, b1))

	b := big.NewInt(0).Mul(b1, n)

	r := big.NewInt(0).Add(b0, b)
	return big.NewInt(0).Add(r, c)
}
