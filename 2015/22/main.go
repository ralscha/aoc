package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"log"
)

func main() {
	inputFile := "./2015/22/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	//part2(input)
}

type spell struct {
	name     string
	cost     int
	duration int
	damage   int
	heal     int
	armor    int
	mana     int
}

type player struct {
	hp, mana, armor int
}

type boss struct {
	hp, damage int
	effects    []spell
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	/*
		spells := []spell{
			{53, 0, 4, 0, 0, 0},
			{73, 0, 2, 2, 0, 0},
			{113, 6, 0, 0, 7, 0},
			{173, 6, 3, 0, 0, 0},
			{229, 5, 0, 0, 0, 101},
		}

		bossHitPoints := conv.MustAtoi(strings.Field(lines[0])[2])
		bossDamage := conv.MustAtoi(strings.Field(lines[1])[1])

		boss := boss{bossHitPoints, bossDamage, []spell{}}
		player := player{
			hp:    50,
			mana:  500,
			armor: 0,
		}
	*/
}
