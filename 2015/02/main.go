package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"aoc/internal/mathx"
	"fmt"
	"log"
	"strings"
)

func main() {
	inputFile := "./2015/02/input.txt"
	input, err := download.ReadInput(inputFile, 2015, 2)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	totalWrappingPaper := 0
	totalRibbon := 0
	for _, line := range conv.SplitNewline(input) {
		parts := strings.Split(line, "x")
		l, w, h := conv.MustAtoi(parts[0]), conv.MustAtoi(parts[1]), conv.MustAtoi(parts[2])
		totalWrappingPaper += 2*l*w + 2*w*h + 2*h*l + mathx.Min(l*w, w*h, h*l)
		totalRibbon += 2*mathx.Min(l+w, w+h, h+l) + l*w*h
	}

	fmt.Printf("Total wrapping paper: %d\n", totalWrappingPaper)
	fmt.Printf("Total ribbon: %d\n", totalRibbon)

}
