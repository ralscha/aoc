package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

func main() {
	input, err := download.ReadInput(2022, 10)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	x := 1
	cycle := 1
	lines := conv.SplitNewline(input)
	results := make([]int, 6)
	for _, line := range lines {
		instructions := 1
		if line[:4] == "addx" {
			instructions = 2
		}
		for range instructions {
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signalStrength := cycle * x
				results = append(results, signalStrength)
			}
			cycle++
		}

		if line[:4] == "addx" {
			value := conv.MustAtoi(line[5:])
			x += value
		}
	}

	total := 0
	for _, result := range results {
		total += result
	}
	fmt.Println(total)
}

func part2(input string) {
	crt := gridutil.NewGrid2D[rune](false)
	for row := range 6 {
		for col := range 40 {
			crt.Set(row, col, '.')
		}
	}

	x := 1
	cycle := 1
	lines := conv.SplitNewline(input)
	crtRow := 0
	crtCol := 0

	for _, line := range lines {
		instructions := 1
		if line[:4] == "addx" {
			instructions = 2
		}
		for range instructions {
			if crtCol == x || crtCol == x+1 || crtCol == x-1 {
				crt.Set(crtRow, crtCol, '#')
			}
			crtCol++
			if crtCol == 40 {
				crtCol = 0
				crtRow++
			}
			cycle++
		}

		if line[:4] == "addx" {
			value := conv.MustAtoi(line[5:])
			x += value
		}
	}

	// Print the CRT display
	for row := range 6 {
		for col := range 40 {
			if val, ok := crt.Get(row, col); ok {
				fmt.Print(string(val))
			}
		}
		fmt.Println()
	}
}
