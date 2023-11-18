package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/19/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 19)
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

type production struct {
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
	minutes := 24
	startProd := production{oreRobot: 1}
	sum := 0
	for ix, blueprint := range blueprints {
		maxGeode := simulate(blueprint, minutes, startProd)
		sum += (ix + 1) * maxGeode
	}
	fmt.Println(sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	blueprints := createBlueprints(lines)
	minutes := 32
	startProd := production{oreRobot: 1}
	result := 1
	for i := 0; i < 3; i++ {
		maxGeode := simulate(blueprints[i], minutes, startProd)
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

func simulate(blueprint blueprint, minutes int, startProd production) int {
	var nextGeneration []production
	nextGeneration = append(nextGeneration, startProd)

	for m := 0; m < minutes; m++ {
		currentGeneration := nextGeneration
		nextGeneration = nil

		generated := make(map[production]bool)
		for _, prod := range currentGeneration {
			// building
			if prod.obsidian >= blueprint.costGeodeRobot.obsidian &&
				prod.ore >= blueprint.costGeodeRobot.ore {
				p := copyProduction(prod)
				p.obsidian -= blueprint.costGeodeRobot.obsidian
				p.ore -= blueprint.costGeodeRobot.ore
				p.ore += p.oreRobot
				p.clay += p.clayRobot
				p.obsidian += p.obsidianRobot
				p.geode += p.geodeRobot
				p.geodeRobot++
				if !generated[p] {
					nextGeneration = append(nextGeneration, p)
					generated[p] = true
				}
			} else {
				if prod.obsidianRobot < blueprint.maxObsidian {
					if prod.clay >= blueprint.costObsidianRobot.clay &&
						prod.ore >= blueprint.costObsidianRobot.ore {
						p := copyProduction(prod)
						p.clay -= blueprint.costObsidianRobot.clay
						p.ore -= blueprint.costObsidianRobot.ore
						p.ore += p.oreRobot
						p.clay += p.clayRobot
						p.obsidian += p.obsidianRobot
						p.geode += p.geodeRobot
						p.obsidianRobot++
						if !generated[p] {
							nextGeneration = append(nextGeneration, p)
							generated[p] = true
						}
					}
				}
				if prod.clayRobot < blueprint.maxClay {
					if prod.ore >= blueprint.costClayRobot.ore {
						p := copyProduction(prod)
						p.ore -= blueprint.costClayRobot.ore
						p.ore += p.oreRobot
						p.clay += p.clayRobot
						p.obsidian += p.obsidianRobot
						p.geode += p.geodeRobot
						p.clayRobot++
						if !generated[p] {
							nextGeneration = append(nextGeneration, p)
							generated[p] = true
						}
					}
				}
				if prod.oreRobot < blueprint.maxOre {
					if prod.ore >= blueprint.costOreRobot.ore {
						p := copyProduction(prod)
						p.ore -= blueprint.costOreRobot.ore
						p.ore += p.oreRobot
						p.clay += p.clayRobot
						p.obsidian += p.obsidianRobot
						p.geode += p.geodeRobot
						p.oreRobot++
						if !generated[p] {
							nextGeneration = append(nextGeneration, p)
							generated[p] = true
						}
					}
				}
				// do nothing
				prod.ore += prod.oreRobot
				prod.clay += prod.clayRobot
				prod.obsidian += prod.obsidianRobot
				prod.geode += prod.geodeRobot
				if !generated[prod] {
					nextGeneration = append(nextGeneration, prod)
					generated[prod] = true
				}
			}
		}
		maxGeode := 0
		maxGeodeRobot := 0
		for _, prod := range nextGeneration {
			if prod.geode > maxGeode {
				maxGeode = prod.geode
			}
			if prod.geodeRobot > maxGeodeRobot {
				maxGeodeRobot = prod.geodeRobot
			}
		}
		// remove all that have less than maxGeode-2 and less than maxGeodeRobot-2
		var newNextGeneration []production
		for _, prod := range nextGeneration {
			if prod.geode >= maxGeode-2 && prod.geodeRobot >= maxGeodeRobot-2 {
				newNextGeneration = append(newNextGeneration, prod)
			}
		}
		nextGeneration = newNextGeneration
	}

	// get max geode
	maxGeode := 0
	for _, prod := range nextGeneration {
		if prod.geode > maxGeode {
			maxGeode = prod.geode
		}
	}
	return maxGeode
}

func copyProduction(prod production) production {
	return production{
		ore:           prod.ore,
		clay:          prod.clay,
		obsidian:      prod.obsidian,
		geode:         prod.geode,
		oreRobot:      prod.oreRobot,
		clayRobot:     prod.clayRobot,
		obsidianRobot: prod.obsidianRobot,
		geodeRobot:    prod.geodeRobot,
	}
}
