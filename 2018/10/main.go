package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2018, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type point struct {
	x, y                 int
	velocityX, velocityY int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)
	points := make([]point, len(lines))
	for i, line := range lines {
		splitted := strings.Split(line, "=")
		last := strings.LastIndex(splitted[1], ">")
		pos := splitted[1][1:last]
		velocity := splitted[2][1 : len(splitted[2])-1]

		posSplitted := strings.Split(pos, ",")
		velocitySplitted := strings.Split(velocity, ",")

		points[i].x = conv.MustAtoi(strings.TrimSpace(posSplitted[0]))
		points[i].y = conv.MustAtoi(strings.TrimSpace(posSplitted[1]))
		points[i].velocityX = conv.MustAtoi(strings.TrimSpace(velocitySplitted[0]))
		points[i].velocityY = conv.MustAtoi(strings.TrimSpace(velocitySplitted[1]))
	}

	minBoundingBox := math.MaxInt
	minBoundingBoxIndex := 0

	for i := 0; i < 100000; i++ {
		minX := math.MaxInt
		maxX := math.MinInt
		minY := math.MaxInt
		maxY := math.MinInt

		for _, p := range points {
			x := p.x + i*p.velocityX
			y := p.y + i*p.velocityY

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}

		if maxX-minX+maxY-minY < minBoundingBox {
			minBoundingBox = maxX - minX + maxY - minY
			minBoundingBoxIndex = i
		}
	}

	fmt.Println("Part2:", minBoundingBoxIndex)

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for _, p := range points {
		x := p.x + minBoundingBoxIndex*p.velocityX
		y := p.y + minBoundingBoxIndex*p.velocityY

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	canvas := make([][]rune, maxY-minY+1)
	for i := range canvas {
		canvas[i] = make([]rune, maxX-minX+1)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	for _, p := range points {
		x := p.x + minBoundingBoxIndex*p.velocityX
		y := p.y + minBoundingBoxIndex*p.velocityY

		canvas[y-minY][x-minX] = '*'
	}

	for _, line := range canvas {
		fmt.Println(string(line))
	}

}
