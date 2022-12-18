package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/18/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
}

type coord struct {
	x, y, z int
}

func part1(input string) {
	var coordinates []coord
	lines := conv.SplitNewline(input)

	droplets := make(map[coord]struct{})
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x, y, z := conv.MustAtoi(splitted[0]), conv.MustAtoi(splitted[1]), conv.MustAtoi(splitted[2])
		coordinates = append(coordinates, coord{x, y, z})
		droplets[coord{x, y, z}] = struct{}{}
	}
	directions := []int{-1, 1}
	surfaceArea := 0
	for point := range droplets {
		for _, d := range directions {
			if _, ok := droplets[coord{point.x + d, point.y, point.z}]; !ok {
				surfaceArea++
			}
			if _, ok := droplets[coord{point.x, point.y + d, point.z}]; !ok {
				surfaceArea++
			}
			if _, ok := droplets[coord{point.x, point.y, point.z + d}]; !ok {
				surfaceArea++
			}
		}

	}
	fmt.Println(surfaceArea)
}
