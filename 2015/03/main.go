package main

import (
	"aoc/internal/container"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2015, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	visitedHouses := container.NewSet[string]()
	pos := gridutil.Coordinate{0, 0}
	visitedHouses.Add(fmt.Sprintf("%d,%d", pos.Row, pos.Col))

	for _, dir := range input {
		switch dir {
		case '^':
			pos.Row--
		case 'v':
			pos.Row++
		case '>':
			pos.Col++
		case '<':
			pos.Col--
		}
		visitedHouses.Add(fmt.Sprintf("%d,%d", pos.Row, pos.Col))
	}

	fmt.Println("Number of houses visited:", visitedHouses.Len())
}

func part2(input string) {
	visitedHouses := container.NewSet[string]()
	santa := gridutil.Coordinate{0, 0}
	roboSanta := gridutil.Coordinate{0, 0}
	visitedHouses.Add(fmt.Sprintf("%d,%d", santa.Row, santa.Col))

	for i, dir := range input {
		pos := &santa
		if i%2 == 1 {
			pos = &roboSanta
		}

		switch dir {
		case '^':
			pos.Row--
		case 'v':
			pos.Row++
		case '>':
			pos.Col++
		case '<':
			pos.Col--
		}
		visitedHouses.Add(fmt.Sprintf("%d,%d", pos.Row, pos.Col))
	}

	fmt.Println("Number of houses visited:", visitedHouses.Len())
}
