package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 17)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	splitted := strings.Fields(lines[0])
	xStr := splitted[2]
	xStr = xStr[2 : len(xStr)-1]
	yStr := splitted[3]
	yStr = yStr[2:]

	xSplitted := strings.Split(xStr, "..")
	ySplitted := strings.Split(yStr, "..")
	xMin := conv.MustAtoi(xSplitted[0])
	xMax := conv.MustAtoi(xSplitted[1])
	yMin := conv.MustAtoi(ySplitted[0])
	yMax := conv.MustAtoi(ySplitted[1])

	target := area{xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax}
	highestY := findHighestY(target)
	fmt.Println(highestY)

	count := findAllVelocities(target)
	fmt.Println(count)
}

type point struct {
	x, y int
}

type velocity struct {
	dx, dy int
}

type area struct {
	xMin, xMax, yMin, yMax int
}

func simulate(v velocity, target area) (bool, int) {
	position := point{0, 0}
	maxY := position.y

	for position.x <= target.xMax && position.y >= target.yMin {
		position.x += v.dx
		position.y += v.dy

		if position.y > maxY {
			maxY = position.y
		}

		if v.dx > 0 {
			v.dx--
		} else if v.dx < 0 {
			v.dx++
		}

		v.dy--

		if position.x >= target.xMin && position.x <= target.xMax && position.y >= target.yMin && position.y <= target.yMax {
			return true, maxY
		}
	}

	return false, maxY
}

func findHighestY(target area) int {
	highestY := 0

	for dx := 0; dx <= target.xMax; dx++ {
		for dy := target.yMin; dy <= -target.yMin; dy++ {
			v := velocity{dx, dy}
			success, maxY := simulate(v, target)
			if success && maxY > highestY {
				highestY = maxY
			}
		}
	}

	return highestY
}

func findAllVelocities(target area) int {
	count := 0

	for dx := 0; dx <= target.xMax; dx++ {
		for dy := target.yMin; dy <= -target.yMin; dy++ {
			v := velocity{dx, dy}
			hit, _ := simulate(v, target)
			if hit {
				count++
			}
		}
	}

	return count
}
