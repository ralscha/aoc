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
	input, err := download.ReadInput(2019, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	start := gridutil.Coordinate{}
	keys := make(map[rune]gridutil.Coordinate)
	doors := make(map[rune]gridutil.Coordinate)
	allKeys := 0

	for row := range grid.Height() {
		for col := range grid.Width() {
			coord := gridutil.Coordinate{Row: row, Col: col}
			value, _ := grid.Get(row, col)
			if value == '@' {
				start = coord
			} else if value >= 'a' && value <= 'z' {
				keys[value] = coord
				allKeys |= 1 << (value - 'a')
			} else if value >= 'A' && value <= 'Z' {
				doors[value] = coord
			}
		}
	}

	type state struct {
		coord gridutil.Coordinate
		keys  int
		steps int
	}

	queue := container.NewQueue[state]()
	startState := state{coord: start, keys: 0, steps: 0}
	queue.Push(startState)

	visited := container.NewSet[[3]int]()
	visited.Add([3]int{start.Row, start.Col, 0})

	for !queue.IsEmpty() {
		currentState := queue.Pop()

		if currentState.keys == allKeys {
			fmt.Println("Part 1", currentState.steps)
			return
		}

		for _, dir := range gridutil.Get4Directions() {
			newCoord := gridutil.Coordinate{Row: currentState.coord.Row + dir.Row, Col: currentState.coord.Col + dir.Col}

			if value, ok := grid.GetC(newCoord); ok && value != '#' {
				canPass := true
				if value >= 'A' && value <= 'Z' {
					requiredKey := value - 'A'
					if (currentState.keys & (1 << requiredKey)) == 0 {
						canPass = false
					}
				}

				if canPass {
					newKeys := currentState.keys
					if value >= 'a' && value <= 'z' {
						keyIndex := value - 'a'
						if (newKeys & (1 << keyIndex)) == 0 {
							newKeys |= 1 << keyIndex
						}
					}

					newVisitedKey := [3]int{newCoord.Row, newCoord.Col, newKeys}
					if !visited.Contains(newVisitedKey) {
						visited.Add(newVisitedKey)
						newState := state{coord: newCoord, keys: newKeys, steps: currentState.steps + 1}
						queue.Push(newState)
					}
				}
			}
		}
	}
}

func part2(input string) {
	fmt.Println("Part 2", "Not yet implemented")
}
