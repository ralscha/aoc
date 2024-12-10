package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/graphutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	graph := createGraph(lines)
	maxHappiness(graph)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	graph := createGraph(lines)
	
	// Add "me" to the graph with 0 happiness to/from everyone
	for _, node := range graph.GetNeighbors("") { // Empty string gets all nodes
		graph.AddNode("me", nil)
		graph.AddEdge("me", node.ID, 0)
		graph.AddEdge(node.ID, "me", 0)
	}
	
	maxHappiness(graph)
}

func createGraph(lines []string) *graphutil.Graph {
	graph := graphutil.NewGraph()
	
	for _, line := range lines {
		splitted := strings.Fields(line)
		name := splitted[0]
		gainLoseValue := conv.MustAtoi(splitted[3])
		if splitted[2] == "lose" {
			gainLoseValue = -gainLoseValue
		}
		neighbour := splitted[10]
		neighbour = neighbour[:len(neighbour)-1]
		
		graph.AddNode(name, nil)
		graph.AddNode(neighbour, nil)
		graph.AddEdge(name, neighbour, gainLoseValue)
	}
	
	return graph
}

func maxHappiness(graph *graphutil.Graph) {
	// Get all names (nodes)
	names := make([]string, 0)
	for _, node := range graph.GetNeighbors("") { // Empty string gets all nodes
		names = append(names, node.ID)
	}

	perms := mathx.Permutations(names)
	maxHappiness := 0
	for _, perm := range perms {
		happiness := 0
		for i, name := range perm {
			// Get happiness values between current person and their neighbors in the seating
			neighbor1 := perm[(i-1+len(perm))%len(perm)]
			neighbor2 := perm[(i+1)%len(perm)]
			
			// Add happiness in both directions (person -> neighbor and neighbor -> person)
			for _, n := range graph.GetNeighbors(name) {
				if n.ID == neighbor1 || n.ID == neighbor2 {
					happiness += n.Edge.Weight
				}
			}
		}
		if happiness > maxHappiness {
			maxHappiness = happiness
		}
	}

	fmt.Println(maxHappiness)
}
