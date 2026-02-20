package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type grid struct {
	cells [][]string
	size  int
}

func newGrid(rows []string) *grid {
	size := len(rows)
	cells := make([][]string, size)
	for i := range cells {
		cells[i] = make([]string, size)
		for j, c := range rows[i] {
			cells[i][j] = string(c)
		}
	}
	return &grid{cells: cells, size: size}
}

func (g *grid) String() string {
	rows := make([]string, g.size)
	for i, row := range g.cells {
		rows[i] = strings.Join(row, "")
	}
	return strings.Join(rows, "/")
}

func (g *grid) rotate() *grid {
	newCells := make([][]string, g.size)
	for i := range newCells {
		newCells[i] = make([]string, g.size)
		for j := range newCells[i] {
			newCells[i][j] = g.cells[g.size-1-j][i]
		}
	}
	return &grid{cells: newCells, size: g.size}
}

func (g *grid) flipHorizontal() *grid {
	newCells := make([][]string, g.size)
	for i := range newCells {
		newCells[i] = make([]string, g.size)
		for j := range newCells[i] {
			newCells[i][j] = g.cells[i][g.size-1-j]
		}
	}
	return &grid{cells: newCells, size: g.size}
}

func (g *grid) getAllTransformations() []string {
	seen := container.NewSet[string]()

	current := g
	for range 4 {
		seen.Add(current.String())
		current = current.rotate()
	}

	current = g.flipHorizontal()
	for range 4 {
		seen.Add(current.String())
		current = current.rotate()
	}

	result := make([]string, 0, seen.Len())
	for _, k := range seen.Values() {
		result = append(result, k)
	}
	return result
}

func (g *grid) subGrid(row, col, size int) *grid {
	sub := make([][]string, size)
	for i := range sub {
		sub[i] = make([]string, size)
		for j := range sub[i] {
			sub[i][j] = g.cells[row+i][col+j]
		}
	}
	return &grid{cells: sub, size: size}
}

func (g *grid) countOn() int {
	count := 0
	for _, row := range g.cells {
		for _, cell := range row {
			if cell == "#" {
				count++
			}
		}
	}
	return count
}

type Rules struct {
	patterns map[string][]string
}

func NewRules(input string) (*Rules, error) {
	lines := conv.SplitNewline(input)
	patterns := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid rule format: %s", line)
		}

		from := newGrid(strings.Split(parts[0], "/"))
		to := strings.Split(parts[1], "/")

		for _, transform := range from.getAllTransformations() {
			patterns[transform] = to
		}
	}

	return &Rules{patterns: patterns}, nil
}

func (r *Rules) apply(g *grid) *grid {
	size := g.size
	var newSize, subSize, newSubSize int

	if size%2 == 0 {
		subSize = 2
		newSubSize = 3
	} else {
		subSize = 3
		newSubSize = 4
	}

	numSubs := size / subSize
	newSize = numSubs * newSubSize

	newCells := make([][]string, newSize)
	for i := range newCells {
		newCells[i] = make([]string, newSize)
	}

	for i := range numSubs {
		for j := range numSubs {
			sub := g.subGrid(i*subSize, j*subSize, subSize)
			var output []string
			for _, pattern := range sub.getAllTransformations() {
				if out, ok := r.patterns[pattern]; ok {
					output = out
					break
				}
			}

			for row := 0; row < newSubSize; row++ {
				for col := 0; col < newSubSize; col++ {
					newCells[i*newSubSize+row][j*newSubSize+col] = string(output[row][col])
				}
			}
		}
	}

	return &grid{cells: newCells, size: newSize}
}

func part1(input string) {
	rules, err := NewRules(input)
	if err != nil {
		log.Fatalf("parsing rules failed: %v", err)
	}

	initial := newGrid([]string{
		".#.",
		"..#",
		"###",
	})

	grid := initial
	for range 5 {
		grid = rules.apply(grid)
	}
	fmt.Println("Part 1", grid.countOn())
}

func part2(input string) {
	rules, err := NewRules(input)
	if err != nil {
		log.Fatalf("parsing rules failed: %v", err)
	}

	initial := newGrid([]string{
		".#.",
		"..#",
		"###",
	})

	grid := initial
	for range 18 {
		grid = rules.apply(grid)
	}
	fmt.Println("Part 2", grid.countOn())
}
