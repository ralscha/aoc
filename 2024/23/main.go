package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
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

	connections := make(map[string]map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		if connections[a] == nil {
			connections[a] = make(map[string]bool)
		}
		if connections[b] == nil {
			connections[b] = make(map[string]bool)
		}
		connections[a][b] = true
		connections[b][a] = true
	}

	triangles := findTriangles(connections)

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

	connections := make(map[string]map[string]bool)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		if connections[a] == nil {
			connections[a] = make(map[string]bool)
		}
		if connections[b] == nil {
			connections[b] = make(map[string]bool)
		}
		connections[a][b] = true
		connections[b][a] = true
	}

	largestClique := findLargestClique(connections)
	slices.Sort(largestClique)
	password := strings.Join(largestClique, ",")
	fmt.Println("Part 2", password)
}

func findLargestClique(connections map[string]map[string]bool) []string {
	var largestClique []string

	var bronKerboschPivot func(r, p, x map[string]bool)
	bronKerboschPivot = func(r, p, x map[string]bool) {
		if len(p) == 0 && len(x) == 0 {
			if len(r) > len(largestClique) {
				largestClique = make([]string, 0, len(r))
				for node := range r {
					largestClique = append(largestClique, node)
				}
			}
			return
		}

		u := getPivot(p, x, connections)
		for v := range difference(p, connections[u]) {
			rv := union(r, v)
			pv := intersection(p, connections[v])
			xv := intersection(x, connections[v])
			bronKerboschPivot(rv, pv, xv)
			p = difference(p, map[string]bool{v: true})
			x = union(x, v)
		}
	}

	nodes := make(map[string]bool)
	for node := range connections {
		nodes[node] = true
	}
	bronKerboschPivot(make(map[string]bool), nodes, make(map[string]bool))

	return largestClique
}

func union(a map[string]bool, b string) map[string]bool {
	result := make(map[string]bool)
	for k := range a {
		result[k] = true
	}
	result[b] = true
	return result
}

func getPivot(p, x map[string]bool, connections map[string]map[string]bool) string {
	var pivot string
	maxDegree := -1
	for node := range unionMaps(p, x) {
		degree := len(connections[node])
		if degree > maxDegree {
			maxDegree = degree
			pivot = node
		}
	}
	return pivot
}

func unionMaps(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range a {
		result[k] = true
	}
	for k := range b {
		result[k] = true
	}
	return result
}

func intersection(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range a {
		if b[k] {
			result[k] = true
		}
	}
	return result
}

func difference(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k := range a {
		if !b[k] {
			result[k] = true
		}
	}
	return result
}

func findTriangles(connections map[string]map[string]bool) [][]string {
	var triangles [][]string
	for a := range connections {
		for b := range connections[a] {
			if a >= b {
				continue
			}
			for c := range connections[b] {
				if b >= c || !connections[a][c] {
					continue
				}
				triangles = append(triangles, []string{a, b, c})
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
