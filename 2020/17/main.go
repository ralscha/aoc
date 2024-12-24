package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type Coordinate4D struct {
	X, Y, Z, W int
}

func getNeighbors3D(c gridutil.Coordinate3D) []gridutil.Coordinate3D {
	neighbors := make([]gridutil.Coordinate3D, 0, 26)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				neighbors = append(neighbors, gridutil.Coordinate3D{X: c.X + dx, Y: c.Y + dy, Z: c.Z + dz})
			}
		}
	}
	return neighbors
}

func getNeighbors4D(c Coordinate4D) []Coordinate4D {
	neighbors := make([]Coordinate4D, 0, 80)
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					neighbors = append(neighbors, Coordinate4D{c.X + dx, c.Y + dy, c.Z + dz, c.W + dw})
				}
			}
		}
	}
	return neighbors
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	active := container.NewSet[gridutil.Coordinate3D]()
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				active.Add(gridutil.Coordinate3D{X: x, Y: y})
			}
		}
	}

	for range 6 {
		nextActive := container.NewSet[gridutil.Coordinate3D]()
		candidates := container.NewSet[gridutil.Coordinate3D]()
		for _, c := range active.Values() {
			candidates.Add(c)
			for _, n := range getNeighbors3D(c) {
				candidates.Add(n)
			}
		}

		for _, c := range candidates.Values() {
			count := 0
			for _, n := range getNeighbors3D(c) {
				if active.Contains(n) {
					count++
				}
			}
			if active.Contains(c) {
				if count == 2 || count == 3 {
					nextActive.Add(c)
				}
			} else {
				if count == 3 {
					nextActive.Add(c)
				}
			}
		}
		active = nextActive
	}

	fmt.Println("Part 1", active.Len())
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	active := container.NewSet[Coordinate4D]()
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				active.Add(Coordinate4D{x, y, 0, 0})
			}
		}
	}

	for range 6 {
		nextActive := container.NewSet[Coordinate4D]()
		candidates := container.NewSet[Coordinate4D]()
		for _, c := range active.Values() {
			candidates.Add(c)
			for _, n := range getNeighbors4D(c) {
				candidates.Add(n)
			}
		}

		for _, c := range candidates.Values() {
			count := 0
			for _, n := range getNeighbors4D(c) {
				if active.Contains(n) {
					count++
				}
			}
			if active.Contains(c) {
				if count == 2 || count == 3 {
					nextActive.Add(c)
				}
			} else {
				if count == 3 {
					nextActive.Add(c)
				}
			}
		}
		active = nextActive
	}

	fmt.Println("Part 2", active.Len())
}
