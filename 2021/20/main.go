package main

import (
	"aoc/internal/container"
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/gridutil"
	"fmt"
	"log"
)

type image struct {
	lit            *container.Set[gridutil.Coordinate]
	minRow, maxRow int
	minCol, maxCol int
	infiniteIsLit  bool
}

func main() {
	input, err := download.ReadInput(2021, 20)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input, 2)
	part1and2(input, 50)
}

func part1and2(input string, steps int) {
	lines := conv.SplitNewline(input)

	// Parse enhancement algorithm
	algorithm := make([]bool, len(lines[0]))
	for i, c := range lines[0] {
		algorithm[i] = c == '#'
	}

	// Parse initial image
	img := newImage()
	for row, line := range lines[2:] {
		for col, c := range line {
			if c == '#' {
				coord := gridutil.Coordinate{Row: row, Col: col}
				img.lit.Add(coord)
				img.updateBounds(coord)
			}
		}
	}

	// Enhance image specified number of times
	for i := 0; i < steps; i++ {
		img = enhance(img, algorithm)
	}

	fmt.Println("Lit pixels:", img.lit.Len())
}

func newImage() *image {
	return &image{
		lit:           container.NewSet[gridutil.Coordinate](),
		minRow:        0,
		maxRow:        0,
		minCol:        0,
		maxCol:        0,
		infiniteIsLit: false,
	}
}

func (img *image) updateBounds(p gridutil.Coordinate) {
	if p.Row < img.minRow {
		img.minRow = p.Row
	}
	if p.Row > img.maxRow {
		img.maxRow = p.Row
	}
	if p.Col < img.minCol {
		img.minCol = p.Col
	}
	if p.Col > img.maxCol {
		img.maxCol = p.Col
	}
}

func (img *image) isLit(p gridutil.Coordinate) bool {
	// If point is within known bounds, check the set
	if p.Row >= img.minRow && p.Row <= img.maxRow &&
		p.Col >= img.minCol && p.Col <= img.maxCol {
		return img.lit.Contains(p)
	}
	// Otherwise, return the infinite space state
	return img.infiniteIsLit
}

func enhance(img *image, algorithm []bool) *image {
	newImg := newImage()

	// Expand bounds by 1 in each direction
	minRow := img.minRow - 1
	maxRow := img.maxRow + 1
	minCol := img.minCol - 1
	maxCol := img.maxCol + 1

	// Process each pixel in the expanded area
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			coord := gridutil.Coordinate{Row: row, Col: col}
			if shouldBeLit(img, coord, algorithm) {
				newImg.lit.Add(coord)
				newImg.updateBounds(coord)
			}
		}
	}

	// Update infinite space state
	if img.infiniteIsLit {
		newImg.infiniteIsLit = algorithm[511] // All 9 bits set
	} else {
		newImg.infiniteIsLit = algorithm[0] // All 9 bits clear
	}

	return newImg
}

func shouldBeLit(img *image, p gridutil.Coordinate, algorithm []bool) bool {
	index := 0
	for row := -1; row <= 1; row++ {
		for col := -1; col <= 1; col++ {
			neighbor := gridutil.Coordinate{
				Row: p.Row + row,
				Col: p.Col + col,
			}
			index <<= 1
			if img.isLit(neighbor) {
				index |= 1
			}
		}
	}
	return algorithm[index]
}
