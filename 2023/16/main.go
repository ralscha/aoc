package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2023, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type direction int

const (
	right direction = iota
	down
	left
	up
)

type point struct {
	x, y int
}

type beam struct {
	pos point
	dir direction
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	startPos := point{x: 0, y: 0}
	startDir := right

	energized := energize(startPos, startDir, grid)

	fmt.Println(energized)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var grid [][]byte
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}

	maxEnergized := 0

	for x := 0; x < len(grid[0]); x++ {
		energized := energize(point{x, 0}, down, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}
	for x := 0; x < len(grid[len(grid)-1]); x++ {
		energized := energize(point{x, len(grid) - 1}, up, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}
	for y := range len(grid) {
		energized := energize(point{0, y}, right, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}
	for y := range len(grid) {
		energized := energize(point{0, len(grid[y]) - 1}, left, grid)
		if energized > maxEnergized {
			maxEnergized = energized
		}
	}

	fmt.Println(maxEnergized)
}

func energize(startPos point, startDir direction, grid [][]byte) int {
	var beams []beam
	beams = append(beams, beam{pos: startPos, dir: startDir})
	energized := container.NewSet[point]()
	cycleDetection := container.NewSet[beam]()
	for len(beams) > 0 {
		currentBeam := beams[0]
		beams = beams[1:]

		for {
			if currentBeam.pos.y < 0 || currentBeam.pos.y >= len(grid) || currentBeam.pos.x < 0 || currentBeam.pos.x >= len(grid[currentBeam.pos.y]) {
				break
			}

			energized.Add(currentBeam.pos)
			tile := grid[currentBeam.pos.y][currentBeam.pos.x]

			switch tile {
			case '/':
				switch currentBeam.dir {
				case right:
					currentBeam.dir = up
					currentBeam.pos.y--
				case down:
					currentBeam.dir = left
					currentBeam.pos.x--
				case left:
					currentBeam.dir = down
					currentBeam.pos.y++
				case up:
					currentBeam.dir = right
					currentBeam.pos.x++
				}
			case '\\':
				switch currentBeam.dir {
				case right:
					currentBeam.dir = down
					currentBeam.pos.y++
				case down:
					currentBeam.dir = right
					currentBeam.pos.x++
				case left:
					currentBeam.dir = up
					currentBeam.pos.y--
				case up:
					currentBeam.dir = left
					currentBeam.pos.x--
				}
			case '|', '-':
				if (tile == '|' && (currentBeam.dir == left || currentBeam.dir == right)) || (tile == '-' && (currentBeam.dir == up || currentBeam.dir == down)) {
					if currentBeam.dir == left || currentBeam.dir == right {
						beams = append(beams, beam{pos: point{x: currentBeam.pos.x, y: currentBeam.pos.y - 1}, dir: up})
						beams = append(beams, beam{pos: point{x: currentBeam.pos.x, y: currentBeam.pos.y + 1}, dir: down})
					} else {
						beams = append(beams, beam{pos: point{x: currentBeam.pos.x - 1, y: currentBeam.pos.y}, dir: left})
						beams = append(beams, beam{pos: point{x: currentBeam.pos.x + 1, y: currentBeam.pos.y}, dir: right})
					}
					break
				}
				fallthrough
			default:
				switch currentBeam.dir {
				case right:
					currentBeam.pos.x++
				case down:
					currentBeam.pos.y++
				case left:
					currentBeam.pos.x--
				case up:
					currentBeam.pos.y--
				}
			}

			if cycleDetection.Contains(currentBeam) {
				break
			}
			cycleDetection.Add(currentBeam)
		}
	}
	return energized.Len()
}
