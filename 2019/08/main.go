package main

import (
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2019, 8)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1(input)
	part2(input)
}

func part1(input string) {
	digits := []rune(input)
	width := 25
	height := 6

	layers := make([][][]int, 0)
	layer := make([][]int, 0)
	row := make([]int, 0)
	for i, digit := range digits {
		if i%width == 0 && i != 0 {
			layer = append(layer, row)
			row = make([]int, 0)
			if len(layer) == height {
				layers = append(layers, layer)
				layer = make([][]int, 0)
			}
		}
		row = append(row, int(digit-'0'))
	}
	layer = append(layer, row)
	layers = append(layers, layer)

	minZeros := math.MaxInt
	result := 0
	for _, layer := range layers {
		zeros := 0
		ones := 0
		twos := 0
		for _, row := range layer {
			for _, digit := range row {
				if digit == 0 {
					zeros++
				} else if digit == 1 {
					ones++
				} else if digit == 2 {
					twos++
				}
			}
		}
		if zeros < minZeros {
			minZeros = zeros
			result = ones * twos
		}
	}

	fmt.Println("Part 1:", result)

}

func part2(input string) {
	digits := []rune(input)
	width := 25
	height := 6

	layers := make([][][]int, 0)
	layer := make([][]int, 0)
	row := make([]int, 0)
	for i, digit := range digits {
		if i%width == 0 && i != 0 {
			layer = append(layer, row)
			row = make([]int, 0)
			if len(layer) == height {
				layers = append(layers, layer)
				layer = make([][]int, 0)
			}
		}
		row = append(row, int(digit-'0'))
	}
	layer = append(layer, row)
	layers = append(layers, layer)

	image := make([][]int, height)
	for i := 0; i < height; i++ {
		image[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for _, layer := range layers {
				if layer[i][j] != 2 {
					image[i][j] = layer[i][j]
					break
				}
			}
		}
	}

	fmt.Println("Part 2:")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if image[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}
