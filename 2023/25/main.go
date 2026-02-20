package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

type edge struct {
	src, dest int
}

type graph struct {
	V     int
	E     int
	edges []edge
}

type subset struct {
	parent, rank int
}

type parsedGraph struct {
	nodes []string
	edges [][2]int
}

func parseGraph(strs []string) parsedGraph {
	nodes := make([]string, 0)
	edges := make([][2]int, 0)
	nodeSet := container.NewSet[string]()
	indexes := make(map[string]int)

	add := func(u string) int {
		if !nodeSet.Contains(u) {
			nodeSet.Add(u)
			indexes[u] = len(nodes)
			nodes = append(nodes, u)
		}
		return indexes[u]
	}

	for _, str := range strs {
		parts := strings.Split(str, ": ")
		left := parts[0]
		u := add(left)

		right := strings.SplitSeq(parts[1], " ")
		for c := range right {
			v := add(c)
			edges = append(edges, [2]int{u, v})
		}
	}
	return parsedGraph{nodes: nodes, edges: edges}
}

func kargerMinCut(graph *graph) (int, []int) {
	subsets := make([]subset, graph.V)
	for v := range subsets {
		subsets[v] = subset{parent: v, rank: 0}
	}

	vertices := graph.V
	for vertices > 2 {
		i := rand.Intn(graph.E)
		subset1 := find(subsets, graph.edges[i].src)
		subset2 := find(subsets, graph.edges[i].dest)

		if subset1 != subset2 {
			vertices--
			union(subsets, subset1, subset2)
		}
	}

	cutedges := 0
	edgeSet := container.NewSet[edge]()
	for _, e := range graph.edges {
		if find(subsets, e.src) != find(subsets, e.dest) {
			cutedges++
			edgeSet.Add(e)
		}
	}

	components := make([]int, graph.V)
	for i := range components {
		components[i] = find(subsets, i)
	}

	return cutedges, components
}

func find(subsets []subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = find(subsets, subsets[i].parent)
	}
	return subsets[i].parent
}

func union(subsets []subset, x, y int) {
	xroot := find(subsets, x)
	yroot := find(subsets, y)

	if subsets[xroot].rank < subsets[yroot].rank {
		subsets[xroot].parent = yroot
	} else if subsets[xroot].rank > subsets[yroot].rank {
		subsets[yroot].parent = xroot
	} else {
		subsets[yroot].parent = xroot
		subsets[xroot].rank++
	}
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	g := parseGraph(lines)
	V := len(g.nodes)
	E := len(g.edges)
	gr := graph{V: V, E: E, edges: make([]edge, E)}

	for i, e := range g.edges {
		gr.edges[i] = edge{src: e[0], dest: e[1]}
	}

	k := 0
	var res int
	var components []int
	found := false

	// Try Karger's algorithm multiple times until we find a minimum cut of 3
	for k < 1000 {
		k++
		res, components = kargerMinCut(&gr)
		if res == 3 {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("no solution found")
		return
	}

	// Count size of each component
	componentSizes := container.NewBag[int]()
	for _, c := range components {
		componentSizes.Add(c)
	}

	// Calculate product of component sizes
	product := 1
	for _, count := range componentSizes.Values() {
		product *= count
	}

	fmt.Println(product)
}
