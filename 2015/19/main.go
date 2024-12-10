package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 19)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type replacement struct {
	from string
	to   string
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	replacements, molecule := convertInput(lines)

	distinctMolecules := container.NewSet[string]()
	for _, replacement := range replacements {
		for i := 0; i < len(molecule); i++ {
			if strings.HasPrefix(molecule[i:], replacement.from) {
				distinctMolecules.Add(molecule[:i] + replacement.to + molecule[i+len(replacement.from):])
			}
		}
	}

	fmt.Println("Part 1", distinctMolecules.Len())
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	replacements, molecule := convertInput(lines)
	targetMolecule := "e"

	slices.SortFunc(replacements, func(a, b replacement) int {
		return len(a.to) - len(b.to)
	})

	steps := 0
	for molecule != targetMolecule {
		for _, replacement := range replacements {
			for i := 0; i < len(molecule); i++ {
				if strings.HasPrefix(molecule[i:], replacement.to) {
					molecule = molecule[:i] + replacement.from + molecule[i+len(replacement.to):]
					steps++
					break
				}
			}
		}
	}

	fmt.Println("Part 2", steps)
}

func convertInput(lines []string) ([]replacement, string) {
	replacements := make([]replacement, 0)
	molecule := ""
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, " => ") {
			splitted := strings.Split(line, " => ")
			from := splitted[0]
			to := splitted[1]
			replacements = append(replacements, replacement{from: from, to: to})
		} else {
			molecule = line
		}
	}
	return replacements, molecule
}
