package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	total := 0

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			aStr, bStr := match[1], match[2]
			a, errA := strconv.Atoi(aStr)
			b, errB := strconv.Atoi(bStr)
			if errA != nil || errB != nil {
				continue
			}
			total += a * b
		}
	}

	fmt.Println("Part 1", total)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	mulRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	instructionRegex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	mulEnabled := true
	total := 0

	for _, line := range lines {
		matches := instructionRegex.FindAllString(line, -1)
		for _, match := range matches {
			if strings.HasPrefix(match, "don't") {
				mulEnabled = false
				continue
			}

			if strings.HasPrefix(match, "do") {
				mulEnabled = true
				continue
			}

			if mulMatch := mulRegex.FindStringSubmatch(match); mulMatch != nil && mulEnabled {
				a, errA := strconv.Atoi(mulMatch[1])
				b, errB := strconv.Atoi(mulMatch[2])
				if errA == nil && errB == nil {
					total += a * b
				}
			}
		}
	}

	fmt.Println("Part 2", total)
}
