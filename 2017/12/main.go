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
	input, err := download.ReadInput(2017, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func buildAdjacencyList(input string) map[int][]int {
	adj := make(map[int][]int)
	lines := conv.SplitNewline(input)
	for _, line := range lines {
		parts := strings.Split(line, " <-> ")
		from := conv.MustAtoi(parts[0])
		toStrs := strings.Split(parts[1], ", ")
		var to []int
		for _, s := range toStrs {
			t := conv.MustAtoi(s)
			to = append(to, t)
		}
		adj[from] = to
	}
	return adj
}

func part1(input string) {
	adj := buildAdjacencyList(input)

	visited := container.NewSet[int]()
	queue := container.NewQueue[int]()

	queue.Push(0)
	visited.Add(0)

	for !queue.IsEmpty() {
		curr := queue.Pop()
		for _, neighbor := range adj[curr] {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(neighbor)
			}
		}
	}

	fmt.Println("Part 1", visited.Len())
}

func part2(input string) {
	adj := buildAdjacencyList(input)

	visited := container.NewSet[int]()
	groups := 0

	for program := range adj {
		if !visited.Contains(program) {
			groups++
			queue := container.NewQueue[int]()
			queue.Push(program)
			visited.Add(program)

			for !queue.IsEmpty() {
				curr := queue.Pop()
				for _, neighbor := range adj[curr] {
					if !visited.Contains(neighbor) {
						visited.Add(neighbor)
						queue.Push(neighbor)
					}
				}
			}
		}
	}

	fmt.Println("Part 2", groups)
}
