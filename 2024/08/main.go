package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2024, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	lines := conv.SplitNewline(input)
	antennas := make(map[rune][]gridutil.Coordinate)

	for row, line := range lines {
		for col, ch := range line {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], gridutil.Coordinate{Row: row, Col: col})
			}
		}
	}

	maxRows := len(lines)
	maxCols := len(lines[0])
	fmt.Println("Part 1", findAntinodes(antennas, maxRows, maxCols, false).Len())
	fmt.Println("Part 2", findAntinodes(antennas, maxRows, maxCols, true).Len())
}

func findAntinodes(antennas map[rune][]gridutil.Coordinate, rows, cols int, part2 bool) *container.Set[gridutil.Coordinate] {
	points := container.NewSet[gridutil.Coordinate]()

	for _, pairs := range antennas {
		for i, a := range pairs {
			for _, b := range pairs[i+1:] {
				if part2 {
					dr := b.Row - a.Row
					dc := b.Col - a.Col
					gcd := mathx.Gcd(mathx.Abs(dr), mathx.Abs(dc))
					if gcd == 0 {
						gcd = 1
					}

					stepR, stepC := dr/gcd, dc/gcd
					for dir := -1; dir <= 1; dir += 2 {
						r, c := a.Row, a.Col
						for r >= 0 && r < rows && c >= 0 && c < cols {
							points.Add(gridutil.Coordinate{Row: r, Col: c})
							r += stepR * dir
							c += stepC * dir
						}
					}
				} else {
					for _, p := range []gridutil.Coordinate{
						{Row: 2*b.Row - a.Row, Col: 2*b.Col - a.Col},
						{Row: 2*a.Row - b.Row, Col: 2*a.Col - b.Col},
					} {
						if p.Row >= 0 && p.Row < rows && p.Col >= 0 && p.Col < cols {
							points.Add(p)
						}
					}
				}
			}
		}
	}

	return points
}
