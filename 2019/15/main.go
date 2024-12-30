package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

const (
	North = 1
	South = 2
	West  = 3
	East  = 4

	Wall   = 0
	Empty  = 1
	Oxygen = 2
)

type state struct {
	pos    gridutil.Coordinate
	steps  int
	oxygen bool
}

func main() {
	input, err := download.ReadInput(2019, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func exploreMap(program []int) (map[gridutil.Coordinate]int, gridutil.Coordinate) {
	computer := intcomputer.NewIntcodeComputer(program)
	area := make(map[gridutil.Coordinate]int)
	var oxygenPos gridutil.Coordinate

	visited := make(map[gridutil.Coordinate]bool)
	var explore func(pos gridutil.Coordinate)
	explore = func(pos gridutil.Coordinate) {
		visited[pos] = true

		for _, dir := range []int{North, South, West, East} {
			nextPos := pos
			switch dir {
			case North:
				nextPos.Row--
			case South:
				nextPos.Row++
			case West:
				nextPos.Col--
			case East:
				nextPos.Col++
			}

			if visited[nextPos] {
				continue
			}

			if err := computer.AddInput(dir); err != nil {
				log.Fatalf("adding input failed: %v", err)
			}

			result, err := computer.Run()
			if err != nil {
				log.Fatalf("running program failed: %v", err)
			}

			if result.Signal != intcomputer.SignalOutput {
				log.Fatal("expected output signal")
			}

			area[nextPos] = result.Value

			if result.Value == Oxygen {
				oxygenPos = nextPos
			}

			if result.Value != Wall {
				explore(nextPos)
				// Move back
				var backDir int
				switch dir {
				case North:
					backDir = South
				case South:
					backDir = North
				case West:
					backDir = East
				case East:
					backDir = West
				}
				if err := computer.AddInput(backDir); err != nil {
					log.Fatalf("adding input failed: %v", err)
				}
				if _, err := computer.Run(); err != nil {
					log.Fatalf("running program failed: %v", err)
				}
			}
		}
	}

	start := gridutil.Coordinate{Row: 0, Col: 0}
	area[start] = Empty
	explore(start)

	return area, oxygenPos
}

func findShortestPath(area map[gridutil.Coordinate]int, start, target gridutil.Coordinate) int {
	visited := make(map[gridutil.Coordinate]bool)
	queue := []state{{pos: start, steps: 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == target {
			return current.steps
		}

		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true

		// Try all directions
		for _, offset := range []gridutil.Coordinate{
			{Row: -1, Col: 0}, // North
			{Row: 1, Col: 0},  // South
			{Row: 0, Col: -1}, // West
			{Row: 0, Col: 1},  // East
		} {
			next := gridutil.Coordinate{
				Row: current.pos.Row + offset.Row,
				Col: current.pos.Col + offset.Col,
			}

			if tile, exists := area[next]; exists && tile != Wall && !visited[next] {
				queue = append(queue, state{pos: next, steps: current.steps + 1})
			}
		}
	}

	return -1
}

func fillWithOxygen(area map[gridutil.Coordinate]int, oxygenPos gridutil.Coordinate) int {
	minutes := 0
	oxygenated := make(map[gridutil.Coordinate]bool)
	queue := []state{{pos: oxygenPos, steps: 0, oxygen: true}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.steps > minutes {
			minutes = current.steps
		}

		if oxygenated[current.pos] {
			continue
		}
		oxygenated[current.pos] = true

		// Spread oxygen to adjacent spaces
		for _, offset := range []gridutil.Coordinate{
			{Row: -1, Col: 0}, // North
			{Row: 1, Col: 0},  // South
			{Row: 0, Col: -1}, // West
			{Row: 0, Col: 1},  // East
		} {
			next := gridutil.Coordinate{
				Row: current.pos.Row + offset.Row,
				Col: current.pos.Col + offset.Col,
			}

			if tile, exists := area[next]; exists && tile != Wall && !oxygenated[next] {
				queue = append(queue, state{pos: next, steps: current.steps + 1, oxygen: true})
			}
		}
	}

	return minutes
}

func part1(input string) {
	program := conv.ToIntSliceComma(input)
	area, oxygenPos := exploreMap(program)
	start := gridutil.Coordinate{Row: 0, Col: 0}
	steps := findShortestPath(area, start, oxygenPos)
	fmt.Println("Part 1", steps)
}

func part2(input string) {
	program := conv.ToIntSliceComma(input)
	area, oxygenPos := exploreMap(program)
	minutes := fillWithOxygen(area, oxygenPos)
	fmt.Println("Part 2", minutes)
}
