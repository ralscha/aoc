package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/grid"
	"fmt"
	"log"
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

type path []grid.Coordinate

func part1(input string) {
	lines := conv.SplitNewline(input)

	paths := createPaths(lines)
	g := grid.NewGrid2D[bool](false)
	g.SetMinRow(0)
	drawPaths(&g, paths)

	abyss := false
	sands := 0
	for !abyss {
		var row int
		var col int
		abyss, row, col = fall(g, 0, 500)
		if !abyss {
			g.Set(row, col, true)
			sands++
		}
	}
	fmt.Println(sands)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	paths := createPaths(lines)
	g := grid.NewGrid2D[bool](false)
	g.SetMinRow(0)
	drawPaths(&g, paths)

	_, maxRow := g.GetMinMaxRow()
	minCol, maxCol := g.GetMinMaxCol()

	for x := minCol - 400; x <= maxCol+400; x++ {
		g.Set(maxRow+2, x, true)
	}

	sands := 0
	for {
		if v, _ := g.Get(0, 500); v {
			break
		}
		_, row, col := fall(g, 0, 500)
		g.Set(row, col, true)
		sands++
	}
	fmt.Println(sands)

}

func fall(g grid.Grid2D[bool], row, col int) (bool, int, int) {
	_, maxRow := g.GetMinMaxRow()
	if row+1 > maxRow {
		return true, row, col
	}
	_, ok := g.Get(row+1, col)
	if ok {
		minCol, maxCol := g.GetMinMaxCol()
		if col-1 < minCol {
			return true, row, col
		}
		_, ok := g.Get(row+1, col-1)
		if ok {
			if col+1 > maxCol {
				return true, row, col
			}
			_, ok := g.Get(row+1, col+1)
			if ok {
				return false, row, col
			} else {
				return fall(g, row+1, col+1)
			}
		} else {
			return fall(g, row+1, col-1)
		}
	} else {
		return fall(g, row+1, col)
	}
}

func createPaths(lines []string) []path {
	var paths []path

	for _, line := range lines {
		splitted := strings.Split(line, " -> ")
		var points []grid.Coordinate
		for _, s := range splitted {
			sp := strings.Split(s, ",")
			col := conv.MustAtoi(sp[0])
			row := conv.MustAtoi(sp[1])
			p := grid.Coordinate{Row: row, Col: col}
			points = append(points, p)
		}
		paths = append(paths, points)
	}
	return paths
}

func drawPaths(g *grid.Grid2D[bool], paths []path) {
	for _, p := range paths {
		prev := p[0]
		g.Set(prev.Row, prev.Col, true)
		for _, curr := range p[1:] {
			if prev.Col == curr.Col {
				if prev.Row < curr.Row {
					for r := prev.Row; r <= curr.Row; r++ {
						g.Set(r, curr.Col, true)
					}
				} else {
					for r := prev.Row; r >= curr.Row; r-- {
						g.Set(r, curr.Col, true)
					}
				}
			} else {
				if prev.Col < curr.Col {
					for c := prev.Col; c <= curr.Col; c++ {
						g.Set(curr.Row, c, true)
					}
				} else {
					for c := prev.Col; c >= curr.Col; c-- {
						g.Set(curr.Row, c, true)
					}
				}
			}
			prev = curr
		}
	}
}
