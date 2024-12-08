package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

type cuboid struct {
	min, max gridutil.Coordinate3D
	state    bool // true for on, false for off
}

func main() {
	input, err := download.ReadInput(2021, 22)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	steps := parseInput(input)
	cubes := container.NewSet[gridutil.Coordinate3D]()

	// Process only steps within -50..50 region
	for _, step := range steps {
		if !isInInitRegion(step) {
			continue
		}

		for x := step.min.X; x <= step.max.X; x++ {
			for y := step.min.Y; y <= step.max.Y; y++ {
				for z := step.min.Z; z <= step.max.Z; z++ {
					pos := gridutil.Coordinate3D{X: x, Y: y, Z: z}
					if step.state {
						cubes.Add(pos)
					} else {
						cubes.Remove(pos)
					}
				}
			}
		}
	}

	fmt.Println("Part 1", cubes.Len())
}

func part2(input string) {
	steps := parseInput(input)
	var regions []cuboid

	for _, step := range steps {
		var newRegions []cuboid

		// If turning on, add this cuboid
		if step.state {
			newRegions = append(newRegions, step)
		}

		// Add intersections with all previous regions
		for _, region := range regions {
			if intersection := intersectCuboids(step, region); intersection != nil {
				// Add intersection with opposite state to cancel out overlap
				intersection.state = !region.state
				newRegions = append(newRegions, *intersection)
			}
		}

		regions = append(regions, newRegions...)
	}

	// Len total cubes that are on
	total := int64(0)
	for _, region := range regions {
		volume := cuboidVolume(region)
		if region.state {
			total += volume
		} else {
			total -= volume
		}
	}

	fmt.Println("Part 2", total)
}

func parseInput(input string) []cuboid {
	lines := conv.SplitNewline(input)
	var steps []cuboid

	for _, line := range lines {
		fields := strings.Fields(line)
		state := fields[0] == "on"

		// Parse coordinates
		coords := make([]int, 6) // [xmin, xmax, ymin, ymax, zmin, zmax]
		parts := strings.Split(fields[1], ",")
		for i, part := range parts {
			nums := strings.Split(strings.Split(part, "=")[1], "..")
			coords[i*2] = conv.MustAtoi(nums[0])
			coords[i*2+1] = conv.MustAtoi(nums[1])
		}

		steps = append(steps, cuboid{
			min:   gridutil.Coordinate3D{X: coords[0], Y: coords[2], Z: coords[4]},
			max:   gridutil.Coordinate3D{X: coords[1], Y: coords[3], Z: coords[5]},
			state: state,
		})
	}

	return steps
}

func isInInitRegion(c cuboid) bool {
	return c.min.X >= -50 && c.max.X <= 50 &&
		c.min.Y >= -50 && c.max.Y <= 50 &&
		c.min.Z >= -50 && c.max.Z <= 50
}

func intersectCuboids(a, b cuboid) *cuboid {
	minValue := gridutil.Coordinate3D{
		X: max(a.min.X, b.min.X),
		Y: max(a.min.Y, b.min.Y),
		Z: max(a.min.Z, b.min.Z),
	}
	maxValue := gridutil.Coordinate3D{
		X: min(a.max.X, b.max.X),
		Y: min(a.max.Y, b.max.Y),
		Z: min(a.max.Z, b.max.Z),
	}

	// Check if there is no intersection
	if minValue.X > maxValue.X || minValue.Y > maxValue.Y || minValue.Z > maxValue.Z {
		return nil
	}

	return &cuboid{min: minValue, max: maxValue}
}

func cuboidVolume(c cuboid) int64 {
	return int64(c.max.X-c.min.X+1) *
		int64(c.max.Y-c.min.Y+1) *
		int64(c.max.Z-c.min.Z+1)
}
