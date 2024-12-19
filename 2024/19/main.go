package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"strings"
)

func main() {
	input, err := download.ReadInput(2024, 19)
	if err != nil {
		panic(err)
	}

	part1(input)
	part2(input)
}

func isDesignPossible(design string, patterns []string) bool {
	dp := make([]bool, len(design)+1)
	dp[0] = true

	for i := 1; i <= len(design); i++ {
		for _, pattern := range patterns {
			if i >= len(pattern) && dp[i-len(pattern)] {
				matches := true
				for j := 0; j < len(pattern); j++ {
					if design[i-len(pattern)+j] != pattern[j] {
						matches = false
						break
					}
				}
				if matches {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(design)]
}

func countDesignWays(design string, patterns []string) int {
	dp := make([]int, len(design)+1)
	dp[0] = 1

	for i := 1; i <= len(design); i++ {
		for _, pattern := range patterns {
			if i >= len(pattern) {
				matches := true
				for j := 0; j < len(pattern); j++ {
					if design[i-len(pattern)+j] != pattern[j] {
						matches = false
						break
					}
				}
				if matches {
					dp[i] += dp[i-len(pattern)]
				}
			}
		}
	}

	return dp[len(design)]
}

func part1(input string) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		fmt.Println("Invalid input")
		return
	}

	patterns := strings.Split(parts[0], ", ")
	designs := conv.SplitNewline(parts[1])
	possibleCount := 0
	for _, design := range designs {
		if isDesignPossible(design, patterns) {
			possibleCount++
		}
	}

	fmt.Println("Part 1", possibleCount)
}

func part2(input string) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		fmt.Println("Invalid input")
		return
	}

	patterns := strings.Split(parts[0], ", ")
	designs := conv.SplitNewline(parts[1])

	totalWays := 0
	for _, design := range designs {
		ways := countDesignWays(design, patterns)
		totalWays += ways
	}

	fmt.Println("Part 2", totalWays)
}
