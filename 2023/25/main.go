package main

import (
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
	var nodes []string
	var edges [][2]int
	indexes := make(map[string]int)

	add := func(u string) int {
		if _, exists := indexes[u]; !exists {
			indexes[u] = len(nodes)
			nodes = append(nodes, u)
		}
		return indexes[u]
	}

	for _, str := range strs {
		parts := strings.Split(str, ": ")
		left := parts[0]
		u := add(left)

		right := strings.Split(parts[1], " ")
		for _, c := range right {
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
	for _, edge := range graph.edges {
		if find(subsets, edge.src) != find(subsets, edge.dest) {
			cutedges++
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

	list := make(map[int]int)
	for _, c := range components {
		list[c]++
	}

	product := 1
	for _, count := range list {
		product *= count
	}

	fmt.Println(product)
}
