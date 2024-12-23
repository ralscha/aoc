package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2020, 12)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	x, y := 0, 0
	direction := "E"

	for _, line := range lines {
		action := line[0]
		value := conv.MustAtoi(line[1:])

		switch action {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			for range value / 90 {
				switch direction {
				case "N":
					direction = "W"
				case "E":
					direction = "N"
				case "S":
					direction = "E"
				default:
					direction = "S"
				}
			}
		case 'R':
			for range value / 90 {
				switch direction {
				case "N":
					direction = "E"
				case "E":
					direction = "S"
				case "S":
					direction = "W"
				default:
					direction = "N"
				}
			}
		case 'F':
			switch direction {
			case "N":
				y += value
			case "S":
				y -= value
			case "E":
				x += value
			default:
				x -= value
			}
		}
	}

	manhattanDistance := mathx.Abs(x) + mathx.Abs(y)
	fmt.Println("Part 1", manhattanDistance)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	shipX, shipY := 0, 0
	waypointX, waypointY := 10, 1

	for _, line := range lines {
		action := line[0]
		value := conv.MustAtoi(line[1:])

		switch action {
		case 'N':
			waypointY += value
		case 'S':
			waypointY -= value
		case 'E':
			waypointX += value
		case 'W':
			waypointX -= value
		case 'L':
			for i := 0; i < value/90; i++ {
				temp := waypointX
				waypointX = -waypointY
				waypointY = temp
			}
		case 'R':
			for i := 0; i < value/90; i++ {
				temp := waypointX
				waypointX = waypointY
				waypointY = -temp
			}
		case 'F':
			shipX += value * waypointX
			shipY += value * waypointY
		}
	}

	manhattanDistance := mathx.Abs(shipX) + mathx.Abs(shipY)
	fmt.Println("Part 2", manhattanDistance)
}
