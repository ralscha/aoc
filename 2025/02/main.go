package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/rangeutil"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2025, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	ranges := parseRanges(input)

	sum := 0
	for _, r := range ranges {
		for n := r.Start; n <= r.End; n++ {
			if isDoubledNumber(n) {
				sum += n
			}
		}
	}

	fmt.Println("Part 1", sum)
}

func part2(input string) {
	ranges := parseRanges(input)

	sum := 0
	for _, r := range ranges {
		for n := r.Start; n <= r.End; n++ {
			if isRepeatedNumber(n) {
				sum += n
			}
		}
	}
	fmt.Println("Part 2", sum)
}

func isDoubledNumber(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	return s[:half] == s[half:]
}

func isRepeatedNumber(n int) bool {
	s := strconv.Itoa(n)
	for size := 1; size <= len(s)/2; size++ {
		if len(s)%size != 0 {
			continue
		}
		repeated := true
		for i := size; i < len(s); i += size {
			if s[i:i+size] != s[0:size] {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}

func parseRanges(input string) []rangeutil.Range {
	input = strings.TrimSpace(input)
	var ranges []rangeutil.Range
	for part := range strings.SplitSeq(input, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		before, after, _ := strings.Cut(part, "-")
		start := conv.MustAtoi(before)
		end := conv.MustAtoi(after)
		ranges = append(ranges, rangeutil.NewRange(start, end))
	}
	return ranges
}
