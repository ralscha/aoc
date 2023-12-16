package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	input, err := download.ReadInput(2021, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 2)
	part1and2(input, 50)
}

type point struct {
	row, col int
}

func part1and2(input string, steps int) {
	lines := conv.SplitNewline(input)

	imageEnhancementAlgorithm := make([]byte, 0, len(lines[0]))
	for _, c := range lines[0] {
		imageEnhancementAlgorithm = append(imageEnhancementAlgorithm, byte(c))
	}
	infiniteSpaceStaysOff := imageEnhancementAlgorithm[0] == '.'

	inputImage := make(map[point]byte)
	for row, line := range lines[2:] {
		for col, c := range line {
			inputImage[point{row, col}] = byte(c)
		}
	}

	enhancedImage := inputImage
	for n := 0; n < steps; n++ {
		infiniteSpaceIsOn := n%2 == 1
		enhancedImage = enhance(enhancedImage, imageEnhancementAlgorithm, infiniteSpaceStaysOff, infiniteSpaceIsOn)
	}

	count := 0
	for _, v := range enhancedImage {
		if v == '#' {
			count++
		}
	}
	fmt.Println(count)
}

func minMax(input map[point]byte) (int, int, int, int) {
	maxRow, maxCol := math.MinInt64, math.MinInt64
	minRow, minCol := math.MaxInt64, math.MaxInt64
	for p := range input {
		maxRow = max(maxRow, p.row)
		maxCol = max(maxCol, p.col)
		minRow = min(minRow, p.row)
		minCol = min(minCol, p.col)
	}
	return minRow, minCol, maxRow, maxCol
}

func enhance(input map[point]byte, algorithm []byte, infiniteSpaceStaysOff, infiniteSpaceIsOn bool) map[point]byte {

	minRow, minCol, maxRow, maxCol := minMax(input)

	var infChar byte = '.'
	if !infiniteSpaceStaysOff && infiniteSpaceIsOn {
		infChar = '#'
	}

	for c := minCol - 3; c <= maxCol+3; c++ {
		input[point{minRow - 1, c}] = infChar
		input[point{minRow - 2, c}] = infChar
		input[point{minRow - 3, c}] = infChar
		input[point{maxRow + 1, c}] = infChar
		input[point{maxRow + 2, c}] = infChar
		input[point{maxRow + 3, c}] = infChar
	}

	for r := minRow - 3; r <= maxRow+3; r++ {
		input[point{r, minCol - 1}] = infChar
		input[point{r, minCol - 2}] = infChar
		input[point{r, minCol - 3}] = infChar
		input[point{r, maxCol + 1}] = infChar
		input[point{r, maxCol + 2}] = infChar
		input[point{r, maxCol + 3}] = infChar
	}

	enhancedImage := map[point]byte{}
	for r := minRow - 2; r <= maxRow+2; r++ {
		for c := minCol - 2; c <= maxCol+2; c++ {
			p := point{r, c}
			enhancedImage[p] = enhancePixel(input, p, algorithm)
		}
	}
	return enhancedImage
}

func enhancePixel(input map[point]byte, p point, algorithm []byte) byte {
	pixelValue := 0
	for _, d := range []point{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	} {
		n := point{p.row + d.row, p.col + d.col}
		pixelValue <<= 1
		if input[n] == '#' {
			pixelValue |= 1
		}
	}
	return algorithm[pixelValue]
}
