package main

import (
	"aoc/internal/conv"
	"aoc/internal/download"
	"fmt"
	"log"
	"unicode"
)

func main() {
	input, err := download.ReadInput(2023, 3)
	if err != nil {
		log.Fatalf("reading input failed: %v", err)
	}

	part1and2(input)
}

type position struct {
	row int
	col int
}

func part1and2(input string) {
	lines := conv.SplitNewline(input)

	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	numRows := len(lines)
	numCols := len(lines[0])
	matrix := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		matrix[i] = []rune(lines[i])
	}
	sum := 0

	starNumbers := make(map[position][]int)

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if unicode.IsDigit(matrix[row][col]) {
				number := string(matrix[row][col])
				startCol := col
				for col+1 < numCols && unicode.IsDigit(matrix[row][col+1]) {
					number += string(matrix[row][col+1])
					col++
				}
				endCol := col
				numberInt := conv.MustAtoi(number)

				hasSymbolNeighbor := false
				for i := startCol; i <= endCol; i++ {
					for _, direction := range directions {
						testRow := row + direction[0]
						testCol := i + direction[1]
						if testRow < 0 || testRow >= numRows || testCol < 0 || testCol >= numCols {
							continue
						}
						testRune := matrix[testRow][testCol]
						if !unicode.IsDigit(testRune) && testRune != '.' {
							if testRune == '*' {
								key := position{testRow, testCol}
								if _, ok := starNumbers[key]; !ok {
									starNumbers[key] = []int{}
								}
								starNumbers[key] = append(starNumbers[key], numberInt)
							}
							hasSymbolNeighbor = true
							break
						}
					}
					if hasSymbolNeighbor {
						break
					}
				}

				if hasSymbolNeighbor {
					sum += numberInt
				}
			}
		}
	}

	sumGears := 0
	for _, numbers := range starNumbers {
		if len(numbers) == 2 {
			sumGears += numbers[0] * numbers[1]
		}
	}

	fmt.Println(sum)
	fmt.Println(sumGears)
}
