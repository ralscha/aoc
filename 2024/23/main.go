package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/graphutil"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	graph := buildGraph(lines)

	triangles := findTriangles(graph)
	count := 0
	for _, triangle := range triangles {
		if startsWithT(triangle) {
			count++
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	graph := buildGraph(lines)

	largestClique := findLargestClique(graph)
	slices.Sort(largestClique)
	password := strings.Join(largestClique, ",")
	fmt.Println("Part 2", password)
}

func buildGraph(lines []string) *graphutil.Graph {
	graph := graphutil.NewGraph()
	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		graph.AddNode(a)
		graph.AddNode(b)
		graph.AddEdge(a, b, 1)
		graph.AddEdge(b, a, 1)
	}
	return graph
}

func findLargestClique(g *graphutil.Graph) []string {
	var largestClique []string

	var bronKerboschPivot func(r, p, x *container.Set[string])
	bronKerboschPivot = func(r, p, x *container.Set[string]) {
		if p.Len() == 0 && x.Len() == 0 {
			if len(r.Values()) > len(largestClique) {
				largestClique = r.Values()
			}
			return
		}

		u := getPivot(p, x, g)
		neighbors := container.NewSet[string]()
		for _, n := range g.GetNeighbors(u) {
			neighbors.Add(n.ID)
		}

		pCopy := container.NewSet[string]()
		for _, v := range p.Values() {
			if !neighbors.Contains(v) {
				pCopy.Add(v)
			}
		}

		for _, v := range pCopy.Values() {
			rNew := r.Copy()
			rNew.Add(v)

			pNew := container.NewSet[string]()
			xNew := container.NewSet[string]()

			vNeighbors := container.NewSet[string]()
			for _, n := range g.GetNeighbors(v) {
				vNeighbors.Add(n.ID)
			}

			for _, pv := range p.Values() {
				if vNeighbors.Contains(pv) {
					pNew.Add(pv)
				}
			}
			for _, xv := range x.Values() {
				if vNeighbors.Contains(xv) {
					xNew.Add(xv)
				}
			}

			bronKerboschPivot(rNew, pNew, xNew)
			p.Remove(v)
			x.Add(v)
		}
	}

	nodes := container.NewSet[string]()
	for _, node := range g.GetNeighbors("") {
		nodes.Add(node.ID)
	}

	bronKerboschPivot(container.NewSet[string](), nodes, container.NewSet[string]())
	return largestClique
}

func getPivot(p, x *container.Set[string], g *graphutil.Graph) string {
	var pivot string
	maxDegree := -1

	union := container.NewSet[string]()
	for _, v := range p.Values() {
		union.Add(v)
	}
	for _, v := range x.Values() {
		union.Add(v)
	}

	for _, node := range union.Values() {
		degree := len(g.GetNeighbors(node))
		if degree > maxDegree {
			maxDegree = degree
			pivot = node
		}
	}
	return pivot
}

func findTriangles(g *graphutil.Graph) [][]string {
	var triangles [][]string
	for _, nodeA := range g.GetNeighbors("") {
		for _, nodeB := range g.GetNeighbors(nodeA.ID) {
			if nodeA.ID >= nodeB.ID {
				continue
			}
			for _, nodeC := range g.GetNeighbors(nodeB.ID) {
				if nodeB.ID >= nodeC.ID {
					continue
				}
				hasEdge := false
				for _, neighbor := range g.GetNeighbors(nodeA.ID) {
					if neighbor.ID == nodeC.ID {
						hasEdge = true
						break
					}
				}
				if hasEdge {
					triangles = append(triangles, []string{nodeA.ID, nodeB.ID, nodeC.ID})
				}
			}
		}
	}
	return triangles
}

func startsWithT(triangle []string) bool {
	for _, computer := range triangle {
		if strings.HasPrefix(computer, "t") {
			return true
		}
	}
	return false
}
