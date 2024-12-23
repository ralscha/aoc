package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2017, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type nodeState int

const (
	Clean nodeState = iota
	Weakened
	Infected
	Flagged
)

type virus struct {
	pos       gridutil.Coordinate
	direction gridutil.Direction
}

type grid struct {
	nodes map[gridutil.Coordinate]nodeState
}

func newGrid(input string) *grid {
	lines := conv.SplitNewline(input)
	grid := &grid{
		nodes: make(map[gridutil.Coordinate]nodeState),
	}

	for r, line := range lines {
		for c, char := range line {
			if char == '#' {
				pos := gridutil.Coordinate{Row: r, Col: c}
				grid.nodes[pos] = Infected
			}
		}
	}

	return grid
}

func (g *grid) getState(pos gridutil.Coordinate) nodeState {
	return g.nodes[pos]
}

func (g *grid) setState(pos gridutil.Coordinate, state nodeState) {
	if state == Clean {
		delete(g.nodes, pos)
	} else {
		g.nodes[pos] = state
	}
}

func (v *virus) turnLeft() {
	switch v.direction {
	case gridutil.DirectionN:
		v.direction = gridutil.DirectionW
	case gridutil.DirectionW:
		v.direction = gridutil.DirectionS
	case gridutil.DirectionS:
		v.direction = gridutil.DirectionE
	case gridutil.DirectionE:
		v.direction = gridutil.DirectionN
	}
}

func (v *virus) turnRight() {
	switch v.direction {
	case gridutil.DirectionN:
		v.direction = gridutil.DirectionE
	case gridutil.DirectionE:
		v.direction = gridutil.DirectionS
	case gridutil.DirectionS:
		v.direction = gridutil.DirectionW
	case gridutil.DirectionW:
		v.direction = gridutil.DirectionN
	}
}

func (v *virus) reverse() {
	switch v.direction {
	case gridutil.DirectionN:
		v.direction = gridutil.DirectionS
	case gridutil.DirectionS:
		v.direction = gridutil.DirectionN
	case gridutil.DirectionE:
		v.direction = gridutil.DirectionW
	case gridutil.DirectionW:
		v.direction = gridutil.DirectionE
	}
}

func (v *virus) move() {
	v.pos = gridutil.Coordinate{
		Row: v.pos.Row + v.direction.Row,
		Col: v.pos.Col + v.direction.Col,
	}
}

func simulate(input string, bursts int, stateTransition func(grid *grid, virus *virus) int) int {
	grid := newGrid(input)
	lines := conv.SplitNewline(input)
	startRow := len(lines) / 2
	startCol := len(lines[0]) / 2

	virus := &virus{
		pos:       gridutil.Coordinate{Row: startRow, Col: startCol},
		direction: gridutil.DirectionN,
	}

	infections := 0
	for range bursts {
		infections += stateTransition(grid, virus)
		virus.move()
	}

	return infections
}

func part1(input string) {
	infections := simulate(input, 10000, func(grid *grid, virus *virus) int {
		state := grid.getState(virus.pos)
		if state == Infected {
			virus.turnRight()
			grid.setState(virus.pos, Clean)
			return 0
		} else {
			virus.turnLeft()
			grid.setState(virus.pos, Infected)
			return 1
		}
	})
	fmt.Println("Part 1", infections)
}

func part2(input string) {
	infections := simulate(input, 10000000, func(grid *grid, virus *virus) int {
		state := grid.getState(virus.pos)
		switch state {
		case Clean:
			virus.turnLeft()
			grid.setState(virus.pos, Weakened)
			return 0
		case Weakened:
			grid.setState(virus.pos, Infected)
			return 1
		case Infected:
			virus.turnRight()
			grid.setState(virus.pos, Flagged)
			return 0
		case Flagged:
			virus.reverse()
			grid.setState(virus.pos, Clean)
			return 0
		}
		return 0
	})
	fmt.Println("Part 2", infections)
}
