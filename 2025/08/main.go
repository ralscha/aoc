package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"slices"
	"sort"
)

func main() {
	input, err := download.ReadInput(2025, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type uf struct {
	parent, size []int
	components   int
}

func newUF(n int) *uf {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &uf{parent, size, n}
}

func (u *uf) find(x int) int {
	for u.parent[x] != x {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *uf) union(a, b int) bool {
	ra, rb := u.find(a), u.find(b)
	if ra == rb {
		return false
	}
	if u.size[ra] < u.size[rb] {
		ra, rb = rb, ra
	}
	u.parent[rb] = ra
	u.size[ra] += u.size[rb]
	u.components--
	return true
}

type edge struct {
	i, j int
	dist int64
}

func solve(input string) ([]gridutil.Coordinate3D, []edge) {
	lines := conv.SplitNewline(input)
	var coords []gridutil.Coordinate3D
	for _, line := range lines {
		if line == "" {
			continue
		}
		p := conv.ToIntSliceComma(line)
		coords = append(coords, gridutil.Coordinate3D{X: p[0], Y: p[1], Z: p[2]})
	}

	n := len(coords)
	edges := make([]edge, 0, n*(n-1)/2)
	for i := range n {
		for j := i + 1; j < n; j++ {
			d := coords[i].Sub(coords[j])
			dist := int64(d.X)*int64(d.X) + int64(d.Y)*int64(d.Y) + int64(d.Z)*int64(d.Z)
			edges = append(edges, edge{i, j, dist})
		}
	}
	sort.Slice(edges, func(a, b int) bool { return edges[a].dist < edges[b].dist })
	return coords, edges
}

func part1(input string) {
	coords, edges := solve(input)
	u := newUF(len(coords))
	for k := range 1000 {
		u.union(edges[k].i, edges[k].j)
	}

	sizes := make(map[int]int)
	for i := range coords {
		sizes[u.find(i)]++
	}
	top := make([]int, 0, len(sizes))
	for _, s := range sizes {
		top = append(top, s)
	}
	slices.SortFunc(top, func(a, b int) int { return b - a })
	fmt.Println("Part 1", top[0]*top[1]*top[2])
}

func part2(input string) {
	coords, edges := solve(input)
	u := newUF(len(coords))
	for _, e := range edges {
		if u.union(e.i, e.j) && u.components == 1 {
			fmt.Println("Part 2", coords[e.i].X*coords[e.j].X)
			return
		}
	}
}
