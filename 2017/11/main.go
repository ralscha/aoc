package main

import (
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	input, err := download.ReadInput(2017, 11)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	moves := strings.Split(strings.TrimSpace(input), ",")
	x, y, z := 0, 0, 0

	for _, move := range moves {
		switch move {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		}
	}

	distance := slices.Max([]int{mathx.Abs(x), mathx.Abs(y), mathx.Abs(z)})
	fmt.Println("Part 1", distance)
}

func part2(input string) {
	moves := strings.Split(strings.TrimSpace(input), ",")
	x, y, z := 0, 0, 0
	maxDistance := 0

	for _, move := range moves {
		switch move {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		}
		distance := slices.Max([]int{mathx.Abs(x), mathx.Abs(y), mathx.Abs(z)})
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println("Part 2", maxDistance)
}
