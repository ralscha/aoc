package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type rule struct {
	destStart   int
	sourceStart int
	sourceEnd   int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	seedsLine := strings.TrimPrefix(lines[0], "seeds: ")
	seedNumbers := strings.Split(seedsLine, " ")
	seeds := make([]int, len(seedNumbers))
	for i, seedStr := range seedNumbers {
		seeds[i] = conv.MustAtoi(seedStr)
	}

	groupRules := make([][]rule, 0)

	group := -1
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		if strings.Contains(line, "map:") {
			group++
			continue
		}
		parts := strings.Fields(line)
		destStart := conv.MustAtoi(parts[0])
		sourceStart := conv.MustAtoi(parts[1])
		length := conv.MustAtoi(parts[2])

		if len(groupRules) <= group {
			groupRules = append(groupRules, make([]rule, 0))
		}
		groupRules[group] = append(groupRules[group], rule{
			destStart:   destStart,
			sourceStart: sourceStart,
			sourceEnd:   sourceStart + length - 1,
		})
	}
	minLocation := -1
	for _, seed := range seeds {
		for _, rules := range groupRules {
			for _, rule := range rules {
				if seed >= rule.sourceStart && seed <= rule.sourceEnd {
					seed = rule.destStart + (seed - rule.sourceStart)
					break
				}
			}
		}

		if minLocation == -1 || seed < minLocation {
			minLocation = seed
		}
	}
	fmt.Println(minLocation)

	minLocation = -1

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedLen := seeds[i+1]
		for seed := seedStart; seed < seedStart+seedLen; seed++ {
			current := seed
			for _, rules := range groupRules {
				for _, rule := range rules {
					if current >= rule.sourceStart && current <= rule.sourceEnd {
						current = rule.destStart + (current - rule.sourceStart)
						break
					}
				}
			}

			if minLocation == -1 || current < minLocation {
				minLocation = current
			}
		}
	}

	fmt.Println(minLocation)
}
