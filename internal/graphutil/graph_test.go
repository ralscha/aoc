package graphutil

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	// Test adding nodes
	g.AddNode("A", nil)
	g.AddNode("B", nil)
	g.AddNode("C", nil)

	if len(g.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(g.Nodes))
	}

	// Test adding edges
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 1)
	g.AddEdge("C", "A", 1)

	if len(g.Edges["A"]) != 1 {
		t.Errorf("Expected 1 edge from A, got %d", len(g.Edges["A"]))
	}
}

func TestGetNeighbors(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", nil)
	g.AddNode("B", nil)
	g.AddNode("C", nil)
	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 1)

	neighbors := g.GetNeighbors("A")
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors, got %d", len(neighbors))
	}

	neighborIDs := make(map[string]bool)
	for _, n := range neighbors {
		neighborIDs[n.ID] = true
	}

	if !neighborIDs["B"] || !neighborIDs["C"] {
		t.Error("Expected neighbors B and C")
	}
}

func TestFindCycle(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Graph
		start     string
		wantLen   int
		wantNil   bool
		wantNodes []string
	}{
		{
			name: "simple cycle",
			setup: func() *Graph {
				g := NewGraph()
				g.AddNode("A", nil)
				g.AddNode("B", nil)
				g.AddNode("C", nil)
				g.AddEdge("A", "B", 1)
				g.AddEdge("B", "C", 1)
				g.AddEdge("C", "A", 1)
				return g
			},
			start:     "A",
			wantLen:   4, // A->B->C->A
			wantNodes: []string{"A", "B", "C", "A"},
			wantNil:   false,
		},
		{
			name: "no cycle",
			setup: func() *Graph {
				g := NewGraph()
				g.AddNode("A", nil)
				g.AddNode("B", nil)
				g.AddEdge("A", "B", 1)
				return g
			},
			start:   "A",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.setup()
			cycle := g.FindCycle(tt.start)
			if tt.wantNil && cycle != nil {
				t.Errorf("Expected nil cycle, got %v", cycle)
			}
			if !tt.wantNil && cycle == nil {
				t.Error("Expected non-nil cycle, got nil")
			}
			if !tt.wantNil && len(cycle) != tt.wantLen {
				t.Errorf("Expected cycle length %d, got %d", tt.wantLen, len(cycle))
			}
			if !tt.wantNil && !reflect.DeepEqual(cycle, tt.wantNodes) {
				t.Errorf("Expected cycle %v, got %v", tt.wantNodes, cycle)
			}
		})
	}
}

func TestFindAllPaths(t *testing.T) {
	g := NewGraph()
	g.AddNode("A", nil)
	g.AddNode("B", nil)
	g.AddNode("C", nil)
	g.AddNode("D", nil)
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 1)
	g.AddEdge("B", "D", 1)
	g.AddEdge("A", "D", 1)

	paths := g.FindAllPaths("A", "D", func(node *Node) bool { return true })

	expectedPaths := [][]string{
		{"A", "D"},
		{"A", "B", "D"},
	}

	if len(paths) != len(expectedPaths) {
		t.Errorf("Expected %d paths, got %d", len(expectedPaths), len(paths))
	}

	// Helper function to check if a path exists in expected paths
	containsPath := func(path []string) bool {
		for _, exp := range expectedPaths {
			if reflect.DeepEqual(path, exp) {
				return true
			}
		}
		return false
	}

	for _, path := range paths {
		if !containsPath(path) {
			t.Errorf("Unexpected path: %v", path)
		}
	}
}

func TestParseNetwork(t *testing.T) {
	input := []string{
		"AAA = (BBB, CCC)",
		"BBB = (DDD, EEE)",
		"CCC = (ZZZ, GGG)",
	}

	g := ParseNetwork(input)

	// Check nodes
	expectedNodes := []string{"AAA", "BBB", "CCC"}
	for _, nodeID := range expectedNodes {
		if _, exists := g.Nodes[nodeID]; !exists {
			t.Errorf("Expected node %s not found", nodeID)
		}
	}

	// Check edges
	tests := []struct {
		from string
		to   []string
	}{
		{"AAA", []string{"BBB", "CCC"}},
		{"BBB", []string{"DDD", "EEE"}},
		{"CCC", []string{"ZZZ", "GGG"}},
	}

	for _, tt := range tests {
		edges := g.Edges[tt.from]
		if len(edges) != len(tt.to) {
			t.Errorf("Expected %d edges from %s, got %d", len(tt.to), tt.from, len(edges))
		}
		for _, to := range tt.to {
			if _, exists := edges[to]; !exists {
				t.Errorf("Expected edge from %s to %s not found", tt.from, to)
			}
		}
	}
}
