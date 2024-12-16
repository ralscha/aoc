package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
)

type State struct {
	pos       gridutil.Coordinate
	direction gridutil.Direction
	score     int
	path      []gridutil.Coordinate
}

func getKey(pos gridutil.Coordinate, dir gridutil.Direction) string {
	return fmt.Sprintf("%d,%d,%d", pos.Row, pos.Col, dir)
}

func main() {
	input, err := download.ReadInput(2024, 16)
	if err != nil {
		panic(err)
	}
	part1and2(input)
}

type pathfindResult struct {
	score int
	paths [][]gridutil.Coordinate
}

func pathfind(grid *gridutil.Grid2D[rune], start, end gridutil.Coordinate, lowestScore int, collectPaths bool) pathfindResult {
	pq := container.NewPriorityQueue[State]()
	visited := make(map[string]int)
	var paths [][]gridutil.Coordinate
	bestScore := -1

	initialState := State{
		pos:       start,
		direction: gridutil.DirectionE,
		score:     0,
		path:      []gridutil.Coordinate{start},
	}
	pq.Push(initialState, 0)

	for !pq.IsEmpty() {
		state := pq.Pop()

		if lowestScore > 0 && state.score > lowestScore {
			continue
		}

		key := getKey(state.pos, state.direction)
		if bestScore, exists := visited[key]; exists && bestScore < state.score {
			continue
		}
		visited[key] = state.score

		if state.pos == end {
			if bestScore == -1 || state.score < bestScore {
				bestScore = state.score
			}
			if collectPaths && state.score == lowestScore {
				paths = append(paths, append([]gridutil.Coordinate(nil), state.path...))
			}
			if !collectPaths && !pq.IsEmpty() && pq.Peek().score > state.score {
				break
			}
			continue
		}

		next := gridutil.Coordinate{
			Row: state.pos.Row + state.direction.Row,
			Col: state.pos.Col + state.direction.Col,
		}

		if val, exists := grid.GetC(next); exists && val != '#' {
			newPath := make([]gridutil.Coordinate, len(state.path), len(state.path)+1)
			copy(newPath, state.path)
			newPath = append(newPath, next)

			pq.Push(State{
				pos:       next,
				direction: state.direction,
				score:     state.score + 1,
				path:      newPath,
			}, state.score+1)
		}

		pq.Push(State{
			pos:       state.pos,
			direction: gridutil.TurnRight(state.direction),
			score:     state.score + 1000,
			path:      state.path,
		}, state.score+1000)

		pq.Push(State{
			pos:       state.pos,
			direction: gridutil.TurnLeft(state.direction),
			score:     state.score + 1000,
			path:      state.path,
		}, state.score+1000)
	}

	return pathfindResult{score: bestScore, paths: paths}
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	var start, end gridutil.Coordinate
	for row := 0; row < grid.Height(); row++ {
		for col := 0; col < grid.Width(); col++ {
			if ch, ok := grid.Get(row, col); ok {
				if ch == 'S' {
					start = gridutil.Coordinate{Row: row, Col: col}
				} else if ch == 'E' {
					end = gridutil.Coordinate{Row: row, Col: col}
				}
			}
		}
	}

	result := pathfind(&grid, start, end, 0, false)
	fmt.Println("Part 1:", result.score)

	result = pathfind(&grid, start, end, result.score, true)
	uniquePaths := container.NewSet[string]()
	for _, path := range result.paths {
		for _, coord := range path {
			uniquePaths.Add(getKey(coord, gridutil.DirectionN))
		}
	}
	fmt.Println("Part 2:", uniquePaths.Len())
}
