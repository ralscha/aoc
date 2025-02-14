package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"regexp"
)

func main() {
	input, err := download.ReadInput(2016, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type node struct {
	size    int
	used    int
	avail   int
	usePerc int
}

func parseInput(input string) gridutil.Grid2D[node] {
	grid := gridutil.NewGrid2D[node](false)
	lines := conv.SplitNewline(input)
	lines = lines[2:]
	re := regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		x := conv.MustAtoi(matches[1])
		y := conv.MustAtoi(matches[2])
		n := node{
			size:    conv.MustAtoi(matches[3]),
			used:    conv.MustAtoi(matches[4]),
			avail:   conv.MustAtoi(matches[5]),
			usePerc: conv.MustAtoi(matches[6]),
		}

		grid.Set(y, x, n)
	}

	return grid
}

func countViablePairs(grid gridutil.Grid2D[node]) int {
	count := 0
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	for y1 := minRow; y1 <= maxRow; y1++ {
		for x1 := minCol; x1 <= maxCol; x1++ {
			nodeA, _ := grid.Get(y1, x1)
			if nodeA.used == 0 {
				continue
			}

			for y2 := minRow; y2 <= maxRow; y2++ {
				for x2 := minCol; x2 <= maxCol; x2++ {
					if y1 == y2 && x1 == x2 {
						continue
					}

					nodeB, _ := grid.Get(y2, x2)
					if nodeA.used <= nodeB.avail {
						count++
					}
				}
			}
		}
	}

	return count
}

type state struct {
	empty gridutil.Coordinate
	goal  gridutil.Coordinate
}

func findMinimumSteps(grid gridutil.Grid2D[node]) int {
	minRow, maxRow := grid.GetMinMaxRow()
	minCol, maxCol := grid.GetMinMaxCol()

	var initialEmpty gridutil.Coordinate
	for y := minRow; y <= maxRow; y++ {
		for x := minCol; x <= maxCol; x++ {
			n, _ := grid.Get(y, x)
			if n.used == 0 {
				initialEmpty = gridutil.Coordinate{Row: y, Col: x}
				break
			}
		}
	}

	validMoves := gridutil.NewGrid2D[bool](false)
	for y := minRow; y <= maxRow; y++ {
		for x := minCol; x <= maxCol; x++ {
			n, _ := grid.Get(y, x)
			validMoves.Set(y, x, n.used <= 400)
		}
	}

	startState := state{
		empty: initialEmpty,
		goal:  gridutil.Coordinate{Row: 0, Col: maxCol},
	}

	queue := container.NewQueue[struct {
		state state
		steps int
	}]()
	queue.Push(struct {
		state state
		steps int
	}{state: startState, steps: 0})

	visited := container.NewSet[state]()
	visited.Add(startState)

	directions := gridutil.Get4Directions()

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.state.goal.Row == 0 && current.state.goal.Col == 0 {
			return current.steps
		}

		for _, dir := range directions {
			newEmpty := gridutil.Coordinate{
				Row: current.state.empty.Row + dir.Row,
				Col: current.state.empty.Col + dir.Col,
			}

			if valid, _ := validMoves.GetC(newEmpty); valid {
				nextState := state{
					empty: newEmpty,
					goal:  current.state.goal,
				}

				if nextState.empty == nextState.goal {
					nextState.goal = current.state.empty
				}

				if !visited.Contains(nextState) {
					visited.Add(nextState)
					queue.Push(struct {
						state state
						steps int
					}{state: nextState, steps: current.steps + 1})
				}
			}
		}
	}

	return -1
}

func part1(input string) {
	grid := parseInput(input)
	result := countViablePairs(grid)
	fmt.Println("Part 1", result)
}

func part2(input string) {
	grid := parseInput(input)
	steps := findMinimumSteps(grid)
	fmt.Println("Part 2", steps)
}
