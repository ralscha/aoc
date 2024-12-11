package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/geomutil"
	"aoc/internal/gridutil"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2022, 15)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type sensor struct {
	pos               gridutil.Coordinate
	closestBeacon     gridutil.Coordinate
	manhattanDistance int
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	sensors := createSensors(lines)

	targetY := 2000000
	covers := container.NewSet[gridutil.Coordinate]()
	beacons := container.NewSet[gridutil.Coordinate]()

	for _, sensor := range sensors {
		if sensor.closestBeacon.Row == targetY {
			beacons.Add(sensor.closestBeacon)
		}
		radius := sensor.manhattanDistance
		distance := mathx.Abs(sensor.pos.Row - targetY)
		if distance <= radius {
			diff := radius - distance
			for x := sensor.pos.Col - diff; x <= sensor.pos.Col+diff; x++ {
				covers.Add(gridutil.Coordinate{Row: targetY, Col: x})
			}
		}
	}
	fmt.Println(covers.Len() - beacons.Len())
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	sensors := createSensors(lines)

	for y := 0; y <= 4000000; y++ {
		for x := 0; x <= 4000000; x++ {
			p := gridutil.Coordinate{Row: y, Col: x}
			jump := false
			for _, sensor := range sensors {
				distance := geomutil.ManhattanDistance(p, sensor.pos)
				if distance <= sensor.manhattanDistance {
					x = sensor.pos.Col + sensor.manhattanDistance - mathx.Abs(sensor.pos.Row-p.Row)
					jump = true
					break
				}
			}
			if !jump {
				fmt.Println(p.Col*4000000 + p.Row)
			}
		}
	}
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

		sensorPos := gridutil.Coordinate{Row: sensorY, Col: sensorX}
		beaconPos := gridutil.Coordinate{Row: beaconY, Col: beaconX}

		sensors = append(sensors, sensor{
			pos:               sensorPos,
			closestBeacon:     beaconPos,
			manhattanDistance: geomutil.ManhattanDistance(sensorPos, beaconPos),
		})
	}
	return sensors
}
