package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"slices"
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
	grid := gridutil.NewGrid2D[bool](false)
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, c == '#')
		}
	}

	maxVisible := 0
	for row := range grid.Height() {
		for col := range grid.Width() {
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			visible := detectableAsteroids(grid, gridutil.Coordinate{Row: row, Col: col})
			if visible > maxVisible {
				maxVisible = visible
			}
		}
	}

	fmt.Println("Part 1", maxVisible)
}

func slope(a, b gridutil.Coordinate) gridutil.Coordinate {
	dCol := b.Col - a.Col
	dRow := b.Row - a.Row
	g := mathx.Abs(mathx.Gcd(dCol, dRow))
	return gridutil.Coordinate{
		Row: dRow / g,
		Col: dCol / g,
	}
}

func detectableAsteroids(grid gridutil.Grid2D[bool], from gridutil.Coordinate) int {
	seenSlopes := container.NewSet[gridutil.Coordinate]()
	for row := range grid.Height() {
		for col := range grid.Width() {
			to := gridutil.Coordinate{Row: row, Col: col}
			if from == to {
				continue
			}
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			slope := slope(from, to)
			seenSlopes.Add(slope)
		}
	}
	return seenSlopes.Len()
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[bool](false)
	for row, line := range lines {
		for col, c := range line {
			grid.Set(row, col, c == '#')
		}
	}

	maxVisible := 0
	maxAsteroid := gridutil.Coordinate{}
	for row := range grid.Height() {
		for col := range grid.Width() {
			isAsteroid, _ := grid.Get(row, col)
			if !isAsteroid {
				continue
			}
			visible := detectableAsteroids(grid, gridutil.Coordinate{Row: row, Col: col})
			if visible > maxVisible {
				maxVisible = visible
				maxAsteroid = gridutil.Coordinate{Row: row, Col: col}
			}
		}
	}

	grid.Set(maxAsteroid.Row, maxAsteroid.Col, false)

	byAngle := map[float64][]gridutil.Coordinate{}
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
			byAngle[angle] = append(byAngle[angle], gridutil.Coordinate{Row: row, Col: col})
		}
	}

	for _, points := range byAngle {
		slices.SortFunc(points, func(a, b gridutil.Coordinate) int {
			return mathx.Abs(a.Col-maxAsteroid.Col) + mathx.Abs(a.Row-maxAsteroid.Row) - mathx.Abs(b.Col-maxAsteroid.Col) + mathx.Abs(b.Row-maxAsteroid.Row)
		})
	}

	angles := make([]float64, 0, len(byAngle))
	for angle := range byAngle {
		angles = append(angles, angle)
	}
	slices.Sort(angles)

	count := 0
	for len(byAngle) > 0 {
		for _, angle := range angles {
			points := byAngle[angle]
			if len(points) == 0 {
				continue
			}
			count++
			if count == 200 {
				fmt.Println("Part 2", points[0].Col*100+points[0].Row)
			}
			byAngle[angle] = points[1:]
			if len(byAngle[angle]) == 0 {
				delete(byAngle, angle)
			}
		}
	}
}
