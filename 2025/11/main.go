package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	adj := parseGraph(input)
	count := countPaths(adj, "you", "out")
	fmt.Println("Part 1", count)
}

func part2(input string) {
	adj := parseGraph(input)
	count := countPaths(adj, "svr", "out", "dac", "fft")
	fmt.Println("Part 2", count)

}

func parseGraph(input string) map[string][]string {
	adj := make(map[string][]string)
	for _, line := range conv.SplitNewline(input) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		left, right, found := strings.Cut(line, ":")
		if !found {
			log.Fatalf("invalid line: %q", line)
		}

		from := strings.TrimSpace(left)
		if from == "" {
			log.Fatalf("missing source device in line: %q", line)
		}

		right = strings.TrimSpace(right)
		if right == "" {
			adj[from] = []string{}
			continue
		}

		adj[from] = strings.Fields(right)
	}

	return adj
}

func countPaths(adj map[string][]string, start, target string, required ...string) int64 {
	reqBit := make(map[string]uint64, len(required))
	for i, node := range required {
		reqBit[node] = 1 << i
	}

	allMask := uint64(0)
	for i := range required {
		allMask |= 1 << i
	}

	type state struct {
		node string
		mask uint64
	}

	memo := make(map[state]int64)
	visiting := container.NewSet[state]()

	var dfs func(node string, mask uint64) int64
	dfs = func(node string, mask uint64) int64 {
		if bit, ok := reqBit[node]; ok {
			mask |= bit
		}

		if node == target {
			if mask == allMask {
				return 1
			}
			return 0
		}

		current := state{node: node, mask: mask}
		if value, ok := memo[current]; ok {
			return value
		}

		visiting.Add(current)
		total := int64(0)
		for _, next := range adj[node] {
			total += dfs(next, mask)
		}
		visiting.Remove(current)

		memo[current] = total
		return total
	}

	return dfs(start, 0)
}
