package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
)

func main() {
	inputFile := "./2015/03/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)

}

func part1(input string) {
	visitedHouses := make(map[string]bool)
	x, y := 0, 0
	for _, dir := range input {
		switch dir {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		visitedHouses[fmt.Sprintf("%d,%d", x, y)] = true
	}

	fmt.Println("Number of houses visited:", len(visitedHouses))
}

func part2(input string) {
	visitedHouses := make(map[string]bool)
	x, y := 0, 0
	x2, y2 := 0, 0
	for i, dir := range input {
		if i%2 == 0 {
			switch dir {
			case '^':
				y++
			case 'v':
				y--
			case '>':
				x++
			case '<':
				x--
			}
			visitedHouses[fmt.Sprintf("%d,%d", x, y)] = true
		} else {
			switch dir {
			case '^':
				y2++
			case 'v':
				y2--
			case '>':
				x2++
			case '<':
				x2--
			}
			visitedHouses[fmt.Sprintf("%d,%d", x2, y2)] = true
		}
	}

	fmt.Println("Number of houses visited:", len(visitedHouses))
}
