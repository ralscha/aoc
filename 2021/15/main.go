package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
)

type point struct {
	x, y int
}

var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}

func main() {
	input, err := download.ReadInput(2021, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := make(map[point]int)
	for x, row := range lines {
		if row == "" {
			continue
		}
		for y, val := range row {
			grid[point{x: x, y: y}] = int(val - '0')
		}
	}

	var maxx, maxy = len(lines[0]) - 1, len(lines) - 1
	start := point{0, 0}
	target := point{maxx, maxy}

	lowestRiskAt := make(map[point]int)
	pq := container.NewPriorityQueue[point]()
	pq.Push(start, 0)

	for !pq.IsEmpty() {
		head := pq.Pop()
		for i := 0; i < 4; i++ {
			next := point{head.x + dx[i], head.y + dy[i]}
			if next.x > target.x || next.x < 0 || next.y > target.y || next.y < 0 {
				continue
			}

			nextRisk := lowestRiskAt[head] + grid[next]
			if sAt, ok := lowestRiskAt[next]; ok && sAt <= nextRisk {
				continue
			}
			lowestRiskAt[next] = nextRisk
			pq.Push(next, nextRisk)

		}
	}

	fmt.Println(lowestRiskAt[target])
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := make(map[point]int)

	for x, row := range lines {
		for y, val := range row {
			grid[point{x: x, y: y}] = int(val - '0')
		}
	}

	var maxx, maxy = len(lines[0]) - 1, len(lines) - 1
	start := point{0, 0}
	target := point{(maxx+1)*5 - 1, (maxy+1)*5 - 1}

	risk := func(pos point) int {
		og := point{pos.x % (maxx + 1), pos.y % (maxy + 1)}
		risk := grid[og] +
			(pos.x)/(maxx+1) + (pos.y)/(maxy+1)
		if risk > 9 {
			return risk - 9
		}
		return risk
	}

	lowestRiskAt := make(map[point]int)
	pq := container.NewPriorityQueue[point]()
	pq.Push(start, 0)

	for pq.Len() > 0 {
		head := pq.Pop()
		for i := 0; i < 4; i++ {
			next := point{head.x + dx[i], head.y + dy[i]}
			if next.x > target.x || next.x < 0 || next.y > target.y || next.y < 0 {
				continue
			}
			nextRisk := lowestRiskAt[head] + risk(next)
			if sAt, ok := lowestRiskAt[next]; ok && sAt <= nextRisk {
				continue
			}
			lowestRiskAt[next] = nextRisk
			pq.Push(next, nextRisk)
		}
	}
	fmt.Println(lowestRiskAt[target])
}
