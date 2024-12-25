package graphutil

import (
	"aoc/internal/container"
	"maps"
	"slices"
)

type Node struct {
	ID   string
	Data interface{}
	Edge *Edge
}

type Edge struct {
	Weight int
}

type Graph struct {
	nodes map[string][]*Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string][]*Node),
	}
}

func (g *Graph) AddNode(id string) {
	if _, exists := g.nodes[id]; !exists {
		g.nodes[id] = make([]*Node, 0)
	}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	g.nodes[from] = append(g.nodes[from], &Node{
		ID:   to,
		Edge: &Edge{Weight: weight},
	})
}

// GetNeighbors returns all neighbors of a node.
// If nodeID is empty, returns all nodes in the graph.
func (g *Graph) GetNeighbors(nodeID string) []*Node {
	if nodeID == "" {
		// Special case: return all nodes
		visited := container.NewSet[string]()
		allNodes := make([]*Node, 0)

		for k, v := range maps.All(g.nodes) {
			id, neighbors := k, v
			if !visited.Contains(id) {
				visited.Add(id)
				allNodes = append(allNodes, &Node{ID: id})
			}
			for _, neighbor := range neighbors {
				if !visited.Contains(neighbor.ID) {
					visited.Add(neighbor.ID)
					allNodes = append(allNodes, &Node{ID: neighbor.ID})
				}
			}
		}
		return allNodes
	}
	return slices.Clone(g.nodes[nodeID])
}
