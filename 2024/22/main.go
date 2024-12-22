package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"maps"
	"slices"
)

func main() {
	input, err := download.ReadInput(2024, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func generateNextSecret(secret int) int {
	result := secret * 64
	secret = secret ^ result
	secret = secret % 16777216

	result = secret / 32
	secret = secret ^ result
	secret = secret % 16777216

	result = secret * 2048
	secret = secret ^ result
	return secret % 16777216
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	nums := conv.ToIntSlice(lines)
	sum := 0
	for _, initial := range nums {
		secret := initial
		for range 2000 {
			secret = generateNextSecret(secret)
		}
		sum += secret
	}

	fmt.Println("Part 1", sum)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	nums := conv.ToIntSlice(lines)
	patternValues := make(map[[4]int]int)
	for _, initial := range nums {
		secret := initial
		lastPrice := secret % 10

		var pattern [][2]int

		for range 2000 {
			secret = generateNextSecret(secret)
			currentPrice := secret % 10
			pattern = append(pattern, [2]int{currentPrice - lastPrice, currentPrice})
			lastPrice = currentPrice
		}

		seen := container.NewSet[[4]int]()
		for i := 0; i <= len(pattern)-4; i++ {
			var p [4]int
			for j := range 4 {
				p[j] = pattern[i+j][0]
			}

			if !seen.Contains(p) {
				seen.Add(p)
				patternValues[p] += pattern[i+3][1]
			}
		}
	}

	maxBananas := slices.Max(slices.Collect(maps.Values(patternValues)))
	fmt.Println("Part 2", maxBananas)
}
