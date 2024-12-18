package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 18)
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[rune](false)

	for y := 0; y <= 70; y++ {
		for x := 0; x <= 70; x++ {
			grid.Set(y, x, '.')
		}
	}

	for i := 0; i < 1024 && i < len(lines); i++ {
		parts := strings.Split(lines[i], ",")
		if len(parts) != 2 {
			continue
		}
		x := conv.MustAtoi(parts[0])
		y := conv.MustAtoi(parts[1])
		grid.Set(y, x, '#')
	}

	result, found := grid.ShortestPathWithBFS(
		gridutil.Coordinate{Row: 0, Col: 0},
		func(pos gridutil.Coordinate, val rune) bool {
			return pos.Row == 70 && pos.Col == 70
		},
		func(pos gridutil.Coordinate, val rune) bool {
			return val == '#'
		},
		gridutil.Get4Directions(),
	)

	if !found {
		fmt.Println("Part 1", "No path")
		return
	}

	fmt.Println("Part 1", len(result.Path)-1)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[rune](false)

	for y := 0; y <= 70; y++ {
		for x := 0; x <= 70; x++ {
			grid.Set(y, x, '.')
		}
	}

	for i, line := range lines {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		x := conv.MustAtoi(parts[0])
		y := conv.MustAtoi(parts[1])
		grid.Set(y, x, '#')

		if i < 1024 {
			continue
		}

		_, found := grid.ShortestPathWithBFS(
			gridutil.Coordinate{Row: 0, Col: 0},
			func(pos gridutil.Coordinate, val rune) bool {
				return pos.Row == 70 && pos.Col == 70
			},
			func(pos gridutil.Coordinate, val rune) bool {
				return val == '#'
			},
			gridutil.Get4Directions(),
		)

		if !found {
			s := fmt.Sprintf("%d,%d", x, y)
			fmt.Println("Part 2", s)
			return
		}
	}

	fmt.Println("Part 2", "Nothing found")
}
