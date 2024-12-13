package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	totalArrangements := 0
	for _, line := range lines {
		splitted := strings.Fields(line)
		field := splitted[0]
		recordedGroups := conv.ToIntSlice(strings.Split(splitted[1], ","))
		
		// Use Set to store unique combinations
		combinations := container.NewSet[string]()
		for _, comb := range generateCombinations(field) {
			combinations.Add(comb)
		}

		// Process each unique combination
		for _, comb := range combinations.Values() {
			groups := countGroups(comb)
			if slices.Equal(groups, recordedGroups) {
				totalArrangements++
			}
		}
	}
	fmt.Println(totalArrangements)
}

func countGroups(field string) []int {
	var counts []int
	currentGroupCount := 0

	for _, ch := range field {
		if ch == '#' {
			currentGroupCount++
		} else if currentGroupCount > 0 {
			counts = append(counts, currentGroupCount)
			currentGroupCount = 0
		}
	}

	if currentGroupCount > 0 {
		counts = append(counts, currentGroupCount)
	}

	return counts
}

func generateCombinations(pattern string) []string {
	if len(pattern) == 0 {
		return []string{""}
	}

	if pattern[0] == '?' {
		dotCombinations := generateCombinations("." + pattern[1:])
		hashCombinations := generateCombinations("#" + pattern[1:])
		return append(dotCombinations, hashCombinations...)
	}
	
	subCombinations := generateCombinations(pattern[1:])
	for i, comb := range subCombinations {
		subCombinations[i] = string(pattern[0]) + comb
	}
	return subCombinations
}

// Cache using container.Bag to store frequency of results
var cache = make(map[string]int)

func part2(input string) {
	lines := conv.SplitNewline(input)
	totalArrangements := 0

	for _, line := range lines {
		splitted := strings.Fields(line)
		field := splitted[0]
		recordedGroups := conv.ToIntSlice(strings.Split(splitted[1], ","))

		// Build expanded field and groups
		newField := strings.Join([]string{field, field, field, field, field}, "?")
		newRecordedGroup := []int{}
		for i := 0; i < 5; i++ {
			newRecordedGroup = append(newRecordedGroup, recordedGroups...)
		}

		clear(cache)
		arrangements := count(newField, newRecordedGroup, 0, 0, 0)
		totalArrangements += arrangements
	}
	fmt.Println(totalArrangements)
}

func count(field string, expectedGroups []int, currentPos int, currentGroupsSize int, currentGroupLength int) int {
	key := fmt.Sprintf("%d_%d_%d", currentPos, currentGroupsSize, currentGroupLength)
	if val, ok := cache[key]; ok {
		return val
	}

	if currentPos == len(field) {
		if currentGroupsSize == len(expectedGroups) && currentGroupLength == 0 {
			return 1
		} else if currentGroupsSize == len(expectedGroups)-1 && expectedGroups[currentGroupsSize] == currentGroupLength {
			return 1
		}
		return 0
	}

	result := 0
	for _, ch := range []byte{'.', '#'} {
		if field[currentPos] == ch || field[currentPos] == '?' {
			if ch == '.' && currentGroupLength == 0 {
				result += count(field, expectedGroups, currentPos+1, currentGroupsSize, 0)
			} else if ch == '.' && currentGroupLength > 0 && currentGroupsSize < len(expectedGroups) && expectedGroups[currentGroupsSize] == currentGroupLength {
				result += count(field, expectedGroups, currentPos+1, currentGroupsSize+1, 0)
			} else if ch == '#' {
				result += count(field, expectedGroups, currentPos+1, currentGroupsSize, currentGroupLength+1)
			}
		}
	}

	cache[key] = result
	return result
}
