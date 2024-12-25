package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	input, err := download.ReadInput(2019, 14)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type reaction struct {
	inputs  map[string]int
	output  string
	outputQ int
}

func parseReactions(input string) map[string]reaction {
	reactions := make(map[string]reaction)
	lines := conv.SplitNewline(input)
	re := regexp.MustCompile(`(\d+) ([A-Z]+)`)
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		inputsStr := parts[0]
		outputStr := parts[1]

		inputs := make(map[string]int)
		matches := re.FindAllStringSubmatch(inputsStr, -1)
		for _, match := range matches {
			quantity := conv.MustAtoi(match[1])
			chemical := match[2]
			inputs[chemical] = quantity
		}

		outputMatches := re.FindStringSubmatch(outputStr)
		outputQ := conv.MustAtoi(outputMatches[1])
		output := outputMatches[2]

		reactions[output] = reaction{inputs: inputs, output: output, outputQ: outputQ}
	}
	return reactions
}

func oreNeeded(fuelAmount int, reactions map[string]reaction) int {
	needed := map[string]int{"FUEL": fuelAmount}
	ore := 0
	surplus := make(map[string]int)

	for len(needed) > 0 {
		for chem, amount := range needed {
			if chem == "ORE" {
				ore += amount
				delete(needed, chem)
				continue
			}

			if surplus[chem] > 0 {
				if surplus[chem] >= amount {
					surplus[chem] -= amount
					delete(needed, chem)
					continue
				} else {
					amount -= surplus[chem]
					surplus[chem] = 0
				}
			}

			reaction := reactions[chem]
			factor := (amount + reaction.outputQ - 1) / reaction.outputQ

			for inputChem, inputAmount := range reaction.inputs {
				needed[inputChem] += inputAmount * factor
			}

			surplus[chem] += reaction.outputQ*factor - amount
			delete(needed, chem)
		}
	}

	return ore
}

func part1(input string) {
	reactions := parseReactions(input)
	fmt.Println("Part 1", oreNeeded(1, reactions))
}

func part2(input string) {
	reactions := parseReactions(input)
	maxOre := 1000000000000
	low := 1
	high := maxOre
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		neededOre := oreNeeded(mid, reactions)
		if neededOre <= maxOre {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Part 2", ans)
}
