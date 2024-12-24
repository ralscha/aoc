package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2018, 25)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

type point struct {
	x, y, z, w int
}

func parsePoint(s string) point {
	var p point
	conv.MustSscanf(s, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.w)
	return p
}

func manhattanDistance(p1, p2 point) int {
	return mathx.Abs(p1.x-p2.x) + mathx.Abs(p1.y-p2.y) + mathx.Abs(p1.z-p2.z) + mathx.Abs(p1.w-p2.w)
}

type dsu struct {
	parent map[point]point
}

func newDSU(points []point) *dsu {
	parent := make(map[point]point)
	for _, p := range points {
		parent[p] = p
	}
	return &dsu{parent: parent}
}

func (dsu *dsu) find(p point) point {
	if dsu.parent[p] == p {
		return p
	}
	dsu.parent[p] = dsu.find(dsu.parent[p])
	return dsu.parent[p]
}

func (dsu *dsu) union(p1, p2 point) {
	root1 := dsu.find(p1)
	root2 := dsu.find(p2)
	if root1 != root2 {
		dsu.parent[root1] = root2
	}
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var points []point
	for _, line := range lines {
		points = append(points, parsePoint(line))
	}

	dsu := newDSU(points)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if manhattanDistance(points[i], points[j]) <= 3 {
				dsu.union(points[i], points[j])
			}
		}
	}

	numConstellations := 0
	seen := container.NewSet[point]()
	for _, p := range points {
		root := dsu.find(p)
		if !seen.Contains(root) {
			numConstellations++
			seen.Add(root)
		}
	}

	fmt.Println("Part 1", numConstellations)
}
