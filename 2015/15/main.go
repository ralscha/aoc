package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type ingredientAmount struct {
	name   string
	amount int
}

func generateCombinations(ingredientNames []string) [][]ingredientAmount {
	var combinations [][]ingredientAmount
	for i := range 100 {
		for j := range 100 - i {
			for k := range 100 - i - j {
				l := 100 - i - j - k
				combinations = append(combinations, []ingredientAmount{
					{ingredientNames[0], i},
					{ingredientNames[1], j},
					{ingredientNames[2], k},
					{ingredientNames[3], l},
				})
			}
		}
	}
	return combinations
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	ingredients := makeIngredients(lines)

	ingredientNames := slices.Collect(maps.Keys(ingredients))
	combinations := generateCombinations(ingredientNames)

	maxCombo := slices.MaxFunc(combinations, func(a, b []ingredientAmount) int {
		scoreA := score(ingredients, a)
		scoreB := score(ingredients, b)
		return scoreA - scoreB
	})

	fmt.Println(score(ingredients, maxCombo))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	ingredients := makeIngredients(lines)

	ingredientNames := slices.Collect(maps.Keys(ingredients))
	combinations := generateCombinations(ingredientNames)

	validCombos := slices.DeleteFunc(combinations, func(combo []ingredientAmount) bool {
		return calories(ingredients, combo) != 500
	})

	if len(validCombos) > 0 {
		maxCombo := slices.MaxFunc(validCombos, func(a, b []ingredientAmount) int {
			scoreA := score(ingredients, a)
			scoreB := score(ingredients, b)
			return scoreA - scoreB
		})
		fmt.Println(score(ingredients, maxCombo))
	} else {
		fmt.Println(0)
	}
}

func score(ingredients map[string]ingredient, ia []ingredientAmount) int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0

	for _, i := range ia {
		if ing, ok := ingredients[i.name]; ok {
			capacity += i.amount * ing.capacity
			durability += i.amount * ing.durability
			flavor += i.amount * ing.flavor
			texture += i.amount * ing.texture
		}
	}

	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0
	}
	return capacity * durability * flavor * texture
}

func calories(ingredients map[string]ingredient, ia []ingredientAmount) int {
	calories := 0
	for _, i := range ia {
		if ing, ok := ingredients[i.name]; ok {
			calories += i.amount * ing.calories
		}
	}
	return calories
}

func makeIngredients(lines []string) map[string]ingredient {
	ingredients := make(map[string]ingredient)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		name := parts[0][:len(parts[0])-1]
		capacity := conv.MustAtoi(parts[2][:len(parts[2])-1])
		durability := conv.MustAtoi(parts[4][:len(parts[4])-1])
		flavor := conv.MustAtoi(parts[6][:len(parts[6])-1])
		texture := conv.MustAtoi(parts[8][:len(parts[8])-1])
		calories := conv.MustAtoi(parts[10])
		ingredients[name] = ingredient{
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}
	}
	return ingredients
}
