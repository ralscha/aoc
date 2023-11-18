package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"math"
)

func main() {
	inputFile := "./2016/06/input.txt"
	input, err := download.ReadInput(inputFile, 2016, 6)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}
	lines := conv.SplitNewline(input)
	part1and2(lines)
}

func part1and2(lines []string) {
	chars := make([][26]int, len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			chars[i][c-'a']++
		}
	}
	resultMax := ""
	resultMin := ""
	for _, char := range chars {
		maxCount := 0
		maxChar := 0
		minCount := math.MaxInt
		minChar := 0
		for i, c := range char {
			if c > maxCount {
				maxCount = c
				maxChar = i
			}
			if c < minCount {
				minCount = c
				minChar = i
			}
		}
		resultMax += string(rune(maxChar + 'a'))
		resultMin += string(rune(minChar + 'a'))
	}
	fmt.Println(resultMax)
	fmt.Println(resultMin)
}
