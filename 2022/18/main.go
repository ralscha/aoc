package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	input, err := download.ReadInput(2022, 18)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

var directions = []gridutil.Coordinate3D{
	{X: -1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 0, Z: 1},
}

func part1(input string) {
	lines := conv.SplitNewline(input)

	cubes := createCubes(lines)
	surfaceArea := 0
	for _, cube := range cubes.Values() {
		for _, d := range directions {
			neighbor := cube.Add(d)
			if !cubes.Contains(neighbor) {
				surfaceArea++
			}
		}
	}
	fmt.Println(surfaceArea)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	cubes := createCubes(lines)

	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := 0, 0, 0
	for _, cube := range cubes.Values() {
		if cube.X < minX {
			minX = cube.X
		}
		if cube.Y < minY {
			minY = cube.Y
		}
		if cube.Z < minZ {
			minZ = cube.Z
		}
		if cube.X > maxX {
			maxX = cube.X
		}
		if cube.Y > maxY {
			maxY = cube.Y
		}
		if cube.Z > maxZ {
			maxZ = cube.Z
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

	queue := container.NewQueue[gridutil.Coordinate3D]()
	start := gridutil.Coordinate3D{X: minX, Y: minY, Z: minZ}
	queue.Push(start)
	visited := container.NewSet[gridutil.Coordinate3D]()
	visited.Add(start)

	for !queue.IsEmpty() {
		cube := queue.Pop()

		for _, d := range directions {
			neighbor := cube.Add(d)
			if !visited.Contains(neighbor) && !cubes.Contains(neighbor) {
				if neighbor.X >= minX && neighbor.X <= maxX &&
					neighbor.Y >= minY && neighbor.Y <= maxY &&
					neighbor.Z >= minZ && neighbor.Z <= maxZ {
					visited.Add(neighbor)
					queue.Push(neighbor)
				}
			}
		}
	}

	surfaceAreaVisited := 0
	for _, cube := range visited.Values() {
		for _, d := range directions {
			neighbor := cube.Add(d)
			if !visited.Contains(neighbor) {
				surfaceAreaVisited++
			}
		}
	}

	fmt.Println(surfaceAreaVisited - boundingBoxSurfaceArea)
}

func createCubes(lines []string) *container.Set[gridutil.Coordinate3D] {
	cubes := container.NewSet[gridutil.Coordinate3D]()
	for _, line := range lines {
		splitted := strings.Split(line, ",")
		x, y, z := conv.MustAtoi(splitted[0]), conv.MustAtoi(splitted[1]), conv.MustAtoi(splitted[2])
		cubes.Add(gridutil.Coordinate3D{X: x, Y: y, Z: z})
	}
	return cubes
}
