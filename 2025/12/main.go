package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

const shapeCount = 6

func main() {
	input, err := download.ReadInput(2025, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

func part1(input string) {
	densities, trees := parseInput(input)

	goodTrees := 0
	for _, line := range trees {
		if treeIsGood(line, densities) {
			goodTrees++
		}
	}

	fmt.Println("Part 1", goodTrees)
}

func parseInput(input string) ([]int, []string) {
	lines := conv.SplitNewline(input)
	densities := make([]int, 0, shapeCount)
	treeLines := make([]string, 0)

	i := 0

	for i < len(lines) && len(densities) < shapeCount {
		line := strings.TrimSpace(lines[i])
		i++

		if line == "" {
			continue
		}

		if !strings.HasSuffix(line, ":") {
			log.Fatalf("invalid shape header: %q", line)
		}

		density := 0
		rowsRead := 0
		for i < len(lines) && rowsRead < 3 {
			row := strings.TrimSpace(lines[i])
			i++
			if row == "" {
				continue
			}
			for _, ch := range row {
				if ch == '#' {
					density++
				}
			}
			rowsRead++
		}

		if rowsRead != 3 {
			log.Fatalf("incomplete shape block in input")
		}

		densities = append(densities, density)
	}

	if len(densities) != shapeCount {
		log.Fatalf("expected %d shape blocks, got %d", shapeCount, len(densities))
	}

	for ; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		treeLines = append(treeLines, line)
	}

	return densities, treeLines
}

func treeIsGood(rawLine string, densities []int) bool {
	left, right, found := strings.Cut(rawLine, ":")
	if !found {
		log.Fatalf("invalid tree line: %q", rawLine)
	}

	var width, height int
	if _, err := fmt.Sscanf(strings.TrimSpace(left), "%dx%d", &width, &height); err != nil {
		log.Fatalf("invalid tree dimensions in line %q: %v", rawLine, err)
	}

	amountTokens := strings.Fields(strings.TrimSpace(right))
	amounts := make([]int, 0, len(amountTokens))
	for _, token := range amountTokens {
		amounts = append(amounts, conv.MustAtoi(token))
	}

	looseWidth := width / 3
	looseHeight := height / 3
	capacityForLoosePresents := looseWidth * looseHeight

	loosePresents := 0
	for _, amount := range amounts {
		loosePresents += amount
	}
	if loosePresents <= capacityForLoosePresents {
		return true
	}

	availableArea := width * height
	tightPlacing := 0
	for i := 0; i < len(densities) && i < len(amounts); i++ {
		tightPlacing += amounts[i] * densities[i]
	}

	if availableArea < tightPlacing {
		return false
	}

	log.Fatalf("unexpected complicated tree line: %q", rawLine)
	return false
}
