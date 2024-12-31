package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	input, err := download.ReadInput(2020, 16)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type rule struct {
	name   string
	ranges []valueRange
}

type valueRange struct {
	min int
	max int
}

func parseInput(input string) ([]rule, []int, [][]int) {
	parts := strings.Split(input, "\n\n")

	ruleLines := conv.SplitNewline(parts[0])
	rules := make([]rule, len(ruleLines))
	ruleRegex := regexp.MustCompile(`^([^:]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	for i, line := range ruleLines {
		matches := ruleRegex.FindStringSubmatch(line)
		min1 := conv.MustAtoi(matches[2])
		max1 := conv.MustAtoi(matches[3])
		min2 := conv.MustAtoi(matches[4])
		max2 := conv.MustAtoi(matches[5])
		rules[i] = rule{
			name: matches[1],
			ranges: []valueRange{
				{min: min1, max: max1},
				{min: min2, max: max2},
			},
		}
	}

	yourTicketStr := conv.SplitNewline(parts[1])[1]
	yourTicket := make([]int, 0)
	for _, s := range strings.Split(yourTicketStr, ",") {
		yourTicket = append(yourTicket, conv.MustAtoi(s))
	}

	nearbyTicketsStr := conv.SplitNewline(parts[2])[1:]
	nearbyTickets := make([][]int, len(nearbyTicketsStr))
	for i, line := range nearbyTicketsStr {
		values := strings.Split(line, ",")
		ticket := make([]int, len(values))
		for j, v := range values {
			ticket[j] = conv.MustAtoi(v)
		}
		nearbyTickets[i] = ticket
	}

	return rules, yourTicket, nearbyTickets
}

func isValidValue(value int, rules []rule) bool {
	for _, r := range rules {
		for _, vr := range r.ranges {
			if value >= vr.min && value <= vr.max {
				return true
			}
		}
	}
	return false
}

func part1(input string) {
	rules, _, nearbyTickets := parseInput(input)

	errorRate := 0
	for _, ticket := range nearbyTickets {
		for _, value := range ticket {
			if !isValidValue(value, rules) {
				errorRate += value
			}
		}
	}

	fmt.Println("Part 1", errorRate)
}

func part2(input string) {
	rules, yourTicket, nearbyTickets := parseInput(input)

	validTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		isValid := true
		for _, value := range ticket {
			if !isValidValue(value, rules) {
				isValid = false
				break
			}
		}
		if isValid {
			validTickets = append(validTickets, ticket)
		}
	}

	numFields := len(yourTicket)
	possibleFields := make([][]string, numFields)
	for i := range numFields {
		possibleFields[i] = make([]string, 0)
		for _, rule := range rules {
			isValidForThisRule := true
			for _, ticket := range validTickets {
				value := ticket[i]
				isRuleValid := false
				for _, vr := range rule.ranges {
					if value >= vr.min && value <= vr.max {
						isRuleValid = true
						break
					}
				}
				if !isRuleValid {
					isValidForThisRule = false
					break
				}
			}
			if isValidForThisRule {
				possibleFields[i] = append(possibleFields[i], rule.name)
			}
		}
	}

	fieldOrder := make(map[int]string)
	assigned := container.NewSet[string]()

	for range numFields {
		for j := range numFields {
			if _, ok := fieldOrder[j]; ok {
				continue
			}
			if len(possibleFields[j]) == 1 {
				field := possibleFields[j][0]
				if !assigned.Contains(field) {
					fieldOrder[j] = field
					assigned.Add(field)
				}
			}
		}

		for j := range numFields {
			if _, ok := fieldOrder[j]; ok {
				continue
			}
			newPossible := make([]string, 0)
			for _, field := range possibleFields[j] {
				if !assigned.Contains(field) {
					newPossible = append(newPossible, field)
				}
			}
			possibleFields[j] = newPossible
		}
	}

	departureProduct := 1
	for i := range numFields {
		if strings.HasPrefix(fieldOrder[i], "departure") {
			departureProduct *= yourTicket[i]
		}
	}

	fmt.Println("Part 2", departureProduct)
}
