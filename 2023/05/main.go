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

type srange struct {
	start int
	end   int
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

	// optimizied solution for part 2
	minLocation = -1

	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedLen := seeds[i+1]
		r := srange{
			start: seedStart,
			end:   seedStart + seedLen - 1,
		}
		var currentRanges []srange
		currentRanges = append(currentRanges, r)

		for _, rules := range groupRules {
			appliedRanges := make([]srange, 0)
			for _, rule := range rules {
				app, unapp := applyRule(currentRanges, rule)
				appliedRanges = append(appliedRanges, app...)
				currentRanges = unapp
			}
			currentRanges = append(currentRanges, appliedRanges...)
		}

		if len(currentRanges) > 0 {
			for _, currentRange := range currentRanges {
				if minLocation == -1 || currentRange.start < minLocation {
					minLocation = currentRange.start
				}
			}
		}
	}

	fmt.Println(minLocation)

	// brute force solution for part 2
	minLocation = -1

	result := make(chan int)
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedLen := seeds[i+1]
		go func() {
			result <- runSeed(seedStart, seedLen, groupRules)
		}()
	}
	for i := 0; i < len(seeds); i += 2 {
		locations := <-result
		if minLocation == -1 || locations < minLocation {
			minLocation = locations
		}
	}

	fmt.Println(minLocation)
}

func applyRule(inputRanges []srange, rule rule) ([]srange, []srange) {
	appliedRanges := make([]srange, 0)
	unappliedRanges := make([]srange, 0)
	for _, inputRange := range inputRanges {
		if inputRange.end < rule.sourceStart || inputRange.start > rule.sourceEnd {
			unappliedRanges = append(unappliedRanges, inputRange)
		} else if inputRange.start >= rule.sourceStart && inputRange.end <= rule.sourceEnd {
			appliedRanges = append(appliedRanges, srange{
				start: rule.destStart + inputRange.start - rule.sourceStart,
				end:   rule.destStart + inputRange.end - rule.sourceStart,
			})
		} else if inputRange.start < rule.sourceStart && inputRange.end <= rule.sourceEnd {
			unappliedRanges = append(unappliedRanges, srange{
				start: inputRange.start,
				end:   rule.sourceStart - 1,
			})
			appliedRanges = append(appliedRanges, srange{
				start: rule.destStart,
				end:   rule.destStart + inputRange.end - rule.sourceStart,
			})
		} else if inputRange.start >= rule.sourceStart && inputRange.end > rule.sourceEnd {
			appliedRanges = append(appliedRanges, srange{
				start: rule.destStart + inputRange.start - rule.sourceStart,
				end:   rule.destStart + rule.sourceEnd - rule.sourceStart,
			})
			unappliedRanges = append(unappliedRanges, srange{
				start: rule.sourceEnd + 1,
				end:   inputRange.end,
			})
		} else if inputRange.start < rule.sourceStart && inputRange.end > rule.sourceEnd {
			unappliedRanges = append(unappliedRanges, srange{
				start: inputRange.start,
				end:   rule.sourceStart - 1,
			})
			appliedRanges = append(appliedRanges, srange{
				start: rule.destStart,
				end:   rule.destStart + rule.sourceEnd - rule.sourceStart,
			})
			unappliedRanges = append(unappliedRanges, srange{
				start: rule.sourceEnd + 1,
				end:   inputRange.end,
			})
		}
	}
	return appliedRanges, unappliedRanges
}

func runSeed(seedStart int, seedLen int, groupRules [][]rule) int {
	minLocation := -1
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
	return minLocation
}
