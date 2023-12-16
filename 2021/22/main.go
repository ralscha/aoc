package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2021, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

type cube struct {
	x, y, z int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	cubes := make(map[cube]bool)
	for _, line := range lines {
		fields := strings.Fields(line)
		command := fields[0]
		xzy := strings.Split(fields[1], ",")
		xPart := strings.Split(xzy[0], "=")
		yPart := strings.Split(xzy[1], "=")
		zPart := strings.Split(xzy[2], "=")

		xPart = strings.Split(xPart[1], "..")
		fromX := conv.MustAtoi(xPart[0])
		toX := conv.MustAtoi(xPart[1])

		yPart = strings.Split(yPart[1], "..")
		fromY := conv.MustAtoi(yPart[0])
		toY := conv.MustAtoi(yPart[1])

		zPart = strings.Split(zPart[1], "..")
		fromZ := conv.MustAtoi(zPart[0])
		toZ := conv.MustAtoi(zPart[1])

		turnOn := command == "on"
		for x := fromX; x <= toX; x++ {
			if x < -50 || x > 50 {
				continue
			}
			for y := fromY; y <= toY; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for z := fromZ; z <= toZ; z++ {
					if z < -50 || z > 50 {
						continue
					}
					cubes[cube{x, y, z}] = turnOn
				}
			}
		}
	}

	count := 0
	for _, on := range cubes {
		if on {
			count++
		}
	}
	fmt.Println(count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)

	var cuboids []cuboid
	for _, line := range lines {
		cuboid := cuboid{}
		fields := strings.Fields(line)
		command := fields[0]
		if command == "on" {
			cuboid.state = 1
		} else {
			cuboid.state = -1
		}
		xzy := strings.Split(fields[1], ",")
		xPart := strings.Split(xzy[0], "=")
		yPart := strings.Split(xzy[1], "=")
		zPart := strings.Split(xzy[2], "=")

		xPart = strings.Split(xPart[1], "..")
		yPart = strings.Split(yPart[1], "..")
		zPart = strings.Split(zPart[1], "..")
		cuboid.xmin, cuboid.xmax = conv.MustAtoi(xPart[0]), conv.MustAtoi(xPart[1])
		cuboid.ymin, cuboid.ymax = conv.MustAtoi(yPart[0]), conv.MustAtoi(yPart[1])
		cuboid.zmin, cuboid.zmax = conv.MustAtoi(zPart[0]), conv.MustAtoi(zPart[1])
		cuboids = append(cuboids, cuboid)
	}

	var cores []cuboid
	for _, cu := range cuboids {
		var toAdd []cuboid
		if cu.state == 1 {
			toAdd = append(toAdd, cu)
		}
		for _, core := range cores {
			inter := intersection(cu, core)
			if inter != nil {
				toAdd = append(toAdd, *inter)
			}
		}
		cores = append(cores, toAdd...)
	}

	count := 0
	for _, c := range cores {
		count += c.state * (c.xmax - c.xmin + 1) * (c.ymax - c.ymin + 1) * (c.zmax - c.zmin + 1)
	}
	fmt.Println(count)
}

type cuboid struct {
	state int
	xmin  int
	xmax  int
	ymin  int
	ymax  int
	zmin  int
	zmax  int
}

func intersection(s, t cuboid) *cuboid {
	n := cuboid{
		state: -t.state,
		xmin:  max(s.xmin, t.xmin),
		xmax:  min(s.xmax, t.xmax),
		ymin:  max(s.ymin, t.ymin),
		ymax:  min(s.ymax, t.ymax),
		zmin:  max(s.zmin, t.zmin),
		zmax:  min(s.zmax, t.zmax),
	}
	if n.xmin > n.xmax || n.ymin > n.ymax || n.zmin > n.zmax {
		return nil
	}
	return &n
}
