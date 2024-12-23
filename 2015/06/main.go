package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
	"strings"
)

func main() {
	input, err := download.ReadInput(2015, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[bool](false)
	grid.SetMaxRowCol(999, 999)

	const sep = " through "

	for _, line := range lines {
		if strings.HasPrefix(line, "turn off ") {
			coords := strings.Split(line[9:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid.Set(x, y, false)
				}
			}
		} else if strings.HasPrefix(line, "turn on ") {
			coords := strings.Split(line[8:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid.Set(x, y, true)
				}
			}
		} else if strings.HasPrefix(line, "toggle ") {
			coords := strings.Split(line[7:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					val, _ := grid.Get(x, y)
					grid.Set(x, y, !val)
				}
			}
		}
	}

	count := 0
	for x := range 1000 {
		for y := range 1000 {
			if val, _ := grid.Get(x, y); val {
				count++
			}
		}
	}
	fmt.Println("Part 1", count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := gridutil.NewGrid2D[int](false)
	grid.SetMaxRowCol(999, 999)

	const sep = " through "

	for _, line := range lines {
		if strings.HasPrefix(line, "turn off ") {
			coords := strings.Split(line[9:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					val, _ := grid.Get(x, y)
					if val > 0 {
						grid.Set(x, y, val-1)
					}
				}
			}
		} else if strings.HasPrefix(line, "turn on ") {
			coords := strings.Split(line[8:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					val, _ := grid.Get(x, y)
					grid.Set(x, y, val+1)
				}
			}
		} else if strings.HasPrefix(line, "toggle ") {
			coords := strings.Split(line[7:], sep)
			start := conv.ToIntSlice(strings.Split(coords[0], ","))
			end := conv.ToIntSlice(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					val, _ := grid.Get(x, y)
					grid.Set(x, y, val+2)
				}
			}
		}
	}

	count := 0
	for x := range 1000 {
		for y := range 1000 {
			val, _ := grid.Get(x, y)
			count += val
		}
	}
	fmt.Println("Part 2", count)
}
