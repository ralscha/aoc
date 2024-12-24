package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	pots, rules := parseInput(input)
	minPot, maxPot := 0, len(pots)-1

	for range 20 {
		pots, minPot, maxPot = nextGeneration(pots, rules, minPot, maxPot)
	}

	sum := calculateSum(pots)
	fmt.Println("Part 1", sum)
}

func part2(input string) {
	pots, rules := parseInput(input)
	var lastPattern string
	var lastPos int

	for gen := range 1000 {
		minPot, maxPot := getMinMaxPots(pots)
		pots, minPot, maxPot = nextGeneration(pots, rules, minPot, maxPot)

		pattern, sum, firstPlant := getCurrentPattern(pots, minPot, maxPot)
		if pattern == lastPattern {
			remaining := 50000000000 - gen - 1
			shift := firstPlant - lastPos
			finalSum := sum + (shift * remaining * strings.Count(pattern, "#"))
			fmt.Println("Part 2", finalSum)
			return
		}

		lastPattern = pattern
		lastPos = firstPlant
	}
}

func parseInput(input string) (map[int]rune, map[string]rune) {
	lines := conv.SplitNewline(input)
	initialStateStr := strings.Split(lines[0], ": ")[1]
	rules := make(map[string]rune)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " => ")
		rules[parts[0]] = rune(parts[1][0])
	}

	pots := make(map[int]rune)
	for i, r := range initialStateStr {
		pots[i] = r
	}
	return pots, rules
}

func nextGeneration(pots map[int]rune, rules map[string]rune, minPot, maxPot int) (map[int]rune, int, int) {
	nextGenPots := make(map[int]rune)
	nextMinPot, nextMaxPot := minPot, maxPot

	for i := minPot - 2; i <= maxPot+2; i++ {
		pattern := getPattern(pots, i)
		if result, ok := rules[pattern]; ok && result == '#' {
			nextGenPots[i] = '#'
			if i < nextMinPot {
				nextMinPot = i
			}
			if i > nextMaxPot {
				nextMaxPot = i
			}
		}
	}
	return nextGenPots, nextMinPot, nextMaxPot
}

func getPattern(pots map[int]rune, pos int) string {
	var pattern strings.Builder
	for i := pos - 2; i <= pos+2; i++ {
		if pots[i] == '#' {
			pattern.WriteRune('#')
		} else {
			pattern.WriteRune('.')
		}
	}
	return pattern.String()
}

func calculateSum(pots map[int]rune) int {
	sum := 0
	for i, p := range pots {
		if p == '#' {
			sum += i
		}
	}
	return sum
}

func getMinMaxPots(pots map[int]rune) (int, int) {
	minPot, maxPot := math.MaxInt, math.MinInt
	for pos := range pots {
		if pos < minPot {
			minPot = pos
		}
		if pos > maxPot {
			maxPot = pos
		}
	}
	return minPot, maxPot
}

func getCurrentPattern(pots map[int]rune, minPot, maxPot int) (string, int, int) {
	var currentPattern strings.Builder
	sum := 0
	firstPlant := math.MaxInt
	for i := minPot - 2; i <= maxPot+2; i++ {
		if pots[i] == '#' {
			if i < firstPlant {
				firstPlant = i
			}
			currentPattern.WriteRune('#')
			sum += i
		} else {
			currentPattern.WriteRune('.')
		}
	}
	return currentPattern.String(), sum, firstPlant
}
