package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"encoding/json"
	"fmt"
	"log"
	"slices"
)

func main() {
	input, err := download.ReadInput(2022, 13)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	var indicesInRightOrder []int

	index := 1
	for i := 0; i < len(lines); i += 3 {
		if parseAndCheck(lines[i], lines[i+1]) {
			indicesInRightOrder = append(indicesInRightOrder, index)
		}
		index++
	}

	sum := 0
	for _, o := range indicesInRightOrder {
		sum += o
	}
	fmt.Println("Part 2", sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var packages [][]any
	for _, line := range lines {
		if line == "" {
			continue
		}
		var t []any
		err := json.Unmarshal([]byte(line), &t)
		if err != nil {
			log.Fatalf("unmarshalling failed: %v", err)
		}
		packages = append(packages, t)
	}

	start := []any{[]any{2.0}}
	packages = append(packages, start)
	end := []any{[]any{6.0}}
	packages = append(packages, end)

	slices.SortFunc(packages, func(a, b []any) int {
		return compare(a, b)
	})

	startIndex := 0
	endIndex := 0
	for i, p := range packages {
		if compare(p, start) == 0 {
			startIndex = i + 1
		} else if compare(p, end) == 0 {
			endIndex = i + 1
		}
	}

	fmt.Println("Part 2", startIndex*endIndex)
}

func parseAndCheck(line1, line2 string) bool {
	var t1, t2 []any
	err := json.Unmarshal([]byte(line1), &t1)
	if err != nil {
		log.Fatalf("unmarshalling failed: %v", err)
		return false
	}
	err = json.Unmarshal([]byte(line2), &t2)
	if err != nil {
		log.Fatalf("unmarshalling failed: %v", err)
		return false
	}
	return compare(t1, t2) == -1 || compare(t1, t2) == 0
}

func compare(input1 any, input2 any) int {
	switch input1.(type) {
	case float64:
		switch input2.(type) {
		case float64:
			if input1.(float64) < input2.(float64) {
				return -1
			} else if input1.(float64) > input2.(float64) {
				return 1
			} else {
				return 0
			}
		case []any:
			return compare([]any{input1}, input2)
		}
	case []any:
		switch input2.(type) {
		case float64:
			return compare(input1, []any{input2})
		case []any:
			for i := 0; i < len(input1.([]any)); i++ {
				if i >= len(input2.([]any)) {
					return 1
				}
				c := compare(input1.([]any)[i], input2.([]any)[i])
				if c != 0 {
					return c
				}
			}
			if len(input1.([]any)) < len(input2.([]any)) {
				return -1
			}
		}
	}
	return 0
}
