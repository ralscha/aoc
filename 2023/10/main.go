package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

type point struct {
	x, y int
}

func makeGrid(input string) map[point]byte {
	lines := conv.SplitNewline(input)
	grid := make(map[point]byte)
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				grid[point{x, y}] = byte(char)
			}
		}
	}
	return grid
}

func findStart(grid map[point]byte) point {
	for point, pipe := range grid {
		if pipe == 'S' {
			return point
		}
	}
	panic("no start found")
}

func main() {
	input, err := download.ReadInput(2023, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	grid := makeGrid(input)
	start := findStart(grid)

	path := make(map[point]struct{})
	path[start] = struct{}{}

	var direction point
	var currentPosition point

	for _, dir := range []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		searchPos := point{start.x + dir.x, start.y + dir.y}
		tile := grid[searchPos]
		if dir.x == 0 && dir.y == 1 {
			if tile == '|' || tile == 'L' || tile == 'J' || tile == 'S' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else if dir.x == 0 && dir.y == -1 {
			if tile == '|' || tile == 'F' || tile == '7' || tile == 'S' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else if dir.x == 1 && dir.y == 0 {
			if tile == '-' || tile == 'L' || tile == 'F' || tile == 'S' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else if dir.x == -1 && dir.y == 0 {
			if tile == '-' || tile == 'J' || tile == '7' || tile == 'S' {
				direction = dir
				currentPosition = searchPos
				break
			}
		}
	}

	steps := 0
	for grid[currentPosition] != 'S' {
		steps++
		tile := grid[currentPosition]
		if tile == '|' {
			path[currentPosition] = struct{}{}
			if direction.x == 0 && direction.y == 1 {
				currentPosition.y++
			} else if direction.x == 0 && direction.y == -1 {
				currentPosition.y--
			}
		} else if tile == '-' {
			path[currentPosition] = struct{}{}
			if direction.x == 1 && direction.y == 0 {
				currentPosition.x++
			} else if direction.x == -1 && direction.y == 0 {
				currentPosition.x--
			}
		} else if tile == '7' {
			path[currentPosition] = struct{}{}
			if direction.x == 1 && direction.y == 0 {
				currentPosition.y++
				direction = point{0, 1}
			} else if direction.x == 0 && direction.y == -1 {
				currentPosition.x--
				direction = point{-1, 0}
			}
		} else if tile == 'J' {
			path[currentPosition] = struct{}{}
			if direction.x == 1 && direction.y == 0 {
				currentPosition.y--
				direction = point{0, -1}
			} else if direction.x == 0 && direction.y == 1 {
				currentPosition.x--
				direction = point{-1, 0}
			}
		} else if tile == 'L' {
			path[currentPosition] = struct{}{}
			if direction.x == -1 && direction.y == 0 {
				currentPosition.y--
				direction = point{0, -1}
			} else if direction.x == 0 && direction.y == 1 {
				currentPosition.x++
				direction = point{1, 0}
			}
		} else if tile == 'F' {
			path[currentPosition] = struct{}{}
			if direction.x == -1 && direction.y == 0 {
				currentPosition.y++
				direction = point{0, 1}
			} else if direction.x == 0 && direction.y == -1 {
				currentPosition.x++
				direction = point{1, 0}
			}
		}
	}

	fmt.Println(steps/2 + 1)

	inside := 0
	lines := conv.SplitNewline(input)
	for y, line := range lines {
		northFacing := 0
		for x := range line {
			p := point{x, y}
			if _, ok := path[p]; ok {
				tile := line[x]
				if tile == '|' || tile == 'L' || tile == 'J' || tile == 'S' {
					northFacing++
				}
				continue
			}
			if northFacing%2 != 0 {
				inside++
			}
		}

	}

	fmt.Println(inside)

}
