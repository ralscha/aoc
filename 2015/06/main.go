package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/06/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	const sep = " through "

	for _, line := range lines {
		if strings.HasPrefix(line, "turn off ") {
			coords := strings.Split(line[9:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid[x][y] = false
				}
			}
		} else if strings.HasPrefix(line, "turn on ") {
			coords := strings.Split(line[8:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid[x][y] = true
				}
			}
		} else if strings.HasPrefix(line, "toggle ") {
			coords := strings.Split(line[7:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] {
				count++
			}
		}
	}
	println(count)
}

func part2(input string) {
	lines := conv.SplitNewline(input)
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	const sep = " through "

	for _, line := range lines {
		if strings.HasPrefix(line, "turn off ") {
			coords := strings.Split(line[9:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					if grid[x][y] > 0 {
						grid[x][y]--
					}
				}
			}
		} else if strings.HasPrefix(line, "turn on ") {
			coords := strings.Split(line[8:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid[x][y]++
				}
			}
		} else if strings.HasPrefix(line, "toggle ") {
			coords := strings.Split(line[7:], sep)
			start := conv.MustAtoiArray(strings.Split(coords[0], ","))
			end := conv.MustAtoiArray(strings.Split(coords[1], ","))
			for x := start[0]; x <= end[0]; x++ {
				for y := start[1]; y <= end[1]; y++ {
					grid[x][y] += 2
				}
			}
		}
	}

	count := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count += grid[x][y]
		}
	}
	println(count)
}
