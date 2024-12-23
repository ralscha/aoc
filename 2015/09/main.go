package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/graphutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	graph := graphutil.NewGraph()

	// Build graph from input
	for _, line := range lines {
		splitted := strings.Fields(line)
		from := splitted[0]
		to := splitted[2]
		weight := conv.MustAtoi(splitted[4])

		graph.AddNode(from)
		graph.AddNode(to)
		graph.AddEdge(from, to, weight)
		graph.AddEdge(to, from, weight) // Add reverse edge since it's undirected
	}

	// Get all locations (nodes)
	locations := make([]string, 0)
	for _, node := range graph.GetNeighbors("") { // Empty string gets all nodes
		locations = append(locations, node.ID)
	}

	fmt.Println("Part 1", shortestRoute(graph, locations))
	fmt.Println("Part 2", longestRoute(graph, locations))
}

func shortestRoute(graph *graphutil.Graph, locations []string) int {
	shortestDistance := math.MaxInt32
	for _, route := range mathx.Permutations(locations) {
		distance := 0
		for i := range len(route) - 1 {
			// Find edge between current and next location
			for _, neighbor := range graph.GetNeighbors(route[i]) {
				if neighbor.ID == route[i+1] {
					distance += neighbor.Edge.Weight
					break
				}
			}
		}
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	return shortestDistance
}

func longestRoute(graph *graphutil.Graph, locations []string) int {
	longestDistance := math.MinInt32
	for _, route := range mathx.Permutations(locations) {
		distance := 0
		for i := range len(route) - 1 {
			// Find edge between current and next location
			for _, neighbor := range graph.GetNeighbors(route[i]) {
				if neighbor.ID == route[i+1] {
					distance += neighbor.Edge.Weight
					break
				}
			}
		}
		if distance > longestDistance {
			longestDistance = distance
		}
	}
	return longestDistance
}
