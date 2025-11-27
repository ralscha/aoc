package main

import (
	"fmt"
	"log"
	"strings"

	"aoc/internal/conv"
	"aoc/internal/download"
	"slices"
)

func main() {
	input, err := download.ReadInput(2020, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type rule struct {
	char       string
	subRules   [][]int
	subRulesOr [][]int
}

func parseRules(rulesStr string) map[int]rule {
	rules := make(map[int]rule)
	for line := range strings.SplitSeq(rulesStr, "\n") {
		parts := strings.SplitN(line, ": ", 2)
		ruleID := conv.MustAtoi(parts[0])
		definition := parts[1]

		if strings.Contains(definition, "\"") {
			rules[ruleID] = rule{char: strings.Trim(definition, "\"")}
		} else if strings.Contains(definition, "|") {
			orParts := strings.Split(definition, " | ")
			var subRulesOr [][]int
			for _, orPart := range orParts {
				var subRules []int
				for _, numStr := range strings.Split(orPart, " ") {
					subRules = append(subRules, conv.MustAtoi(numStr))
				}
				subRulesOr = append(subRulesOr, subRules)
			}
			rules[ruleID] = rule{subRulesOr: subRulesOr}
		} else {
			var subRules []int
			for _, numStr := range strings.Split(definition, " ") {
				subRules = append(subRules, conv.MustAtoi(numStr))
			}
			rules[ruleID] = rule{subRules: [][]int{subRules}}
		}
	}
	return rules
}

func generateMatches(ruleID int, rules map[int]rule, memo map[int][]string) []string {
	if matches, ok := memo[ruleID]; ok {
		return matches
	}

	rule := rules[ruleID]
	var result []string

	if rule.char != "" {
		result = append(result, rule.char)
	} else {
		processSubRules := func(subRuleSet []int) {
			subResults := [][]string{{""}}
			for _, subRuleID := range subRuleSet {
				var nextSubResults []string
				subMatches := generateMatches(subRuleID, rules, memo)
				for _, prefix := range subResults[len(subResults)-1] {
					for _, match := range subMatches {
						nextSubResults = append(nextSubResults, prefix+match)
					}
				}
				subResults = append(subResults, nextSubResults)
			}
			result = append(result, subResults[len(subResults)-1]...)
		}

		if len(rule.subRules) > 0 {
			for _, subRuleSet := range rule.subRules {
				processSubRules(subRuleSet)
			}
		}
		if len(rule.subRulesOr) > 0 {
			for _, subRuleSet := range rule.subRulesOr {
				processSubRules(subRuleSet)
			}
		}
	}

	memo[ruleID] = result
	return result
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	messages := conv.SplitNewline(parts[1])

	count := 0
	memo := make(map[int][]string)
	possibleMatches := generateMatches(0, rules, memo)
	for _, message := range messages {
		if slices.Contains(possibleMatches, message) {
			count++
		}
	}

	fmt.Println("Part 1", count)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	rules := parseRules(parts[0])
	messages := conv.SplitNewline(parts[1])

	rules[8] = rule{subRulesOr: [][]int{{42}, {42, 8}}}
	rules[11] = rule{subRulesOr: [][]int{{42, 31}, {42, 11, 31}}}

	memo := make(map[int][]string)
	matches42 := generateMatches(42, rules, memo)
	matches31 := generateMatches(31, rules, memo)

	count := 0
	for _, message := range messages {
		if matchesRule0(message, matches42, matches31) {
			count++
		}
	}

	fmt.Println("Part 2", count)
}

func matchesRule0(message string, matches42, matches31 []string) bool {
	n := len(message)
	len42 := len(matches42[0])
	len31 := len(matches31[0])

	for i := 1; i*len42 < n; i++ {
		prefix := message[:i*len42]
		if !matchesRepeatedly(prefix, matches42) {
			continue
		}

		for j := 1; i*len42+j*len31 <= n; j++ {
			suffix := message[i*len42 : i*len42+j*len31]
			if matchesRepeatedly(suffix, matches31) && i > j {
				remaining := message[i*len42+j*len31:]
				if matchesRepeatedly(remaining, matches42) {
					return true
				}
			}
		}
	}

	return false
}

func matchesRepeatedly(s string, matches []string) bool {
	l := len(matches[0])
	for i := 0; i < len(s); i += l {
		if !slices.Contains(matches, s[i:i+l]) {
			return false
		}
	}
	return true
}
