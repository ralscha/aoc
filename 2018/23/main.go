package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
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
	fmt.Println("Part 2", "not yet implemented")
}
