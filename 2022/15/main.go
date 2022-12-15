package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2022/15/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type point struct {
	x, y int
}
type sensor struct {
	pos               point
	closestBeacon     point
	manhattanDistance int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sensors := createSensors(lines)

	targetY := 2000000
	covers := make(map[point]bool)
	beacons := make(map[point]bool)

	for _, sensor := range sensors {
		if sensor.closestBeacon.y == targetY {
			beacons[sensor.closestBeacon] = true
		}
		radius := sensor.manhattanDistance
		distance := mathx.Abs(sensor.pos.y - targetY)
		if distance <= radius {
			diff := radius - distance
			for x := sensor.pos.x - diff; x <= sensor.pos.x+diff; x++ {
				covers[point{x, targetY}] = true
			}
		}
	}
	fmt.Println(len(covers) - len(beacons))
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	sensors := createSensors(lines)

	for y := 0; y <= 4000000; y++ {
		for x := 0; x <= 4000000; x++ {
			p := point{x, y}
			jump := false
			for _, sensor := range sensors {
				distance := manhattanDistance(p, sensor.pos)
				if distance <= sensor.manhattanDistance {
					x = sensor.pos.x + sensor.manhattanDistance - mathx.Abs(sensor.pos.y-p.y)
					jump = true
					break
				}
			}
			if !jump {
				fmt.Println(p.x*4000000 + p.y)
			}
		}
	}
}

func manhattanDistance(a, b point) int {
	return mathx.Abs(a.x-b.x) + mathx.Abs(a.y-b.y)
}

func createSensors(lines []string) []sensor {
	var sensors []sensor
	for _, line := range lines {
		splitted := strings.Fields(line)
		s := strings.TrimPrefix(splitted[2], "x=")
		sensorX := conv.MustAtoi(s[:len(s)-1])
		s = strings.TrimPrefix(splitted[3], "y=")
		sensorY := conv.MustAtoi(s[:len(s)-1])
		s = strings.TrimPrefix(splitted[8], "x=")
		beaconX := conv.MustAtoi(s[:len(s)-1])
		beaconY := conv.MustAtoi(strings.TrimPrefix(splitted[9], "y="))

		sensors = append(sensors, sensor{
			pos:               point{sensorX, sensorY},
			closestBeacon:     point{beaconX, beaconY},
			manhattanDistance: manhattanDistance(point{sensorX, sensorY}, point{beaconX, beaconY}),
		})
	}
	return sensors
}
