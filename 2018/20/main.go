package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2018, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	adj := buildGraph(input)
	distances := bfs(gridutil.Coordinate{}, adj)

	maxDistance := 0
	for _, dist := range distances {
		if dist > maxDistance {
			maxDistance = dist
		}
	}

	fmt.Println("Part 1", maxDistance)
}

func part2(input string) {
	adj := buildGraph(input)
	distances := bfs(gridutil.Coordinate{}, adj)

	count := 0
	for _, dist := range distances {
		if dist >= 1000 {
			count++
		}
	}

	fmt.Println("Part 2", count)
}

func buildGraph(regex string) map[gridutil.Coordinate]*container.Set[gridutil.Coordinate] {
	adj := make(map[gridutil.Coordinate]*container.Set[gridutil.Coordinate])
	currentPositions := container.NewSet[gridutil.Coordinate]()
	currentPositions.Add(gridutil.Coordinate{})
	var stack []*container.Set[gridutil.Coordinate]

	for i := 1; i < len(regex)-1; i++ {
		char := rune(regex[i])
		nextPositions := container.NewSet[gridutil.Coordinate]()

		switch char {
		case 'N':
			for _, pos := range currentPositions.Values() {
				nextPos := gridutil.Coordinate{Col: pos.Col, Row: pos.Row + 1}
				addEdge(adj, pos, nextPos)
				nextPositions.Add(nextPos)
			}
			currentPositions = nextPositions
		case 'S':
			for _, pos := range currentPositions.Values() {
				nextPos := gridutil.Coordinate{Col: pos.Col, Row: pos.Row - 1}
				addEdge(adj, pos, nextPos)
				nextPositions.Add(nextPos)
			}
			currentPositions = nextPositions
		case 'E':
			for _, pos := range currentPositions.Values() {
				nextPos := gridutil.Coordinate{Col: pos.Col + 1, Row: pos.Row}
				addEdge(adj, pos, nextPos)
				nextPositions.Add(nextPos)
			}
			currentPositions = nextPositions
		case 'W':
			for _, pos := range currentPositions.Values() {
				nextPos := gridutil.Coordinate{Col: pos.Col - 1, Row: pos.Row}
				addEdge(adj, pos, nextPos)
				nextPositions.Add(nextPos)
			}
			currentPositions = nextPositions
		case '(':
			stack = append(stack, currentPositions.Copy())
		case '|':
			currentPositions = stack[len(stack)-1].Copy()
		case ')':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, pos := range top.Values() {
				currentPositions.Add(pos)
			}
		}
	}
	return adj
}

func addEdge(adj map[gridutil.Coordinate]*container.Set[gridutil.Coordinate], from, to gridutil.Coordinate) {
	if _, ok := adj[from]; !ok {
		adj[from] = container.NewSet[gridutil.Coordinate]()
	}
	adj[from].Add(to)

	if _, ok := adj[to]; !ok {
		adj[to] = container.NewSet[gridutil.Coordinate]()
	}
	adj[to].Add(from)
}

func bfs(start gridutil.Coordinate, adj map[gridutil.Coordinate]*container.Set[gridutil.Coordinate]) map[gridutil.Coordinate]int {
	distances := make(map[gridutil.Coordinate]int)
	queue := container.NewQueue[gridutil.Coordinate]()
	queue.Push(start)
	distances[start] = 0

	for !queue.IsEmpty() {
		current := queue.Pop()
		dist := distances[current]

		if neighbors, ok := adj[current]; ok {
			for _, neighbor := range neighbors.Values() {
				if _, visited := distances[neighbor]; !visited {
					distances[neighbor] = dist + 1
					queue.Push(neighbor)
				}
			}
		}
	}
	return distances
}
