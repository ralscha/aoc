package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2018, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var depth, targetX, targetY int
	conv.MustSscanf(lines[0], "depth: %d", &depth)
	conv.MustSscanf(lines[1], "target: %d,%d", &targetX, &targetY)

	riskLevel := 0
	cache := make(map[gridutil.Coordinate]int)

	for y := 0; y <= targetY; y++ {
		for x := 0; x <= targetX; x++ {
			erosionLevel := getErosionLevel(x, y, depth, targetX, targetY, cache)
			regionType := erosionLevel % 3
			riskLevel += regionType
		}
	}

	fmt.Println("Part 1", riskLevel)
}

func getErosionLevel(x, y int, depth int, targetX, targetY int, cache map[gridutil.Coordinate]int) int {
	coord := gridutil.Coordinate{Col: x, Row: y}
	if val, ok := cache[coord]; ok {
		return val
	}

	var geologicIndex int
	if x == 0 && y == 0 || x == targetX && y == targetY {
		geologicIndex = 0
	} else if y == 0 {
		geologicIndex = x * 16807
	} else if x == 0 {
		geologicIndex = y * 48271
	} else {
		erosionLeft := getErosionLevel(x-1, y, depth, targetX, targetY, cache)
		erosionUp := getErosionLevel(x, y-1, depth, targetX, targetY, cache)
		geologicIndex = erosionLeft * erosionUp
	}

	erosionLevel := (geologicIndex + depth) % 20183
	cache[coord] = erosionLevel
	return erosionLevel
}

type state struct {
	minutes int
	x, y    int
	cannot  int
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	var depth, targetX, targetY int
	conv.MustSscanf(lines[0], "depth: %d", &depth)
	conv.MustSscanf(lines[1], "target: %d,%d", &targetX, &targetY)
	cache := make(map[gridutil.Coordinate]int)
	pq := container.NewPriorityQueue[state]()
	pq.Push(state{minutes: 0, x: 0, y: 0, cannot: 1}, 0)

	best := make(map[string]int)
	maxSearch := 1000

	for !pq.IsEmpty() {
		current := pq.Pop()

		key := fmt.Sprintf("%d,%d,%d", current.x, current.y, current.cannot)
		if minutes, exists := best[key]; exists && minutes <= current.minutes {
			continue
		}
		best[key] = current.minutes

		if current.x == targetX && current.y == targetY && current.cannot == 1 {
			fmt.Println("Part 2", current.minutes)
			return
		}

		regionType := getErosionLevel(current.x, current.y, depth, targetX, targetY, cache) % 3
		for i := range 3 {
			if i != current.cannot && i != regionType {
				pq.Push(state{
					minutes: current.minutes + 7,
					x:       current.x,
					y:       current.y,
					cannot:  i,
				}, current.minutes+7)
			}
		}

		for _, dir := range []struct{ dx, dy int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			newX := current.x + dir.dx
			newY := current.y + dir.dy

			if newX < 0 || newY < 0 || newX > targetX+maxSearch || newY > targetY+maxSearch {
				continue
			}

			newRegionType := getErosionLevel(newX, newY, depth, targetX, targetY, cache) % 3
			if newRegionType == current.cannot {
				continue
			}

			pq.Push(state{
				minutes: current.minutes + 1,
				x:       newX,
				y:       newY,
				cannot:  current.cannot,
			}, current.minutes+1)
		}
	}
}
