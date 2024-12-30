package main

import (
	"aoc/2019/intcomputer"
	"aoc/internal/container"
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

func exploreMap(program []int) (gridutil.Grid2D[int], gridutil.Coordinate) {
	computer := intcomputer.NewIntcodeComputer(program)
	area := gridutil.NewGrid2D[int](false)
	var oxygenPos gridutil.Coordinate

	visited := make(map[gridutil.Coordinate]bool)
	var explore func(pos gridutil.Coordinate)
	explore = func(pos gridutil.Coordinate) {
		visited[pos] = true

		dirMap := map[int]gridutil.Direction{
			North: {Row: -1, Col: 0},
			South: {Row: 1, Col: 0},
			West:  {Row: 0, Col: -1},
			East:  {Row: 0, Col: 1},
		}

		for dir, offset := range dirMap {
			nextPos := gridutil.Coordinate{Row: pos.Row + offset.Row, Col: pos.Col + offset.Col}

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

			area.SetC(nextPos, result.Value)

			if result.Value == Oxygen {
				oxygenPos = nextPos
			}

			if result.Value != Wall {
				explore(nextPos)

				backDir := map[int]int{
					North: South,
					South: North,
					West:  East,
					East:  West,
				}[dir]

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
	area.SetC(start, Empty)
	explore(start)

	return area, oxygenPos
}

func findShortestPath(area gridutil.Grid2D[int], start, target gridutil.Coordinate) int {
	visited := make(map[gridutil.Coordinate]bool)
	queue := container.NewQueue[state]()
	queue.Push(state{pos: start, steps: 0})

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.pos == target {
			return current.steps
		}

		if visited[current.pos] {
			continue
		}
		visited[current.pos] = true

		for _, dir := range gridutil.Get4Directions() {
			next := gridutil.Coordinate{
				Row: current.pos.Row + dir.Row,
				Col: current.pos.Col + dir.Col,
			}

			if val, exists := area.GetC(next); exists && val != Wall && !visited[next] {
				queue.Push(state{pos: next, steps: current.steps + 1})
			}
		}
	}

	return -1
}

func fillWithOxygen(area gridutil.Grid2D[int], oxygenPos gridutil.Coordinate) int {
	minutes := 0
	oxygenated := make(map[gridutil.Coordinate]bool)
	queue := container.NewQueue[state]()
	queue.Push(state{pos: oxygenPos, steps: 0, oxygen: true})

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.steps > minutes {
			minutes = current.steps
		}

		if oxygenated[current.pos] {
			continue
		}
		oxygenated[current.pos] = true

		for _, dir := range gridutil.Get4Directions() {
			next := gridutil.Coordinate{
				Row: current.pos.Row + dir.Row,
				Col: current.pos.Col + dir.Col,
			}

			if val, exists := area.GetC(next); exists && val != Wall && !oxygenated[next] {
				queue.Push(state{pos: next, steps: current.steps + 1, oxygen: true})
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
