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

type orientation struct {
	x, y int
}

var north = orientation{
	x: 0,
	y: -1,
}
var south = orientation{
	x: 0,
	y: 1,
}
var east = orientation{
	x: 1,
	y: 0,
}
var west = orientation{
	x: -1,
	y: 0,
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

	var direction orientation
	var currentPosition point

	for _, dir := range []orientation{north, south, east, west} {
		searchPos := point{start.x + dir.x, start.y + dir.y}
		tile := grid[searchPos]
		if dir == south {
			if tile == '|' || tile == 'L' || tile == 'J' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else if dir == north {
			if tile == '|' || tile == 'F' || tile == '7' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else if dir == west {
			if tile == '-' || tile == 'L' || tile == 'F' {
				direction = dir
				currentPosition = searchPos
				break
			}
		} else {
			if tile == '-' || tile == 'J' || tile == '7' {
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
			if direction == south {
				currentPosition.y++
			} else if direction == north {
				currentPosition.y--
			}
		} else if tile == '-' {
			path[currentPosition] = struct{}{}
			if direction == east {
				currentPosition.x++
			} else if direction == west {
				currentPosition.x--
			}
		} else if tile == '7' {
			path[currentPosition] = struct{}{}
			if direction == east {
				currentPosition.y++
				direction = south
			} else if direction == north {
				currentPosition.x--
				direction = west
			}
		} else if tile == 'J' {
			path[currentPosition] = struct{}{}
			if direction == east {
				currentPosition.y--
				direction = north
			} else if direction == south {
				currentPosition.x--
				direction = west
			}
		} else if tile == 'L' {
			path[currentPosition] = struct{}{}
			if direction == west {
				currentPosition.y--
				direction = north
			} else if direction == south {
				currentPosition.x++
				direction = east
			}
		} else if tile == 'F' {
			path[currentPosition] = struct{}{}
			if direction == west {
				currentPosition.y++
				direction = south
			} else if direction == north {
				currentPosition.x++
				direction = east
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
