package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
)

func main() {
	input, err := download.ReadInput(2020, 7)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parseRules(input string) map[string]map[string]int {
	rules := make(map[string]map[string]int)
	lines := conv.SplitNewline(input)
	reOuter := regexp.MustCompile(`^(.+?) bags contain`)
	reInner := regexp.MustCompile(`(\d+) (.+?) bag`)

	for _, line := range lines {
		matches := reOuter.FindStringSubmatch(line)
		if len(matches) != 2 {
			continue
		}
		outerBag := matches[1]
		innerBags := make(map[string]int)
		for _, match := range reInner.FindAllStringSubmatch(line, -1) {
			count := conv.MustAtoi(match[1])
			innerBag := match[2]
			innerBags[innerBag] = count
		}
		rules[outerBag] = innerBags
	}
	return rules
}

func part1(input string) {
	rules := parseRules(input)
	canContainShinyGold := container.NewSet[string]()
	canContainShinyGold.Add("shiny gold")

	added := true
	for added {
		added = false
		for outerBag, innerBags := range rules {
			if canContainShinyGold.Contains(outerBag) {
				continue
			}
			for innerBag := range innerBags {
				if canContainShinyGold.Contains(innerBag) {
					canContainShinyGold.Add(outerBag)
					added = true
					break
				}
			}
		}
	}

	fmt.Println("Part 1", canContainShinyGold.Len()-1)
}

func part2(input string) {
	rules := parseRules(input)

	var countBags func(bagColor string) int
	countBags = func(bagColor string) int {
		total := 0
		for innerBag, count := range rules[bagColor] {
			total += count * (1 + countBags(innerBag))
		}
		return total
	}

	fmt.Println("Part 2", countBags("shiny gold"))
}
