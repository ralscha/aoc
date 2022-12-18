package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"container/list"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	inputFile := "./2022/18/input.txt"
	input, err := download.ReadInput(inputFile, 2022, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

var directions = []point3d{
	{-1, 0, 0},
	{1, 0, 0},
	{0, -1, 0},
	{0, 1, 0},
	{0, 0, -1},
	{0, 0, 1},
}

type point3d struct {
	x, y, z int
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	cubes := createCubs(lines)
	surfaceArea := 0
	for cube := range cubes {
		for _, d := range directions {
			if _, ok := cubes[point3d{cube.x + d.x, cube.y + d.y, cube.z + d.z}]; !ok {
				surfaceArea++
			}
		}

	}
	fmt.Println(surfaceArea)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	cubes := createCubs(lines)

	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := 0, 0, 0
	for cube := range cubes {
		if cube.x < minX {
			minX = cube.x
		}
		if cube.y < minY {
			minY = cube.y
		}
		if cube.z < minZ {
			minZ = cube.z
		}
		if cube.x > maxX {
			maxX = cube.x
		}
		if cube.y > maxY {
			maxY = cube.y
		}
		if cube.z > maxZ {
			maxZ = cube.z
		}
	}

	rangeX := maxX - minX + 3
	rangeY := maxY - minY + 3
	rangeZ := maxZ - minZ + 3
	boundingBoxSurfaceArea := 2 * (rangeX*rangeY + rangeZ*rangeY + rangeX*rangeZ)

	minX--
	minY--
	minZ--
	maxX++
	maxY++
	maxZ++

	queue := list.New()
	start := point3d{minX, minY, minZ}
	queue.PushBack(start)
	visited := make(map[point3d]struct{})
	visited[start] = struct{}{}

	for queue.Len() > 0 {
		s := queue.Front()
		cube := s.Value.(point3d)
		queue.Remove(s)

		for _, d := range directions {
			neighbour := point3d{cube.x + d.x, cube.y + d.y, cube.z + d.z}
			if _, ok := visited[neighbour]; !ok {
				if _, ok := cubes[neighbour]; !ok {
					if neighbour.x >= minX && neighbour.x <= maxX &&
						neighbour.y >= minY && neighbour.y <= maxY &&
						neighbour.z >= minZ && neighbour.z <= maxZ {
						visited[neighbour] = struct{}{}
						queue.PushBack(neighbour)
					}
				}
			}
		}
	}

	surfaceAreaVisited := 0
	for cube := range visited {
		for _, d := range directions {
			neighbour := point3d{cube.x + d.x, cube.y + d.y, cube.z + d.z}
			if _, ok := visited[neighbour]; !ok {
				surfaceAreaVisited++
			}
		}
	}

	fmt.Println(surfaceAreaVisited - boundingBoxSurfaceArea)
}

func createCubs(lines []string) map[point3d]struct{} {
	cubes := make(map[point3d]struct{})
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x, y, z := conv.MustAtoi(splitted[0]), conv.MustAtoi(splitted[1]), conv.MustAtoi(splitted[2])
		cubes[point3d{x, y, z}] = struct{}{}
	}
	return cubes
}
