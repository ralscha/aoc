package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2015/09/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 9)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	distances := make(map[string]map[string]int)

	for _, line := range lines {
		splitted := strings.Fields(line)
		from := splitted[0]
		to := splitted[2]
		weight := conv.MustAtoi(splitted[4])

		if _, ok := distances[from]; !ok {
			distances[from] = make(map[string]int)
		}
		if _, ok := distances[to]; !ok {
			distances[to] = make(map[string]int)
		}
		distances[from][to] = weight
		distances[to][from] = weight
	}

	fmt.Println("shortest:", shortestRoute(distances))
	fmt.Println("longest :", longestRoute(distances))
}

func shortestRoute(distances map[string]map[string]int) int {
	locations := make([]string, 0, len(distances))
	for location := range distances {
		locations = append(locations, location)
	}

	shortestDistance := math.MaxInt32
	for _, route := range mathx.StringPermutations(locations) {
		distance := 0
		for i := 0; i < len(route)-1; i++ {
			distance += distances[route[i]][route[i+1]]
		}
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}

	return shortestDistance
}

func longestRoute(distances map[string]map[string]int) int {
	locations := make([]string, 0, len(distances))
	for location := range distances {
		locations = append(locations, location)
	}

	longestDistance := math.MinInt32
	for _, route := range mathx.StringPermutations(locations) {
		distance := 0
		for i := 0; i < len(route)-1; i++ {
			distance += distances[route[i]][route[i+1]]
		}
		if distance > longestDistance {
			longestDistance = distance
		}
	}

	return longestDistance
}
