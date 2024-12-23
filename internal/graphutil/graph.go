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

func (g *Graph) FindCycle(startID string) []string {
	visited := container.NewSet[string]()
	path := []string{}
	var paths [][]string

	var dfs func(current string)
	dfs = func(current string) {
		if visited.Contains(current) {
			// Found a cycle
			cycleStart := -1
			for i, node := range path {
				if node == current {
					cycleStart = i
					break
				}
			}
			if cycleStart != -1 {
				cycle := slices.Clone(path[cycleStart:])
				paths = append(paths, cycle)
			}
			return
		}

		visited.Add(current)
		path = append(path, current)

		for _, neighbor := range g.nodes[current] {
			dfs(neighbor.ID)
		}

		path = path[:len(path)-1]
		visited.Remove(current)
	}

	dfs(startID)

	if len(paths) > 0 {
		return paths[0]
	}
	return nil
}

func (g *Graph) FindAllPaths(start, end string, condition func(node *Node) bool) [][]string {
	var paths [][]string
	visited := container.NewSet[string]()
	currentPath := []string{start}

	var dfs func(current string)
	dfs = func(current string) {
		if current == end {
			path := slices.Clone(currentPath)
			paths = append(paths, path)
			return
		}

		visited.Add(current)
		for _, neighbor := range g.nodes[current] {
			if !visited.Contains(neighbor.ID) && (condition == nil || condition(neighbor)) {
				currentPath = append(currentPath, neighbor.ID)
				dfs(neighbor.ID)
				currentPath = currentPath[:len(currentPath)-1]
			}
		}
		visited.Remove(current)
	}

	dfs(start)
	return paths
}
