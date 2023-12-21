package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
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

type pos struct {
	x, y int
}

var directions = []pos{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type state struct {
	pos   pos
	steps int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	var garden [][]bool
	var startX, startY int

	row := 0
	for _, line := range lines {
		gardenRow := make([]bool, len(line))
		for col, char := range line {
			if char == 'S' {
				startX, startY = row, col
				gardenRow[col] = true
			} else {
				gardenRow[col] = char == '.'
			}
		}
		garden = append(garden, gardenRow)
		row++
	}

	ways := make(map[state]int)
	ways[state{pos{startX, startY}, 0}] = 1

	for step := 1; step <= 64; step++ {
		nextWays := make(map[state]int)
		for s, count := range ways {
			for _, dir := range directions {
				nextPos := pos{s.pos.x + dir.x, s.pos.y + dir.y}
				if nextPos.x >= 0 && nextPos.x < len(garden) && nextPos.y >= 0 && nextPos.y < len(garden[0]) && garden[nextPos.x][nextPos.y] {
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
	grid := make([][]bool, len(lines))
	for i, line := range lines {
		grid[i] = make([]bool, len(line))
		for j, char := range line {
			if char != '#' {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	locs := make(map[pos]bool)
	for y, line := range lines {
		for x, char := range line {
			if char == 'S' {
				locs[pos{y: y, x: x}] = true
			}
		}
	}

	var steps [3]*big.Int
	s := 0
	for i := 1; i <= len(grid)*3+65; i++ {
		nlocs := make(map[pos]bool)
		for l := range locs {
			y, x := l.y, l.x
			if isValid(grid, y-1, x) {
				nlocs[pos{y: y - 1, x: x}] = true
			}
			if isValid(grid, y+1, x) {
				nlocs[pos{y: y + 1, x: x}] = true
			}
			if isValid(grid, y, x-1) {
				nlocs[pos{y: y, x: x - 1}] = true
			}
			if isValid(grid, y, x+1) {
				nlocs[pos{y: y, x: x + 1}] = true
			}
		}
		if i%len(grid) == len(grid)/2 {
			steps[s] = big.NewInt(int64(len(nlocs)))
			s++
			if s == 3 {
				break
			}
		}
		locs = nlocs
	}

	result := f(big.NewInt(int64(26501365/len(grid))), steps)
	fmt.Println(result)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func isValid(grid [][]bool, y, x int) bool {
	return grid[mod(y, len(grid))][mod(x, len(grid[0]))]
}

func f(n *big.Int, steps [3]*big.Int) *big.Int {
	a0 := steps[0]
	a1 := steps[1]
	a2 := steps[2]

	b0 := big.NewInt(0).Set(a0)
	b1 := big.NewInt(0).Sub(a1, a0)
	b2 := big.NewInt(0).Sub(a2, a1)

	c := big.NewInt(0).Mul(n, big.NewInt(0).Sub(n, big.NewInt(1)))
	c = big.NewInt(0).Div(c, big.NewInt(2))
	c = big.NewInt(0).Mul(c, big.NewInt(0).Sub(b2, b1))

	b := big.NewInt(0).Mul(b1, n)

	r := big.NewInt(0).Add(b0, b)
	return big.NewInt(0).Add(r, c)
}
