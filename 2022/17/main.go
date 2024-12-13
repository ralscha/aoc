package main

import (
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func addCoord(a, b gridutil.Coordinate) gridutil.Coordinate {
	return gridutil.Coordinate{
		Row: a.Row + b.Row,
		Col: a.Col + b.Col,
	}
}

func main() {
	input, err := download.ReadInput(2022, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(jets string) {
	rocks := [][]gridutil.Coordinate{
		{{Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}, {Row: 0, Col: 3}},                   // horizontal line
		{{Row: 2, Col: 1}, {Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 1, Col: 2}, {Row: 0, Col: 1}}, // plus
		{{Row: 2, Col: 2}, {Row: 1, Col: 2}, {Row: 0, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: 2}}, // reverse L
		{{Row: 3, Col: 0}, {Row: 2, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: 0}},                   // vertical line
		{{Row: 1, Col: 0}, {Row: 1, Col: 1}, {Row: 0, Col: 0}, {Row: 0, Col: 1}},                   // square
	}

	grid := map[gridutil.Coordinate]struct{}{}
	move := func(rock []gridutil.Coordinate, delta gridutil.Coordinate) bool {
		nrock := make([]gridutil.Coordinate, len(rock))
		for i, p := range rock {
			p = addCoord(p, delta)
			if _, ok := grid[p]; ok || p.Col < 0 || p.Col >= 7 || p.Row < 0 {
				return false
			}
			nrock[i] = p
		}
		copy(rock, nrock)
		return true
	}

	cache := map[[2]int][]int{}

	height, jet := 0, 0
	for i := 0; i < 1000000000000; i++ {
		if i == 2022 {
			fmt.Println(height)
		}

		k := [2]int{i % len(rocks), jet}
		if c, ok := cache[k]; ok {
			if n, d := 1000000000000-i, i-c[0]; n%d == 0 {
				fmt.Println(height + n/d*(height-c[1]))
				break
			}
		}
		cache[k] = []int{i, height}

		var rock []gridutil.Coordinate
		for _, p := range rocks[i%len(rocks)] {
			rock = append(rock, addCoord(p, gridutil.Coordinate{Row: height + 3, Col: 2}))
		}

		for {
			move(rock, gridutil.Coordinate{Row: 0, Col: int(jets[jet]) - int('=')})
			jet = (jet + 1) % len(jets)

			if !move(rock, gridutil.Coordinate{Row: -1, Col: 0}) {
				for _, p := range rock {
					grid[p] = struct{}{}
					if p.Row+1 > height {
						height = p.Row + 1
					}
				}
				break
			}
		}
	}
}
