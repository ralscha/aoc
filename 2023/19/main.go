package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type condition struct {
	attribute byte
	value     int
	compare   byte
}

type rule struct {
	condition condition
	target    string
}

type workflow struct {
	name  string
	rules []rule
}

type part struct {
	x, m, a, s int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)

	workflows := make(map[string]workflow)
	var parts []part

	i := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		workflow := parseWorkflow(line)
		workflows[workflow.name] = workflow
		i++
	}

	for _, line := range lines[i+1:] {
		part := parsePart(line)
		parts = append(parts, part)
	}

	totalRating := 0
	for _, part := range parts {
		accepted, rating := processPart(part, workflows)
		if accepted {
			totalRating += rating
		}
	}

	fmt.Println(totalRating)

	total := acceptedCombination(map[byte]partRange{
		'x': {1, 4001},
		'm': {1, 4001},
		'a': {1, 4001},
		's': {1, 4001},
	}, "in", workflows)

	fmt.Println(total)
}

func parseRule(ruleStr string) rule {
	parts := strings.Split(ruleStr, ":")
	if len(parts) > 1 {
		attribute := parts[0][0]
		compare := parts[0][1]
		value := conv.MustAtoi(parts[0][2:])
		target := parts[1]
		return rule{condition: condition{
			attribute: attribute,
			value:     value,
			compare:   compare,
		}, target: target}
	} else {
		return rule{condition: condition{},
			target: parts[0]}
	}
}

func parseWorkflow(workflowStr string) workflow {
	parts := strings.Split(workflowStr, "{")
	name := parts[0]
	ruleStrs := strings.Split(parts[1][:len(parts[1])-1], ",")
	rules := make([]rule, len(ruleStrs))
	for i, ruleStr := range ruleStrs {
		rules[i] = parseRule(ruleStr)
	}
	return workflow{name: name, rules: rules}
}

func parsePart(partStr string) part {
	partStr = partStr[1 : len(partStr)-1]
	ratings := strings.Split(partStr, ",")
	x := conv.MustAtoi(strings.Split(ratings[0], "=")[1])
	m := conv.MustAtoi(strings.Split(ratings[1], "=")[1])
	a := conv.MustAtoi(strings.Split(ratings[2], "=")[1])
	s := conv.MustAtoi(strings.Split(ratings[3], "=")[1])
	return part{x: x, m: m, a: a, s: s}
}

func evaluateCondition(part part, condition condition) bool {
	if condition.compare == 0 {
		return true
	}
	if condition.compare == '>' {
		switch condition.attribute {
		case 'x':
			return part.x > condition.value
		case 'm':
			return part.m > condition.value
		case 'a':
			return part.a > condition.value
		case 's':
			return part.s > condition.value
		}
	} else if condition.compare == '<' {
		switch condition.attribute {
		case 'x':
			return part.x < condition.value
		case 'm':
			return part.m < condition.value
		case 'a':
			return part.a < condition.value
		case 's':
			return part.s < condition.value
		}
	}
	return false
}

func processPart(part part, workflows map[string]workflow) (bool, int) {
	currentWorkflow := workflows["in"]
	for {
		accepted := false
		for _, rule := range currentWorkflow.rules {
			if evaluateCondition(part, rule.condition) {
				if rule.target == "A" {
					accepted = true
				} else if rule.target == "R" {
					return false, 0
				} else {
					currentWorkflow = workflows[rule.target]
				}
				break
			}
		}
		if accepted {
			return true, part.x + part.m + part.a + part.s
		}
	}
}

type partRange struct {
	start, stop int
}

func acceptedCombination(ranges map[byte]partRange, currentWorkflow string, workflows map[string]workflow) int {
	if currentWorkflow == "A" {
		return (ranges['x'].stop - ranges['x'].start) *
			(ranges['m'].stop - ranges['m'].start) *
			(ranges['a'].stop - ranges['a'].start) *
			(ranges['s'].stop - ranges['s'].start)
	} else if currentWorkflow == "R" {
		return 0
	}

	rangesCopy := make(map[byte]partRange)
	for k, v := range ranges {
		rangesCopy[k] = v
	}

	workflow := workflows[currentWorkflow]

	total := 0

	for _, rule := range workflow.rules {
		if rule.condition.compare == 0 {
			total += acceptedCombination(rangesCopy, rule.target, workflows)
			break
		}

		var destRange, afterRange partRange
		if rule.condition.compare == '>' {
			destRange = partRange{rule.condition.value + 1, rangesCopy[rule.condition.attribute].stop}
			afterRange = partRange{rangesCopy[rule.condition.attribute].start, rule.condition.value + 1}
		} else {
			destRange = partRange{rangesCopy[rule.condition.attribute].start, rule.condition.value}
			afterRange = partRange{rule.condition.value, rangesCopy[rule.condition.attribute].stop}
		}

		rangesCopy[rule.condition.attribute] = destRange
		total += acceptedCombination(rangesCopy, rule.target, workflows)
		rangesCopy[rule.condition.attribute] = afterRange
	}
	return total
}
