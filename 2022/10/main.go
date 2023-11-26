package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
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
		for i := 0; i < instructions; i++ {
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
	crt := make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
	}
	for i := range crt {
		for j := range crt[i] {
			crt[i][j] = "."
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
		for i := 0; i < instructions; i++ {
			if crtCol == x || crtCol == x+1 || crtCol == x-1 {
				crt[crtRow][crtCol] = "#"
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

	for i := range crt {
		for j := range crt[i] {
			fmt.Print(crt[i][j])
		}
		fmt.Println()
	}
}
