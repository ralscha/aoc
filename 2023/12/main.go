package main

import (
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
		combinations := generateCombinations(field)
		for _, comb := range combinations {
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
		} else {
			if currentGroupCount > 0 {
				counts = append(counts, currentGroupCount)
				currentGroupCount = 0
			}
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
	} else {
		subCombinations := generateCombinations(pattern[1:])
		for i, comb := range subCombinations {
			subCombinations[i] = string(pattern[0]) + comb
		}
		return subCombinations
	}
}

var cache = map[string]int{}

func part2(input string) {
	lines := conv.SplitNewline(input)
	totalArrangements := 0

	for _, line := range lines {
		splitted := strings.Fields(line)
		field := splitted[0]
		recordedGroups := conv.ToIntSlice(strings.Split(splitted[1], ","))

		newField := field
		newRecordedGroup := recordedGroups
		for i := 0; i < 4; i++ {
			newRecordedGroup = append(newRecordedGroup, recordedGroups...)
			newField += "?" + field
		}

		clear(cache)
		arrangments := count(newField, newRecordedGroup, 0, 0, 0)
		totalArrangements += arrangments

	}
	fmt.Println(totalArrangements)
}

func count(field string, expectedGroups []int, currentPos int, currentPosBlocks int, currentHashLength int) int {
	key := fmt.Sprintf("%d_%d_%d", currentPos, currentPosBlocks, currentHashLength)
	if val, ok := cache[key]; ok {
		return val
	}

	if currentPos == len(field) {
		if currentPosBlocks == len(expectedGroups) && currentHashLength == 0 {
			return 1
		} else if currentPosBlocks == len(expectedGroups)-1 && expectedGroups[currentPosBlocks] == currentHashLength {
			return 1
		} else {
			return 0
		}
	}

	groupCount := 0
	for _, ch := range []byte{'.', '#'} {
		if field[currentPos] == ch || field[currentPos] == '?' {
			if ch == '.' && currentHashLength == 0 {
				groupCount += count(field, expectedGroups, currentPos+1, currentPosBlocks, 0)
			} else if ch == '.' && currentHashLength > 0 && currentPosBlocks < len(expectedGroups) && expectedGroups[currentPosBlocks] == currentHashLength {
				groupCount += count(field, expectedGroups, currentPos+1, currentPosBlocks+1, 0)
			} else if ch == '#' {
				groupCount += count(field, expectedGroups, currentPos+1, currentPosBlocks, currentHashLength+1)
			}
		}
	}

	cache[key] = groupCount
	return groupCount
}
