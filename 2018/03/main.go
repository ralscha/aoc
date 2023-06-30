package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/grid"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2018/03/input.txt"
	input, err := download.ReadInput(inputFile, 2018, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	claims := make([]claim, len(lines))
	for i, line := range lines {
		var c claim
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.width, &c.height)
		if err != nil {
			log.Fatalf("scanning claim failed: %v", err)
		}
		claims[i] = c
	}

	g := grid.NewGrid2D[int](false)
	g.SetMaxRowCol(1000, 1000)
	for _, c := range claims {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if count, ok := g.Get(i, j); ok {
					g.Set(i, j, count+1)
				} else {
					g.Set(i, j, 1)
				}
			}
		}
	}

	overlaps := 0
	minCol, maxCol := g.GetMinMaxCol()
	minRow, maxRow := g.GetMinMaxRow()
	for i := minCol; i <= maxCol; i++ {
		for j := minRow; j <= maxRow; j++ {
			if count, ok := g.Get(i, j); ok {
				if count > 1 {
					overlaps++
				}
			}
		}
	}

	fmt.Println(overlaps)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	claims := make([]claim, len(lines))
	for i, line := range lines {
		var c claim
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.width, &c.height)
		if err != nil {
			log.Fatalf("scanning claim failed: %v", err)
		}
		claims[i] = c
	}

	g := grid.NewGrid2D[int](false)
	g.SetMaxRowCol(1000, 1000)
	for _, c := range claims {
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if count, ok := g.Get(i, j); ok {
					g.Set(i, j, count+1)
				} else {
					g.Set(i, j, 1)
				}
			}
		}
	}

	for _, c := range claims {
		overlaps := false
		for i := c.x; i < c.x+c.width; i++ {
			for j := c.y; j < c.y+c.height; j++ {
				if count, ok := g.Get(i, j); ok {
					if count > 1 {
						overlaps = true
						break
					}
				}
			}
			if overlaps {
				break
			}
		}
		if !overlaps {
			fmt.Println(c.id)
			return
		}
	}
}
