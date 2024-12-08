package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 5)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type rule struct {
	before, after int
}

func isUpdateValid(update []int, rules []rule) bool {
	indexMap := make(map[int]int)
	for i, page := range update {
		indexMap[page] = i
	}

	for _, rule := range rules {
		beforeIndex, beforeExists := indexMap[rule.before]
		afterIndex, afterExists := indexMap[rule.after]
		if beforeExists && afterExists && beforeIndex > afterIndex {
			return false
		}
	}

	return true
}

func part1(input string) {
	rules, updates := parseInput(input)
	var totalMiddlePages int

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			middlePage := update[len(update)/2]
			totalMiddlePages += middlePage
		}
	}
	fmt.Println("Part 1", totalMiddlePages)
}

func parseInput(input string) ([]rule, [][]int) {
	sections := strings.Split(input, "\n\n")

	var rules []rule
	lines := conv.SplitNewline(sections[0])
	for _, line := range lines {
		parts := strings.Split(line, "|")
		before := conv.MustAtoi(parts[0])
		after := conv.MustAtoi(parts[1])
		rules = append(rules, rule{before: before, after: after})
	}

	var updates [][]int
	lines = conv.SplitNewline(sections[1])
	for _, line := range lines {
		updates = append(updates, conv.ToIntSlice(strings.Split(line, ",")))
	}
	return rules, updates
}

func violatesRule(first, second int, pagesBefore map[int]*container.Set[int]) bool {
	if set, exists := pagesBefore[first]; exists && set.Contains(second) {
		return true
	}
	return false
}

func reorderUpdate(update []int, rules []rule) []int {
	// which pages (values) must come before this page (key)
	pagesBefore := make(map[int]*container.Set[int])

	for _, r := range rules {
		if slices.Contains(update, r.before) && slices.Contains(update, r.after) {
			if _, exists := pagesBefore[r.after]; !exists {
				pagesBefore[r.after] = container.NewSet[int]()
			}
			pagesBefore[r.after].Add(r.before)
		}
	}

	reordered := make([]int, len(update))
	copy(reordered, update)

	changed := true
	for changed {
		changed = false
	outer:
		for i := range len(reordered) {
			for j := i + 1; j < len(reordered); j++ {
				if violatesRule(reordered[i], reordered[j], pagesBefore) {
					reordered[i], reordered[j] = reordered[j], reordered[i]
					changed = true
					break outer
				}
			}
		}
	}

	return reordered
}

func part2(input string) {
	rules, updates := parseInput(input)
	var totalMiddlePages int

	for _, update := range updates {
		if !isUpdateValid(update, rules) {
			reordered := reorderUpdate(update, rules)
			totalMiddlePages += reordered[len(reordered)/2]
		}
	}

	fmt.Println("Part 2", totalMiddlePages)
}
