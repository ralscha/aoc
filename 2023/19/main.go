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

	part1(input)
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

func part1(input string) {
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
