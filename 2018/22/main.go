package main

import (
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

func part2(input string) {
	fmt.Println("Part 2", "not yet implemented")
}
