package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"slices"
)

type nanobot struct {
	pos    gridutil.Coordinate3D
	radius int
}

func main() {
	input, err := download.ReadInput(2018, 23)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parseNanobots(input string) []nanobot {
	lines := conv.SplitNewline(input)
	nanobots := make([]nanobot, len(lines))
	for i, line := range lines {
		var x, y, z, r int
		conv.MustSscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		nanobots[i] = nanobot{
			pos:    gridutil.Coordinate3D{X: x, Y: y, Z: z},
			radius: r,
		}
	}
	return nanobots
}

func part1(input string) {
	nanobots := parseNanobots(input)

	var strongestNanobot nanobot
	maxRadius := -1
	for _, bot := range nanobots {
		if bot.radius > maxRadius {
			maxRadius = bot.radius
			strongestNanobot = bot
		}
	}

	inRangeCount := 0
	for _, bot := range nanobots {
		distance := strongestNanobot.pos.ManhattanDistance(bot.pos)
		if distance <= strongestNanobot.radius {
			inRangeCount++
		}
	}

	fmt.Println("Part 1", inRangeCount)
}

func part2(input string) {
	nanobots := parseNanobots(input)

	neighbors := make(map[int][]int)
	for i, bot0 := range nanobots {
		for j, bot1 := range nanobots {
			if i == j || bot0.pos.ManhattanDistance(bot1.pos) > bot0.radius+bot1.radius {
				continue
			}
			neighbors[j] = append(neighbors[j], i)
		}
	}
	cliques := bronKerbosch(neighbors)
	clique := cliques[0]

	var maxMin int
	var origin gridutil.Coordinate3D
	for _, i := range clique {
		bot := nanobots[i]
		d := origin.ManhattanDistance(bot.pos) - bot.radius
		if d > maxMin {
			maxMin = d
		}
	}
	fmt.Println("Part 2", maxMin)
}

func bronKerbosch(g map[int][]int) [][]int {
	p := container.NewSet[int]()
	for v := range g {
		p.Add(v)
	}
	state := &bkState{g: g}
	r := container.NewSet[int]()
	x := container.NewSet[int]()
	bronKerbosch1(state, r, p, x)
	return state.maxCliques
}

type bkState struct {
	g          map[int][]int
	maxCliques [][]int
}

func bronKerbosch1(state *bkState, r, p, x *container.Set[int]) {
	if p.Len() == 0 && x.Len() == 0 {
		if len(state.maxCliques) > 0 && len(state.maxCliques[0]) > r.Len() {
			return
		}
		if len(state.maxCliques) > 0 && len(state.maxCliques[0]) < r.Len() {
			state.maxCliques = nil
		}
		clique := make([]int, 0, r.Len())
		for _, v := range r.Values() {
			clique = append(clique, v)
		}
		slices.Sort(clique)
		state.maxCliques = append(state.maxCliques, clique)
		return
	}
	u := -1
	if p.Len() > 0 {
		for _, v := range p.Values() {
			u = v
			break
		}
	} else {
		for _, v := range x.Values() {
			u = v
			break
		}
	}
	nu := state.g[u]
	nuSet := container.NewSet[int]()
	for _, uu := range nu {
		nuSet.Add(uu)
	}
	for _, v := range p.Values() {
		if nuSet.Contains(v) {
			continue
		}
		ns := state.g[v]
		p1 := container.NewSet[int]()
		x1 := container.NewSet[int]()
		for _, n := range ns {
			if p.Contains(n) {
				p1.Add(n)
			}
			if x.Contains(n) {
				x1.Add(n)
			}
		}
		r.Add(v)
		bronKerbosch1(state, r, p1, x1)
		r.Remove(v)
		p.Remove(v)
		x.Add(v)
	}
}
