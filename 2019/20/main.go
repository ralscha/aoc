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
	input, err := download.ReadInput(2019, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	maze := parseMaze(input)
	steps := maze.findShortestPathPart1()
	fmt.Println("Part 1", steps)
}

func part2(input string) {
	maze := parseMaze(input)
	steps := maze.findShortestPathPart2()
	fmt.Println("Part 2", steps)
}

type state struct {
	pos   gridutil.Coordinate
	level int
}

type portal struct {
	label    string
	entrance gridutil.Coordinate
	isOuter  bool
}

type maze struct {
	grid    gridutil.Grid2D[rune]
	portals map[string][]portal
	width   int
	height  int
}

func (m *maze) isOuterPortal(p gridutil.Coordinate) bool {
	return p.Col <= 2 || p.Row <= 2 || p.Col >= m.width-3 || p.Row >= m.height-3
}

func parseMaze(input string) *maze {
	lines := conv.SplitNewline(input)
	height := len(lines)
	width := len(lines[0])
	grid := gridutil.NewCharGrid2D(lines)

	maze := &maze{
		grid:    grid,
		portals: make(map[string][]portal),
		width:   width,
		height:  height,
	}

	isLetter := func(r rune) bool {
		return r >= 'A' && r <= 'Z'
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width-1; x++ {
			if val1, exists1 := grid.Get(y, x); exists1 {
				if val2, exists2 := grid.Get(y, x+1); exists2 {
					if isLetter(val1) && isLetter(val2) {
						label := string(val1) + string(val2)
						if x > 0 {
							if val, exists := grid.Get(y, x-1); exists && val == '.' {
								entrance := gridutil.Coordinate{Row: y, Col: x - 1}
								maze.portals[label] = append(maze.portals[label], portal{label, entrance, maze.isOuterPortal(entrance)})
							}
						}
						if x+2 < width {
							if val, exists := grid.Get(y, x+2); exists && val == '.' {
								entrance := gridutil.Coordinate{Row: y, Col: x + 2}
								maze.portals[label] = append(maze.portals[label], portal{label, entrance, maze.isOuterPortal(entrance)})
							}
						}
					}
				}
			}
		}
	}

	for y := 0; y < height-1; y++ {
		for x := 0; x < width; x++ {
			if val1, exists1 := grid.Get(y, x); exists1 {
				if val2, exists2 := grid.Get(y+1, x); exists2 {
					if isLetter(val1) && isLetter(val2) {
						label := string(val1) + string(val2)
						if y > 0 {
							if val, exists := grid.Get(y-1, x); exists && val == '.' {
								entrance := gridutil.Coordinate{Row: y - 1, Col: x}
								maze.portals[label] = append(maze.portals[label], portal{label, entrance, maze.isOuterPortal(entrance)})
							}
						}
						if y+2 < height {
							if val, exists := grid.Get(y+2, x); exists && val == '.' {
								entrance := gridutil.Coordinate{Row: y + 2, Col: x}
								maze.portals[label] = append(maze.portals[label], portal{label, entrance, maze.isOuterPortal(entrance)})
							}
						}
					}
				}
			}
		}
	}

	return maze
}

func (m *maze) getNeighbors(s state, maxLevel int) []state {
	neighbors := make([]state, 0)

	for _, dir := range gridutil.Get4Directions() {
		newPos := gridutil.Coordinate{Row: s.pos.Row + dir.Row, Col: s.pos.Col + dir.Col}
		if val, exists := m.grid.GetC(newPos); exists && val == '.' {
			neighbors = append(neighbors, state{newPos, s.level})
		}
	}

	for label, portals := range m.portals {
		if len(portals) != 2 {
			continue
		}

		var currentPortal *portal
		var otherPortal *portal
		for i := range portals {
			if portals[i].entrance == s.pos {
				currentPortal = &portals[i]
				otherPortal = &portals[1-i]
				break
			}
		}

		if currentPortal == nil {
			continue
		}

		if label == "AA" || label == "ZZ" {
			if s.level == 0 {
				if label == "ZZ" {
					neighbors = append(neighbors, state{otherPortal.entrance, 0})
				}
			}
			continue
		}

		if currentPortal.isOuter {
			if s.level > 0 {
				neighbors = append(neighbors, state{otherPortal.entrance, s.level - 1})
			}
		} else {
			if s.level < maxLevel {
				neighbors = append(neighbors, state{otherPortal.entrance, s.level + 1})
			}
		}
	}

	return neighbors
}

func (m *maze) findShortestPathPart1() int {
	start := m.portals["AA"][0].entrance
	end := m.portals["ZZ"][0].entrance

	queue := container.NewQueue[gridutil.Coordinate]()
	visited := container.NewSet[gridutil.Coordinate]()
	distance := make(map[gridutil.Coordinate]int)
	queue.Push(start)
	visited.Add(start)
	distance[start] = 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current == end {
			return distance[current]
		}

		for _, dir := range gridutil.Get4Directions() {
			next := gridutil.Coordinate{Row: current.Row + dir.Row, Col: current.Col + dir.Col}
			if val, exists := m.grid.GetC(next); exists && val == '.' && !visited.Contains(next) {
				visited.Add(next)
				distance[next] = distance[current] + 1
				queue.Push(next)
			}
		}

		for _, portals := range m.portals {
			if len(portals) != 2 {
				continue
			}

			if current == portals[0].entrance {
				next := portals[1].entrance
				if !visited.Contains(next) {
					visited.Add(next)
					distance[next] = distance[current] + 1
					queue.Push(next)
				}
			} else if current == portals[1].entrance {
				next := portals[0].entrance
				if !visited.Contains(next) {
					visited.Add(next)
					distance[next] = distance[current] + 1
					queue.Push(next)
				}
			}
		}
	}

	return -1
}

func (m *maze) findShortestPathPart2() int {
	start := state{m.portals["AA"][0].entrance, 0}
	end := state{m.portals["ZZ"][0].entrance, 0}
	maxLevel := 25

	queue := container.NewQueue[state]()
	visited := container.NewSet[state]()
	distance := make(map[state]int)
	queue.Push(start)
	visited.Add(start)
	distance[start] = 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current == end {
			return distance[current]
		}

		for _, next := range m.getNeighbors(current, maxLevel) {
			if !visited.Contains(next) {
				visited.Add(next)
				distance[next] = distance[current] + 1
				queue.Push(next)
			}
		}
	}

	return -1
}
