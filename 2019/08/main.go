package main

import (
	"aoc/internal/download"
	"aoc/internal/gridutil"
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

	layers := make([]gridutil.Grid2D[int], 0)
	currentLayer := gridutil.NewGrid2D[int](false)

	for i, digit := range digits {
		row := i / width % height
		col := i % width
		if i > 0 && i%(width*height) == 0 {
			layers = append(layers, currentLayer)
			currentLayer = gridutil.NewGrid2D[int](false)
		}
		currentLayer.Set(row, col, int(digit-'0'))
	}
	layers = append(layers, currentLayer)

	minZeros := math.MaxInt
	result := 0
	for _, layer := range layers {
		zeros := 0
		ones := 0
		twos := 0
		for row := 0; row < height; row++ {
			for col := 0; col < width; col++ {
				digit, _ := layer.Get(row, col)
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

	layers := make([]gridutil.Grid2D[int], 0)
	currentLayer := gridutil.NewGrid2D[int](false)

	for i, digit := range digits {
		row := i / width % height
		col := i % width
		if i > 0 && i%(width*height) == 0 {
			layers = append(layers, currentLayer)
			currentLayer = gridutil.NewGrid2D[int](false)
		}
		currentLayer.Set(row, col, int(digit-'0'))
	}
	layers = append(layers, currentLayer)

	finalImage := gridutil.NewGrid2D[int](false)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			for _, layer := range layers {
				pixel, _ := layer.Get(row, col)
				if pixel != 2 { // not transparent
					finalImage.Set(row, col, pixel)
					break
				}
			}
		}
	}

	fmt.Println("Part 2:")
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			pixel, _ := finalImage.Get(row, col)
			if pixel == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}
