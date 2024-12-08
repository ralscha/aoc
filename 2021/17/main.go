package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

type targetArea struct {
	xMin, xMax, yMin, yMax int
}

func main() {
	input, err := download.ReadInput(2021, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	target := parseInput(input)
	validVelocities := findValidVelocities(target)

	// Find highest Y position
	highestY := 0
	for _, vel := range validVelocities.Values() {
		_, maxY := simulate(vel, target)
		if maxY > highestY {
			highestY = maxY
		}
	}

	fmt.Println("Part 1", highestY)
	fmt.Println("Part 2", validVelocities.Len())
}

func parseInput(input string) targetArea {
	lines := conv.SplitNewline(input)
	parts := strings.Fields(lines[0])

	// Parse X range
	xStr := strings.TrimPrefix(parts[2], "x=")
	xStr = strings.TrimSuffix(xStr, ",")
	xRange := strings.Split(xStr, "..")
	xMin := conv.MustAtoi(xRange[0])
	xMax := conv.MustAtoi(xRange[1])

	// Parse Y range
	yStr := strings.TrimPrefix(parts[3], "y=")
	yRange := strings.Split(yStr, "..")
	yMin := conv.MustAtoi(yRange[0])
	yMax := conv.MustAtoi(yRange[1])

	return targetArea{xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax}
}

func findValidVelocities(target targetArea) *container.Set[gridutil.Direction] {
	validVelocities := container.NewSet[gridutil.Direction]()

	// We can determine reasonable bounds for velocity search:
	// - X velocity must be positive and not overshoot target in first step
	// - Y velocity must not undershoot target after first step down
	for dx := 0; dx <= target.xMax; dx++ {
		for dy := target.yMin; dy <= -target.yMin; dy++ {
			vel := gridutil.Direction{Row: dy, Col: dx}
			if hit, _ := simulate(vel, target); hit {
				validVelocities.Add(vel)
			}
		}
	}

	return validVelocities
}

func simulate(velocity gridutil.Direction, target targetArea) (bool, int) {
	pos := gridutil.Coordinate{Row: 0, Col: 0}
	vel := velocity
	maxY := pos.Row

	for pos.Col <= target.xMax && pos.Row >= target.yMin {
		// Update position
		pos.Col += vel.Col
		pos.Row += vel.Row

		// Track maximum height
		if pos.Row > maxY {
			maxY = pos.Row
		}

		// Update velocity
		if vel.Col > 0 {
			vel.Col--
		} else if vel.Col < 0 {
			vel.Col++
		}
		vel.Row--

		// Check if we hit the target
		if pos.Col >= target.xMin && pos.Col <= target.xMax &&
			pos.Row >= target.yMin && pos.Row <= target.yMax {
			return true, maxY
		}
	}

	return false, maxY
}
