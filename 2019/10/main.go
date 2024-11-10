package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	grid2 "aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	input, err := download.ReadInput(2019, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := grid2.NewGrid2D[bool](false)
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, c == '#')
		}
	}

	maxVisible := 0
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			visible := detectableAsteroids(grid, grid2.Coordinate{Row: row, Col: col})
			if visible > maxVisible {
				maxVisible = visible
			}
		}
	}

	fmt.Println("Part 1:", maxVisible)
}

func slope(a, b grid2.Coordinate) grid2.Coordinate {
	dCol := b.Col - a.Col
	dRow := b.Row - a.Row
	g := mathx.Abs(mathx.Gcd(dCol, dRow))
	return grid2.Coordinate{
		Row: dRow / g,
		Col: dCol / g,
	}
}

func detectableAsteroids(grid grid2.Grid2D[bool], from grid2.Coordinate) int {
	seenSlopes := make(map[grid2.Coordinate]bool)
	for row := range grid.Height() {
		for col := range grid.Width() {
			to := grid2.Coordinate{Row: row, Col: col}
			if from == to {
				continue
			}
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			slope := slope(from, to)
			seenSlopes[slope] = true
		}
	}
	return len(seenSlopes)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := grid2.NewGrid2D[bool](false)
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, c == '#')
		}
	}

	maxVisible := 0
	maxAsteroid := grid2.Coordinate{}
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			visible := detectableAsteroids(grid, grid2.Coordinate{Row: row, Col: col})
			if visible > maxVisible {
				maxVisible = visible
				maxAsteroid = grid2.Coordinate{Row: row, Col: col}
			}
		}
	}

	grid.Set(maxAsteroid.Row, maxAsteroid.Col, false)

	byAngle := map[float64][]grid2.Coordinate{}
	for row := range grid.Height() {
		for col := range grid.Width() {
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			dx, dy := float64(col-maxAsteroid.Col), float64(row-maxAsteroid.Row)
			angle := math.Atan2(dx, -dy)
			if angle < 0 {
				angle += 2 * math.Pi
			}
			byAngle[angle] = append(byAngle[angle], grid2.Coordinate{Row: row, Col: col})
		}
	}

	for _, points := range byAngle {
		sort.Slice(points, func(i, j int) bool {
			di := mathx.Abs(points[i].Col-maxAsteroid.Col) + mathx.Abs(points[i].Row-maxAsteroid.Row)
			dj := mathx.Abs(points[j].Col-maxAsteroid.Col) + mathx.Abs(points[j].Row-maxAsteroid.Row)
			return di < dj
		})
	}

	angles := make([]float64, 0, len(byAngle))
	for angle := range byAngle {
		angles = append(angles, angle)
	}
	sort.Float64s(angles)

	count := 0
	for len(byAngle) > 0 {
		for _, angle := range angles {
			points := byAngle[angle]
			if len(points) == 0 {
				continue
			}
			count++
			if count == 200 {
				fmt.Println("Part 2:", points[0].Col*100+points[0].Row)
			}
			byAngle[angle] = points[1:]
			if len(byAngle[angle]) == 0 {
				delete(byAngle, angle)
			}
		}
	}
}
