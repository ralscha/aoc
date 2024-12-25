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
	input, err := download.ReadInput(2020, 21)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func parseInput(input string) ([]string, map[string]*container.Set[string], map[string]int) {
	lines := conv.SplitNewline(input)
	possible := make(map[string]*container.Set[string])
	ingredientsCount := make(map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		ingredients := strings.Split(parts[0], " ")
		var allergens []string
		if len(parts) > 1 {
			allergensStr := strings.TrimSuffix(parts[1], ")")
			allergens = strings.Split(allergensStr, ", ")
		}

		for _, ingredient := range ingredients {
			ingredientsCount[ingredient]++
		}

		for _, allergen := range allergens {
			if _, ok := possible[allergen]; !ok {
				possible[allergen] = container.NewSet[string]()
				for _, ingredient := range ingredients {
					possible[allergen].Add(ingredient)
				}
			} else {
				currentSet := container.NewSet[string]()
				for _, ingredient := range ingredients {
					currentSet.Add(ingredient)
				}
				possible[allergen] = possible[allergen].Intersection(currentSet)
			}
		}
	}

	return lines, possible, ingredientsCount
}

func part1(input string) {
	_, possible, ingredientsCount := parseInput(input)

	allergenicIngredients := container.NewSet[string]()
	for _, ingredientsSet := range possible {
		for _, ingredient := range ingredientsSet.Values() {
			allergenicIngredients.Add(ingredient)
		}
	}

	safeCount := 0
	for ingredient, count := range ingredientsCount {
		if !allergenicIngredients.Contains(ingredient) {
			safeCount += count
		}
	}

	fmt.Println("Part 1", safeCount)
}

func part2(input string) {
	_, possible, _ := parseInput(input)

	resolved := make(map[string]string)
	for len(resolved) < len(possible) {
		for allergen, ingredients := range possible {
			if ingredients.Len() == 1 {
				var ingredient string
				for _, k := range ingredients.Values() {
					ingredient = k
				}
				resolved[allergen] = ingredient
				for a := range possible {
					possible[a].Remove(ingredient)
				}
			}
		}
	}

	var sortedAllergens []string
	for allergen := range resolved {
		sortedAllergens = append(sortedAllergens, allergen)
	}
	slices.Sort(sortedAllergens)

	var canonicalDangerousList []string
	for _, allergen := range sortedAllergens {
		canonicalDangerousList = append(canonicalDangerousList, resolved[allergen])
	}

	fmt.Println("Part 2", strings.Join(canonicalDangerousList, ","))
}
