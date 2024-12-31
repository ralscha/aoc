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

func reachableKeys(grid *gridutil.Grid2D[rune], start gridutil.Coordinate, haveKeys int) map[rune]struct {
	dist int
	pos  gridutil.Coordinate
} {
	type queueItem struct {
		pos  gridutil.Coordinate
		dist int
	}

	queue := container.NewQueue[queueItem]()
	queue.Push(queueItem{pos: start, dist: 0})
	distances := make(map[gridutil.Coordinate]int)
	distances[start] = 0
	reachable := make(map[rune]struct {
		dist int
		pos  gridutil.Coordinate
	})

	for !queue.IsEmpty() {
		current := queue.Pop()

		for _, dir := range gridutil.Get4Directions() {
			newPos := gridutil.Coordinate{Row: current.pos.Row + dir.Row, Col: current.pos.Col + dir.Col}

			if _, visited := distances[newPos]; visited {
				continue
			}

			if ch, ok := grid.GetC(newPos); ok && ch != '#' {
				distances[newPos] = current.dist + 1

				if ch >= 'A' && ch <= 'Z' && (haveKeys&(1<<(ch-'A'))) == 0 {
					continue
				}

				if ch >= 'a' && ch <= 'z' && (haveKeys&(1<<(ch-'a'))) == 0 {
					reachable[ch] = struct {
						dist int
						pos  gridutil.Coordinate
					}{dist: current.dist + 1, pos: newPos}
				} else {
					queue.Push(queueItem{pos: newPos, dist: current.dist + 1})
				}
			}
		}
	}

	return reachable
}

func reachableKeys4(grid *gridutil.Grid2D[rune], positions [4]gridutil.Coordinate, haveKeys int) map[rune]struct {
	dist     int
	pos      gridutil.Coordinate
	robotIdx int
} {
	allReachable := make(map[rune]struct {
		dist     int
		pos      gridutil.Coordinate
		robotIdx int
	})

	for i, pos := range positions {
		for key, info := range reachableKeys(grid, pos, haveKeys) {
			allReachable[key] = struct {
				dist     int
				pos      gridutil.Coordinate
				robotIdx int
			}{
				dist:     info.dist,
				pos:      info.pos,
				robotIdx: i,
			}
		}
	}

	return allReachable
}

type cacheKey struct {
	positions [4]gridutil.Coordinate
	haveKeys  int
}

func minSteps(grid *gridutil.Grid2D[rune], positions [4]gridutil.Coordinate, haveKeys int, cache map[cacheKey]int) int {
	key := cacheKey{positions: positions, haveKeys: haveKeys}
	if cached, ok := cache[key]; ok {
		return cached
	}

	reachable := reachableKeys4(grid, positions, haveKeys)
	if len(reachable) == 0 {
		return 0
	}

	minDist := 1<<31 - 1
	for key, info := range reachable {
		newPositions := positions
		newPositions[info.robotIdx] = info.pos
		newKeys := haveKeys | (1 << (key - 'a'))

		dist := info.dist + minSteps(grid, newPositions, newKeys, cache)
		if dist < minDist {
			minDist = dist
		}
	}

	cache[key] = minDist
	return minDist
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	var start gridutil.Coordinate
	for row := range grid.Height() {
		for col := range grid.Width() {
			if value, _ := grid.Get(row, col); value == '@' {
				start = gridutil.Coordinate{Row: row, Col: col}
				break
			}
		}
	}

	grid.Set(start.Row, start.Col, '#')
	grid.Set(start.Row-1, start.Col, '#')
	grid.Set(start.Row+1, start.Col, '#')
	grid.Set(start.Row, start.Col-1, '#')
	grid.Set(start.Row, start.Col+1, '#')

	positions := [4]gridutil.Coordinate{
		{Row: start.Row - 1, Col: start.Col - 1},
		{Row: start.Row - 1, Col: start.Col + 1},
		{Row: start.Row + 1, Col: start.Col - 1},
		{Row: start.Row + 1, Col: start.Col + 1},
	}
	for _, pos := range positions {
		grid.Set(pos.Row, pos.Col, '@')
	}

	cache := make(map[cacheKey]int)
	result := minSteps(&grid, positions, 0, cache)
	fmt.Println("Part 2", result)
}
