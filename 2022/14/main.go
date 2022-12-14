package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2022/14/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type point struct {
	x, y int
}

type path struct {
	points []point
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	paths, max, min, maxHeight := createPaths(lines)
	grid := createGrid(max, min, maxHeight)
	drawGrid(paths, grid, min)

	abyss := false
	sands := 0
	for !abyss {
		var p point
		abyss, p = fall(grid, point{500 - min, 0})
		if !abyss {
			grid[p.x][p.y] = true
			sands++
		}
	}
	fmt.Println(sands)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	paths, max, min, maxHeight := createPaths(lines)

	maxHeight += 3
	min = min - 400
	max = max + 400

	grid := createGrid(max, min, maxHeight)
	for x := min; x < max; x++ {
		grid[x-min][maxHeight-1] = true
	}

	drawGrid(paths, grid, min)

	sands := 0
	for {
		if grid[500-min][0] {
			break
		}
		_, p := fall(grid, point{500 - min, 0})
		grid[p.x][p.y] = true
		sands++
	}
	fmt.Println(sands)

}

func fall(grid [][]bool, p point) (bool, point) {
	if p.y+1 >= len(grid[0]) {
		return true, p
	}

	if grid[p.x][p.y+1] {
		if p.x-1 < 0 {
			return true, p
		}
		if grid[p.x-1][p.y+1] {
			if p.x+1 >= len(grid) {
				return true, p
			}
			if grid[p.x+1][p.y+1] {
				return false, p
			} else {
				return fall(grid, point{p.x + 1, p.y + 1})
			}
		} else {
			return fall(grid, point{p.x - 1, p.y + 1})
		}
	} else {
		return fall(grid, point{p.x, p.y + 1})
	}
}

func createGrid(max int, min int, maxHeight int) [][]bool {
	grid := make([][]bool, max-min+1)
	for r := range grid {
		grid[r] = make([]bool, maxHeight+1)
		for c := range grid[r] {
			grid[r][c] = false
		}
	}
	return grid
}

func createPaths(lines []string) ([]path, int, int, int) {
	max := 0
	min := math.MaxInt32
	maxHeight := 0

	var paths []path

	for _, line := range lines {
		splitted := strings.Split(line, " -> ")
		var points []point
		for _, s := range splitted {
			sp := strings.Split(s, ",")
			x := conv.MustAtoi(sp[0])
			y := conv.MustAtoi(sp[1])
			if x > max {
				max = x
			}
			if x < min {
				min = x
			}
			if y > maxHeight {
				maxHeight = y
			}
			p := point{x, y}
			points = append(points, p)
		}
		paths = append(paths, path{points})
	}
	return paths, max, min, maxHeight
}

func drawGrid(paths []path, grid [][]bool, min int) {
	for _, p := range paths {
		prev := p.points[0]
		grid[prev.x-min][prev.y] = true
		for _, curr := range p.points[1:] {
			if prev.x == curr.x {
				if prev.y < curr.y {
					for i := prev.y; i <= curr.y; i++ {
						grid[curr.x-min][i] = true
					}
				} else {
					for i := prev.y; i >= curr.y; i-- {
						grid[curr.x-min][i] = true
					}
				}
			} else {
				if prev.x < curr.x {
					for i := prev.x; i <= curr.x; i++ {
						grid[i-min][curr.y] = true
					}
				} else {
					for i := prev.x; i >= curr.x; i-- {
						grid[i-min][curr.y] = true
					}
				}
			}
			prev = curr
		}
	}
}
