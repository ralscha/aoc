package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/15/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 15)
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

func part1(input string) {
	lines := conv.SplitNewline(input)
	ingredients := makeIngredients(lines)
	var ingredientsArray []string
	for name := range ingredients {
		ingredientsArray = append(ingredientsArray, name)
	}

	maxScore := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-i-j; k++ {
				l := 100 - i - j - k
				score := score(ingredients, []ingredientAmount{
					{ingredientsArray[0], i},
					{ingredientsArray[1], j},
					{ingredientsArray[2], k},
					{ingredientsArray[3], l},
				})
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}
	fmt.Println(maxScore)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	ingredients := makeIngredients(lines)
	var ingredientsArray []string
	for name := range ingredients {
		ingredientsArray = append(ingredientsArray, name)
	}

	maxScore := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100-i; j++ {
			for k := 0; k < 100-i-j; k++ {
				l := 100 - i - j - k
				amounts := []ingredientAmount{
					{ingredientsArray[0], i},
					{ingredientsArray[1], j},
					{ingredientsArray[2], k},
					{ingredientsArray[3], l},
				}
				score := score(ingredients, amounts)
				calories := calories(ingredients, amounts)
				if calories == 500 && score > maxScore {
					maxScore = score
				}
			}
		}
	}
	fmt.Println(maxScore)
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
