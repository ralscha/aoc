package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
)

func main() {
	input, err := download.ReadInput(2016, 24)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func solve(input string, part2 bool) int {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewCharGrid2D(lines)

	locations := make(map[int]gridutil.Coordinate)

	for r := 0; r < grid.Height(); r++ {
		for c := 0; c < grid.Width(); c++ {
			if val, ok := grid.Get(r, c); ok && '0' <= val && val <= '9' {
				num, _ := strconv.Atoi(string(val))
				locations[num] = gridutil.Coordinate{Row: r, Col: c}
			}
		}
	}

	allPairsDistances := make(map[[2]int]int)

	for i := 0; i < len(locations); i++ {
		for j := i + 1; j < len(locations); j++ {
			start, ok1 := locations[i]
			end, ok2 := locations[j]
			if !ok1 || !ok2 {
				continue
			}

			isGoal := func(currentPos gridutil.Coordinate, current rune) bool {
				return currentPos == end
			}
			isObstacle := func(currentPos gridutil.Coordinate, current rune) bool {
				return current == '#'
			}
			directions := gridutil.Get4Directions()

			pathResult, found := grid.ShortestPathWithBFS(start, isGoal, isObstacle, directions)
			if found {
				dist := len(pathResult.Path) - 1
				allPairsDistances[[2]int{i, j}] = dist
				allPairsDistances[[2]int{j, i}] = dist
			} else {
				return math.MaxInt32
			}
		}
	}

	numLocations := len(locations)
	targets := make([]int, 0)
	for i := 1; i < numLocations; i++ {
		targets = append(targets, i)
	}

	minSteps := math.MaxInt32

	slices.Sort(targets)
	permutations := mathx.Permutations(targets)

	startLocation := 0

	for _, perm := range permutations {
		currentSteps := 0
		currentLocation := startLocation

		for _, nextLocation := range perm {
			if dist, ok := allPairsDistances[[2]int{currentLocation, nextLocation}]; ok {
				currentSteps += dist
				currentLocation = nextLocation
			} else {
				panic("distance not found")
			}
		}

		if part2 {
			if dist, ok := allPairsDistances[[2]int{currentLocation, startLocation}]; ok {
				currentSteps += dist
			} else {
				panic("distance back to start not found")
			}
		}

		minSteps = min(minSteps, currentSteps)
	}

	return minSteps
}

func part1(input string) {
	result := solve(input, false)
	fmt.Println("Part 1", result)
}

func part2(input string) {
	result := solve(input, true)
	fmt.Println("Part 2", result)
}
