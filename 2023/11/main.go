package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2023, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type point struct {
	x int
	y int
}

func part1and2(input string) {

	lines := conv.SplitNewline(input)
	galaxies := make(map[point]struct{})
	maxHeight := len(lines)
	maxWidth := len(lines[0])

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies[point{x, y}] = struct{}{}
			}
		}
	}

	emptyRows := container.NewSet[int]()
	emptyCols := container.NewSet[int]()

	for y := 0; y < maxHeight; y++ {
		rowEmpty := true
		for x := 0; x < maxWidth; x++ {
			if _, ok := galaxies[point{x, y}]; ok {
				rowEmpty = false
				break
			}
		}
		if rowEmpty {
			emptyRows.Add(y)
		}
	}

	for x := 0; x < maxWidth; x++ {
		colEmpty := true
		for y := 0; y < maxHeight; y++ {
			if _, ok := galaxies[point{x, y}]; ok {
				colEmpty = false
				break
			}
		}
		if colEmpty {
			emptyCols.Add(x)
		}
	}

	var galaxyPoints []point
	for p := range galaxies {
		galaxyPoints = append(galaxyPoints, p)
	}

	sumPart1 := 0
	sumPart2 := 0
	for i := 0; i < len(galaxyPoints)-1; i++ {
		for j := i + 1; j < len(galaxyPoints); j++ {
			start := galaxyPoints[i]
			end := galaxyPoints[j]

			ax, bx := min(start.x, end.x), max(start.x, end.x)
			ay, by := min(start.y, end.y), max(start.y, end.y)

			manhattanDistance := (bx - ax) + (by - ay)
			sumPart1 += manhattanDistance
			sumPart2 += manhattanDistance
			for x := ax; x <= bx; x++ {
				if emptyCols.Contains(x) {
					sumPart1 += 1
					sumPart2 += 1000000 - 1
				}
			}
			for y := ay; y <= by; y++ {
				if emptyRows.Contains(y) {
					sumPart1 += 1
					sumPart2 += 1000000 - 1
				}
			}
		}
	}
	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}
