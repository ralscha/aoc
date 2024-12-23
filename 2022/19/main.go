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
	input, err := download.ReadInput(2022, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type cost struct {
	ore      int
	clay     int
	obsidian int
}

type blueprint struct {
	id                int
	costOreRobot      cost
	costClayRobot     cost
	costObsidianRobot cost
	costGeodeRobot    cost
	maxOre            int
	maxClay           int
	maxObsidian       int
}

type state struct {
	minute        int
	ore           int
	clay          int
	obsidian      int
	geode         int
	oreRobot      int
	clayRobot     int
	obsidianRobot int
	geodeRobot    int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	blueprints := createBlueprints(lines)
	sum := 0
	for ix, blueprint := range blueprints {
		maxGeode := findMaxGeodes(blueprint, 24)
		sum += (ix + 1) * maxGeode
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	blueprints := createBlueprints(lines)
	result := 1
	for i := range min(3, len(blueprints)) {
		maxGeode := findMaxGeodes(blueprints[i], 32)
		result *= maxGeode
	}
	fmt.Println(result)
}

func createBlueprints(lines []string) []blueprint {
	var blueprints []blueprint
	for _, line := range lines {
		splitted := strings.Fields(line)
		id := conv.MustAtoi(splitted[1][:len(splitted[1])-1])
		costOreRobot := cost{ore: conv.MustAtoi(splitted[6])}
		costClayRobot := cost{ore: conv.MustAtoi(splitted[12])}
		costObsidianRobot := cost{ore: conv.MustAtoi(splitted[18]), clay: conv.MustAtoi(splitted[21])}
		costGeoRobot := cost{ore: conv.MustAtoi(splitted[27]), obsidian: conv.MustAtoi(splitted[30])}
		maxOre := max(costOreRobot.ore, costClayRobot.ore, costObsidianRobot.ore, costGeoRobot.ore)
		maxClay := max(costOreRobot.clay, costClayRobot.clay, costObsidianRobot.clay, costGeoRobot.clay)
		maxObsidian := max(costOreRobot.obsidian, costClayRobot.obsidian, costObsidianRobot.obsidian, costGeoRobot.obsidian)
		blueprints = append(blueprints, blueprint{id: id, costOreRobot: costOreRobot, costClayRobot: costClayRobot,
			costObsidianRobot: costObsidianRobot, costGeodeRobot: costGeoRobot,
			maxOre: maxOre, maxClay: maxClay, maxObsidian: maxObsidian})
	}
	return blueprints
}

func findMaxGeodes(bp blueprint, totalMinutes int) int {
	seen := container.NewSet[state]()
	maxGeodes := 0

	var dfs func(s state)
	dfs = func(s state) {
		// Update max geodes if we found a better solution
		if s.geode > maxGeodes {
			maxGeodes = s.geode
		}

		// If we've reached the time limit or seen this state, return
		if s.minute == totalMinutes || seen.Contains(s) {
			return
		}

		// Calculate theoretical max geodes possible from this point
		// Current geodes + current production + theoretical max new robots we could build
		remainingTime := totalMinutes - s.minute
		theoreticalMax := s.geode + (s.geodeRobot * remainingTime) +
			((remainingTime * (remainingTime - 1)) / 2)

		// If even the theoretical maximum can't beat our current best, prune this branch
		if theoreticalMax <= maxGeodes {
			return
		}

		seen.Add(s)

		// Try building each type of robot, but only if we need more of that resource
		nextMin := s.minute + 1

		// Always try to build geode robot first if possible
		if s.ore >= bp.costGeodeRobot.ore && s.obsidian >= bp.costGeodeRobot.obsidian {
			dfs(state{
				minute:        nextMin,
				ore:           s.ore + s.oreRobot - bp.costGeodeRobot.ore,
				clay:          s.clay + s.clayRobot,
				obsidian:      s.obsidian + s.obsidianRobot - bp.costGeodeRobot.obsidian,
				geode:         s.geode + s.geodeRobot,
				oreRobot:      s.oreRobot,
				clayRobot:     s.clayRobot,
				obsidianRobot: s.obsidianRobot,
				geodeRobot:    s.geodeRobot + 1,
			})
			return // If we can build a geode robot, we should always do it
		}

		// Try building obsidian robot
		if s.obsidianRobot < bp.maxObsidian &&
			s.ore >= bp.costObsidianRobot.ore && s.clay >= bp.costObsidianRobot.clay {
			dfs(state{
				minute:        nextMin,
				ore:           s.ore + s.oreRobot - bp.costObsidianRobot.ore,
				clay:          s.clay + s.clayRobot - bp.costObsidianRobot.clay,
				obsidian:      s.obsidian + s.obsidianRobot,
				geode:         s.geode + s.geodeRobot,
				oreRobot:      s.oreRobot,
				clayRobot:     s.clayRobot,
				obsidianRobot: s.obsidianRobot + 1,
				geodeRobot:    s.geodeRobot,
			})
		}

		// Try building clay robot
		if s.clayRobot < bp.maxClay && s.ore >= bp.costClayRobot.ore {
			dfs(state{
				minute:        nextMin,
				ore:           s.ore + s.oreRobot - bp.costClayRobot.ore,
				clay:          s.clay + s.clayRobot,
				obsidian:      s.obsidian + s.obsidianRobot,
				geode:         s.geode + s.geodeRobot,
				oreRobot:      s.oreRobot,
				clayRobot:     s.clayRobot + 1,
				obsidianRobot: s.obsidianRobot,
				geodeRobot:    s.geodeRobot,
			})
		}

		// Try building ore robot
		if s.oreRobot < bp.maxOre && s.ore >= bp.costOreRobot.ore {
			dfs(state{
				minute:        nextMin,
				ore:           s.ore + s.oreRobot - bp.costOreRobot.ore,
				clay:          s.clay + s.clayRobot,
				obsidian:      s.obsidian + s.obsidianRobot,
				geode:         s.geode + s.geodeRobot,
				oreRobot:      s.oreRobot + 1,
				clayRobot:     s.clayRobot,
				obsidianRobot: s.obsidianRobot,
				geodeRobot:    s.geodeRobot,
			})
		}

		// Try doing nothing (but only if we're not at resource caps)
		if s.ore < (bp.maxOre * 2) {
			dfs(state{
				minute:        nextMin,
				ore:           s.ore + s.oreRobot,
				clay:          s.clay + s.clayRobot,
				obsidian:      s.obsidian + s.obsidianRobot,
				geode:         s.geode + s.geodeRobot,
				oreRobot:      s.oreRobot,
				clayRobot:     s.clayRobot,
				obsidianRobot: s.obsidianRobot,
				geodeRobot:    s.geodeRobot,
			})
		}
	}

	// Start DFS with initial state
	dfs(state{oreRobot: 1})
	return maxGeodes
}
