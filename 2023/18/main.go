package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := download.ReadInput(2023, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type point struct {
	x, y int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	pos := point{0, 0}
	points := []point{pos}
	perimeter := 0
	for _, line := range lines {
		splitted := strings.Fields(line)
		directionString := splitted[0]
		distance := conv.MustAtoi(splitted[1])

		var direction point
		switch directionString {
		case "R":
			direction = point{1, 0}
		case "D":
			direction = point{0, -1}
		case "L":
			direction = point{-1, 0}
		case "U":
			direction = point{0, 1}
		}

		pos = point{pos.x + direction.x*distance, pos.y + direction.y*distance}
		perimeter += distance
		points = append(points, pos)
	}

	shoelace(points, perimeter)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	pos := point{0, 0}
	points := []point{pos}
	perimeter := 0
	for _, line := range lines {
		splitted := strings.Fields(line)
		color := splitted[2]
		color = color[1 : len(color)-1]

		distanceString := color[1:6]
		distance, err := strconv.ParseInt(distanceString, 16, 64)
		distanceInt := int(distance)
		if err != nil {
			log.Fatalf("parsing hex failed: %v", err)
		}
		directionString := color[6:]

		var direction point
		switch directionString {
		case "0":
			direction = point{1, 0}
		case "1":
			direction = point{0, -1}
		case "2":
			direction = point{-1, 0}
		case "3":
			direction = point{0, 1}
		}

		pos = point{pos.x + direction.x*distanceInt, pos.y + direction.y*distanceInt}
		perimeter += distanceInt
		points = append(points, pos)
	}

	shoelace(points, perimeter)
}

func shoelace(points []point, perimeter int) {
	points = reverse(points)
	a := 0
	for i := range len(points) - 1 {
		a += (points[i].y + points[i+1].y) * (points[i].x - points[i+1].x)
	}
	fmt.Println(perimeter/2 + a/2 + 1)
}

func reverse(points []point) []point {
	var res []point
	for i := len(points) - 1; i >= 0; i-- {
		res = append(res, points[i])
	}
	return res
}
